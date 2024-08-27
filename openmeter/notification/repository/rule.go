package repository

import (
	"context"
	"fmt"

	entdb "github.com/openmeterio/openmeter/openmeter/ent/db"
	channeldb "github.com/openmeterio/openmeter/openmeter/ent/db/notificationchannel"
	ruledb "github.com/openmeterio/openmeter/openmeter/ent/db/notificationrule"
	"github.com/openmeterio/openmeter/openmeter/notification"
	"github.com/openmeterio/openmeter/pkg/clock"
	"github.com/openmeterio/openmeter/pkg/framework/entutils"
	"github.com/openmeterio/openmeter/pkg/models"
	"github.com/openmeterio/openmeter/pkg/pagination"
	"github.com/openmeterio/openmeter/pkg/sortx"
)

func (r repository) ListRules(ctx context.Context, params notification.ListRulesInput) (pagination.PagedResponse[notification.Rule], error) {
	db := r.client()

	query := db.NotificationRule.Query().
		Where(ruledb.DeletedAtIsNil()) // Do not return deleted Rules

	if len(params.Namespaces) > 0 {
		query = query.Where(ruledb.NamespaceIn(params.Namespaces...))
	}

	if len(params.Rules) > 0 {
		query = query.Where(ruledb.IDIn(params.Rules...))
	}

	if !params.IncludeDisabled {
		query = query.Where(ruledb.Disabled(false))
	}

	if len(params.Types) > 0 {
		query = query.Where(ruledb.TypeIn(params.Types...))
	}

	if len(params.Channels) > 0 {
		query = query.Where(ruledb.HasChannelsWith(channeldb.IDIn(params.Channels...)))
	}

	// Eager load Channels
	query = query.WithChannels()

	order := entutils.GetOrdering(sortx.OrderDefault)
	if !params.Order.IsDefaultValue() {
		order = entutils.GetOrdering(params.Order)
	}

	switch params.OrderBy {
	case notification.RuleOrderByCreatedAt:
		query = query.Order(ruledb.ByCreatedAt(order...))
	case notification.RuleOrderByUpdatedAt:
		query = query.Order(ruledb.ByUpdatedAt(order...))
	case notification.RuleOrderByType:
		query = query.Order(ruledb.ByType(order...))
	case notification.RuleOrderByID:
		fallthrough
	default:
		query = query.Order(ruledb.ByID(order...))
	}

	response := pagination.PagedResponse[notification.Rule]{
		Page: params.Page,
	}

	paged, err := query.Paginate(ctx, params.Page)
	if err != nil {
		return response, err
	}

	result := make([]notification.Rule, 0, len(paged.Items))
	for _, ruleRow := range paged.Items {
		if ruleRow == nil {
			r.logger.Warn("invalid query result: nil notification rule received")
			continue
		}

		result = append(result, *RuleFromDBEntity(*ruleRow))
	}

	response.TotalCount = paged.TotalCount
	response.Items = result

	return response, nil
}

func (r repository) CreateRule(ctx context.Context, params notification.CreateRuleInput) (*notification.Rule, error) {
	db := r.client()

	query := db.NotificationRule.Create().
		SetType(params.Type).
		SetName(params.Name).
		SetNamespace(params.Namespace).
		SetDisabled(params.Disabled).
		SetConfig(params.Config).
		AddChannelIDs(params.Channels...)

	queryRow, err := query.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create notification rule: %w", err)
	}

	if queryRow == nil {
		return nil, fmt.Errorf("invalid query result: nil notification rule received")
	}

	channelsQuery := db.NotificationChannel.Query().
		Where(channeldb.Namespace(params.Namespace)).
		Where(channeldb.IDIn(params.Channels...))

	channelRows, err := channelsQuery.All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to query notification channels: %w", err)
	}

	queryRow.Edges.Channels = channelRows

	return RuleFromDBEntity(*queryRow), nil
}

func (r repository) DeleteRule(ctx context.Context, params notification.DeleteRuleInput) error {
	db := r.client()

	query := db.NotificationRule.UpdateOneID(params.ID).
		Where(ruledb.Namespace(params.Namespace)).
		SetDeletedAt(clock.Now().UTC()).
		SetDisabled(true)

	_, err := query.Save(ctx)
	if err != nil {
		if entdb.IsNotFound(err) {
			return notification.NotFoundError{
				NamespacedID: models.NamespacedID{
					Namespace: params.Namespace,
					ID:        params.ID,
				},
			}
		}

		return fmt.Errorf("failed top delete notification rule: %w", err)
	}

	return nil
}

func (r repository) GetRule(ctx context.Context, params notification.GetRuleInput) (*notification.Rule, error) {
	db := r.client()

	query := db.NotificationRule.Query().
		Where(ruledb.ID(params.ID)).
		Where(ruledb.Namespace(params.Namespace)).
		WithChannels()

	ruleRow, err := query.First(ctx)
	if err != nil {
		if entdb.IsNotFound(err) {
			return nil, notification.NotFoundError{
				NamespacedID: models.NamespacedID{
					Namespace: params.Namespace,
					ID:        params.ID,
				},
			}
		}

		return nil, fmt.Errorf("failed to fetch notification rule: %w", err)
	}

	if ruleRow == nil {
		return nil, fmt.Errorf("invalid query result: nil notification rule received")
	}

	return RuleFromDBEntity(*ruleRow), nil
}

func (r repository) UpdateRule(ctx context.Context, params notification.UpdateRuleInput) (*notification.Rule, error) {
	db := r.client()

	query := db.NotificationRule.UpdateOneID(params.ID).
		SetUpdatedAt(clock.Now().UTC()).
		SetDisabled(params.Disabled).
		SetConfig(params.Config).
		SetName(params.Name).
		ClearChannels().
		AddChannelIDs(params.Channels...)

	queryRow, err := query.Save(ctx)
	if err != nil {
		if entdb.IsNotFound(err) {
			return nil, notification.NotFoundError{
				NamespacedID: models.NamespacedID{
					Namespace: params.Namespace,
					ID:        params.ID,
				},
			}
		}

		return nil, fmt.Errorf("failed to update notification rule: %w", err)
	}

	if queryRow == nil {
		return nil, fmt.Errorf("invalid query result: nil notification rule received")
	}

	channelsQuery := db.NotificationChannel.Query().
		Where(channeldb.Namespace(params.Namespace)).
		Where(channeldb.IDIn(params.Channels...))

	channelRows, err := channelsQuery.All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to query notification channels: %w", err)
	}

	queryRow.Edges.Channels = channelRows

	return RuleFromDBEntity(*queryRow), nil
}
