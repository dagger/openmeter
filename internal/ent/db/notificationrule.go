// Code generated by ent, DO NOT EDIT.

package db

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/openmeterio/openmeter/internal/ent/db/notificationrule"
	"github.com/openmeterio/openmeter/internal/notification"
)

// NotificationRule is the model entity for the NotificationRule schema.
type NotificationRule struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Namespace holds the value of the "namespace" field.
	Namespace string `json:"namespace,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// The event type the rule associated with
	Type notification.RuleType `json:"type,omitempty"`
	// The name of the rule
	Name string `json:"name,omitempty"`
	// Whether the rule is disabled or not
	Disabled bool `json:"disabled,omitempty"`
	// Config holds the value of the "config" field.
	Config notification.RuleConfig `json:"config,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the NotificationRuleQuery when eager-loading is set.
	Edges        NotificationRuleEdges `json:"edges"`
	selectValues sql.SelectValues
}

// NotificationRuleEdges holds the relations/edges for other nodes in the graph.
type NotificationRuleEdges struct {
	// Channels holds the value of the channels edge.
	Channels []*NotificationChannel `json:"channels,omitempty"`
	// Events holds the value of the events edge.
	Events []*NotificationEvent `json:"events,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ChannelsOrErr returns the Channels value or an error if the edge
// was not loaded in eager-loading.
func (e NotificationRuleEdges) ChannelsOrErr() ([]*NotificationChannel, error) {
	if e.loadedTypes[0] {
		return e.Channels, nil
	}
	return nil, &NotLoadedError{edge: "channels"}
}

// EventsOrErr returns the Events value or an error if the edge
// was not loaded in eager-loading.
func (e NotificationRuleEdges) EventsOrErr() ([]*NotificationEvent, error) {
	if e.loadedTypes[1] {
		return e.Events, nil
	}
	return nil, &NotLoadedError{edge: "events"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*NotificationRule) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case notificationrule.FieldDisabled:
			values[i] = new(sql.NullBool)
		case notificationrule.FieldID, notificationrule.FieldNamespace, notificationrule.FieldType, notificationrule.FieldName:
			values[i] = new(sql.NullString)
		case notificationrule.FieldCreatedAt, notificationrule.FieldUpdatedAt, notificationrule.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case notificationrule.FieldConfig:
			values[i] = notificationrule.ValueScanner.Config.ScanValue()
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the NotificationRule fields.
func (nr *NotificationRule) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case notificationrule.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				nr.ID = value.String
			}
		case notificationrule.FieldNamespace:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field namespace", values[i])
			} else if value.Valid {
				nr.Namespace = value.String
			}
		case notificationrule.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				nr.CreatedAt = value.Time
			}
		case notificationrule.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				nr.UpdatedAt = value.Time
			}
		case notificationrule.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				nr.DeletedAt = new(time.Time)
				*nr.DeletedAt = value.Time
			}
		case notificationrule.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				nr.Type = notification.RuleType(value.String)
			}
		case notificationrule.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				nr.Name = value.String
			}
		case notificationrule.FieldDisabled:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field disabled", values[i])
			} else if value.Valid {
				nr.Disabled = value.Bool
			}
		case notificationrule.FieldConfig:
			if value, err := notificationrule.ValueScanner.Config.FromValue(values[i]); err != nil {
				return err
			} else {
				nr.Config = value
			}
		default:
			nr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the NotificationRule.
// This includes values selected through modifiers, order, etc.
func (nr *NotificationRule) Value(name string) (ent.Value, error) {
	return nr.selectValues.Get(name)
}

// QueryChannels queries the "channels" edge of the NotificationRule entity.
func (nr *NotificationRule) QueryChannels() *NotificationChannelQuery {
	return NewNotificationRuleClient(nr.config).QueryChannels(nr)
}

// QueryEvents queries the "events" edge of the NotificationRule entity.
func (nr *NotificationRule) QueryEvents() *NotificationEventQuery {
	return NewNotificationRuleClient(nr.config).QueryEvents(nr)
}

// Update returns a builder for updating this NotificationRule.
// Note that you need to call NotificationRule.Unwrap() before calling this method if this NotificationRule
// was returned from a transaction, and the transaction was committed or rolled back.
func (nr *NotificationRule) Update() *NotificationRuleUpdateOne {
	return NewNotificationRuleClient(nr.config).UpdateOne(nr)
}

// Unwrap unwraps the NotificationRule entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (nr *NotificationRule) Unwrap() *NotificationRule {
	_tx, ok := nr.config.driver.(*txDriver)
	if !ok {
		panic("db: NotificationRule is not a transactional entity")
	}
	nr.config.driver = _tx.drv
	return nr
}

// String implements the fmt.Stringer.
func (nr *NotificationRule) String() string {
	var builder strings.Builder
	builder.WriteString("NotificationRule(")
	builder.WriteString(fmt.Sprintf("id=%v, ", nr.ID))
	builder.WriteString("namespace=")
	builder.WriteString(nr.Namespace)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(nr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(nr.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := nr.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", nr.Type))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(nr.Name)
	builder.WriteString(", ")
	builder.WriteString("disabled=")
	builder.WriteString(fmt.Sprintf("%v", nr.Disabled))
	builder.WriteString(", ")
	builder.WriteString("config=")
	builder.WriteString(fmt.Sprintf("%v", nr.Config))
	builder.WriteByte(')')
	return builder.String()
}

// NotificationRules is a parsable slice of NotificationRule.
type NotificationRules []*NotificationRule
