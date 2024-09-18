// Code generated by ent, DO NOT EDIT.

package billingprofile

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/openmeterio/openmeter/openmeter/ent/db/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldContainsFold(FieldID, id))
}

// Namespace applies equality check predicate on the "namespace" field. It's identical to NamespaceEQ.
func Namespace(v string) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldEQ(FieldNamespace, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldEQ(FieldDeletedAt, v))
}

// Key applies equality check predicate on the "key" field. It's identical to KeyEQ.
func Key(v string) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldEQ(FieldKey, v))
}

// WorkflowConfigID applies equality check predicate on the "workflow_config_id" field. It's identical to WorkflowConfigIDEQ.
func WorkflowConfigID(v string) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldEQ(FieldWorkflowConfigID, v))
}

// Default applies equality check predicate on the "default" field. It's identical to DefaultEQ.
func Default(v bool) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldEQ(FieldDefault, v))
}

// NamespaceEQ applies the EQ predicate on the "namespace" field.
func NamespaceEQ(v string) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldEQ(FieldNamespace, v))
}

// NamespaceNEQ applies the NEQ predicate on the "namespace" field.
func NamespaceNEQ(v string) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldNEQ(FieldNamespace, v))
}

// NamespaceIn applies the In predicate on the "namespace" field.
func NamespaceIn(vs ...string) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldIn(FieldNamespace, vs...))
}

// NamespaceNotIn applies the NotIn predicate on the "namespace" field.
func NamespaceNotIn(vs ...string) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldNotIn(FieldNamespace, vs...))
}

// NamespaceGT applies the GT predicate on the "namespace" field.
func NamespaceGT(v string) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldGT(FieldNamespace, v))
}

// NamespaceGTE applies the GTE predicate on the "namespace" field.
func NamespaceGTE(v string) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldGTE(FieldNamespace, v))
}

// NamespaceLT applies the LT predicate on the "namespace" field.
func NamespaceLT(v string) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldLT(FieldNamespace, v))
}

// NamespaceLTE applies the LTE predicate on the "namespace" field.
func NamespaceLTE(v string) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldLTE(FieldNamespace, v))
}

// NamespaceContains applies the Contains predicate on the "namespace" field.
func NamespaceContains(v string) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldContains(FieldNamespace, v))
}

// NamespaceHasPrefix applies the HasPrefix predicate on the "namespace" field.
func NamespaceHasPrefix(v string) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldHasPrefix(FieldNamespace, v))
}

// NamespaceHasSuffix applies the HasSuffix predicate on the "namespace" field.
func NamespaceHasSuffix(v string) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldHasSuffix(FieldNamespace, v))
}

// NamespaceEqualFold applies the EqualFold predicate on the "namespace" field.
func NamespaceEqualFold(v string) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldEqualFold(FieldNamespace, v))
}

// NamespaceContainsFold applies the ContainsFold predicate on the "namespace" field.
func NamespaceContainsFold(v string) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldContainsFold(FieldNamespace, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldNotNull(FieldDeletedAt))
}

// KeyEQ applies the EQ predicate on the "key" field.
func KeyEQ(v string) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldEQ(FieldKey, v))
}

// KeyNEQ applies the NEQ predicate on the "key" field.
func KeyNEQ(v string) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldNEQ(FieldKey, v))
}

// KeyIn applies the In predicate on the "key" field.
func KeyIn(vs ...string) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldIn(FieldKey, vs...))
}

// KeyNotIn applies the NotIn predicate on the "key" field.
func KeyNotIn(vs ...string) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldNotIn(FieldKey, vs...))
}

// KeyGT applies the GT predicate on the "key" field.
func KeyGT(v string) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldGT(FieldKey, v))
}

// KeyGTE applies the GTE predicate on the "key" field.
func KeyGTE(v string) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldGTE(FieldKey, v))
}

// KeyLT applies the LT predicate on the "key" field.
func KeyLT(v string) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldLT(FieldKey, v))
}

// KeyLTE applies the LTE predicate on the "key" field.
func KeyLTE(v string) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldLTE(FieldKey, v))
}

// KeyContains applies the Contains predicate on the "key" field.
func KeyContains(v string) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldContains(FieldKey, v))
}

// KeyHasPrefix applies the HasPrefix predicate on the "key" field.
func KeyHasPrefix(v string) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldHasPrefix(FieldKey, v))
}

// KeyHasSuffix applies the HasSuffix predicate on the "key" field.
func KeyHasSuffix(v string) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldHasSuffix(FieldKey, v))
}

// KeyEqualFold applies the EqualFold predicate on the "key" field.
func KeyEqualFold(v string) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldEqualFold(FieldKey, v))
}

// KeyContainsFold applies the ContainsFold predicate on the "key" field.
func KeyContainsFold(v string) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldContainsFold(FieldKey, v))
}

// WorkflowConfigIDEQ applies the EQ predicate on the "workflow_config_id" field.
func WorkflowConfigIDEQ(v string) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldEQ(FieldWorkflowConfigID, v))
}

// WorkflowConfigIDNEQ applies the NEQ predicate on the "workflow_config_id" field.
func WorkflowConfigIDNEQ(v string) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldNEQ(FieldWorkflowConfigID, v))
}

// WorkflowConfigIDIn applies the In predicate on the "workflow_config_id" field.
func WorkflowConfigIDIn(vs ...string) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldIn(FieldWorkflowConfigID, vs...))
}

// WorkflowConfigIDNotIn applies the NotIn predicate on the "workflow_config_id" field.
func WorkflowConfigIDNotIn(vs ...string) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldNotIn(FieldWorkflowConfigID, vs...))
}

// WorkflowConfigIDGT applies the GT predicate on the "workflow_config_id" field.
func WorkflowConfigIDGT(v string) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldGT(FieldWorkflowConfigID, v))
}

// WorkflowConfigIDGTE applies the GTE predicate on the "workflow_config_id" field.
func WorkflowConfigIDGTE(v string) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldGTE(FieldWorkflowConfigID, v))
}

// WorkflowConfigIDLT applies the LT predicate on the "workflow_config_id" field.
func WorkflowConfigIDLT(v string) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldLT(FieldWorkflowConfigID, v))
}

// WorkflowConfigIDLTE applies the LTE predicate on the "workflow_config_id" field.
func WorkflowConfigIDLTE(v string) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldLTE(FieldWorkflowConfigID, v))
}

// WorkflowConfigIDContains applies the Contains predicate on the "workflow_config_id" field.
func WorkflowConfigIDContains(v string) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldContains(FieldWorkflowConfigID, v))
}

// WorkflowConfigIDHasPrefix applies the HasPrefix predicate on the "workflow_config_id" field.
func WorkflowConfigIDHasPrefix(v string) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldHasPrefix(FieldWorkflowConfigID, v))
}

// WorkflowConfigIDHasSuffix applies the HasSuffix predicate on the "workflow_config_id" field.
func WorkflowConfigIDHasSuffix(v string) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldHasSuffix(FieldWorkflowConfigID, v))
}

// WorkflowConfigIDEqualFold applies the EqualFold predicate on the "workflow_config_id" field.
func WorkflowConfigIDEqualFold(v string) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldEqualFold(FieldWorkflowConfigID, v))
}

// WorkflowConfigIDContainsFold applies the ContainsFold predicate on the "workflow_config_id" field.
func WorkflowConfigIDContainsFold(v string) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldContainsFold(FieldWorkflowConfigID, v))
}

// DefaultEQ applies the EQ predicate on the "default" field.
func DefaultEQ(v bool) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldEQ(FieldDefault, v))
}

// DefaultNEQ applies the NEQ predicate on the "default" field.
func DefaultNEQ(v bool) predicate.BillingProfile {
	return predicate.BillingProfile(sql.FieldNEQ(FieldDefault, v))
}

// HasBillingInvoices applies the HasEdge predicate on the "billing_invoices" edge.
func HasBillingInvoices() predicate.BillingProfile {
	return predicate.BillingProfile(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BillingInvoicesTable, BillingInvoicesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBillingInvoicesWith applies the HasEdge predicate on the "billing_invoices" edge with a given conditions (other predicates).
func HasBillingInvoicesWith(preds ...predicate.BillingInvoice) predicate.BillingProfile {
	return predicate.BillingProfile(func(s *sql.Selector) {
		step := newBillingInvoicesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBillingWorkflowConfig applies the HasEdge predicate on the "billing_workflow_config" edge.
func HasBillingWorkflowConfig() predicate.BillingProfile {
	return predicate.BillingProfile(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BillingWorkflowConfigTable, BillingWorkflowConfigColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBillingWorkflowConfigWith applies the HasEdge predicate on the "billing_workflow_config" edge with a given conditions (other predicates).
func HasBillingWorkflowConfigWith(preds ...predicate.BillingWorkflowConfig) predicate.BillingProfile {
	return predicate.BillingProfile(func(s *sql.Selector) {
		step := newBillingWorkflowConfigStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.BillingProfile) predicate.BillingProfile {
	return predicate.BillingProfile(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.BillingProfile) predicate.BillingProfile {
	return predicate.BillingProfile(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.BillingProfile) predicate.BillingProfile {
	return predicate.BillingProfile(sql.NotPredicates(p))
}