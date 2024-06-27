// Code generated by ent, DO NOT EDIT.

package grant

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/openmeterio/openmeter/internal/credit"
	"github.com/openmeterio/openmeter/internal/credit/postgresadapter/ent/db/predicate"
	"github.com/openmeterio/openmeter/pkg/recurrence"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Grant {
	return predicate.Grant(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Grant {
	return predicate.Grant(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Grant {
	return predicate.Grant(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Grant {
	return predicate.Grant(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Grant {
	return predicate.Grant(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Grant {
	return predicate.Grant(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Grant {
	return predicate.Grant(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Grant {
	return predicate.Grant(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Grant {
	return predicate.Grant(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Grant {
	return predicate.Grant(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Grant {
	return predicate.Grant(sql.FieldContainsFold(FieldID, id))
}

// Namespace applies equality check predicate on the "namespace" field. It's identical to NamespaceEQ.
func Namespace(v string) predicate.Grant {
	return predicate.Grant(sql.FieldEQ(FieldNamespace, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldEQ(FieldDeletedAt, v))
}

// OwnerID applies equality check predicate on the "owner_id" field. It's identical to OwnerIDEQ.
func OwnerID(v credit.GrantOwner) predicate.Grant {
	vc := string(v)
	return predicate.Grant(sql.FieldEQ(FieldOwnerID, vc))
}

// Amount applies equality check predicate on the "amount" field. It's identical to AmountEQ.
func Amount(v float64) predicate.Grant {
	return predicate.Grant(sql.FieldEQ(FieldAmount, v))
}

// Priority applies equality check predicate on the "priority" field. It's identical to PriorityEQ.
func Priority(v uint8) predicate.Grant {
	return predicate.Grant(sql.FieldEQ(FieldPriority, v))
}

// EffectiveAt applies equality check predicate on the "effective_at" field. It's identical to EffectiveAtEQ.
func EffectiveAt(v time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldEQ(FieldEffectiveAt, v))
}

// ExpiresAt applies equality check predicate on the "expires_at" field. It's identical to ExpiresAtEQ.
func ExpiresAt(v time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldEQ(FieldExpiresAt, v))
}

// VoidedAt applies equality check predicate on the "voided_at" field. It's identical to VoidedAtEQ.
func VoidedAt(v time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldEQ(FieldVoidedAt, v))
}

// ResetMaxRollover applies equality check predicate on the "reset_max_rollover" field. It's identical to ResetMaxRolloverEQ.
func ResetMaxRollover(v float64) predicate.Grant {
	return predicate.Grant(sql.FieldEQ(FieldResetMaxRollover, v))
}

// ResetMinRollover applies equality check predicate on the "reset_min_rollover" field. It's identical to ResetMinRolloverEQ.
func ResetMinRollover(v float64) predicate.Grant {
	return predicate.Grant(sql.FieldEQ(FieldResetMinRollover, v))
}

// RecurrenceAnchor applies equality check predicate on the "recurrence_anchor" field. It's identical to RecurrenceAnchorEQ.
func RecurrenceAnchor(v time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldEQ(FieldRecurrenceAnchor, v))
}

// NamespaceEQ applies the EQ predicate on the "namespace" field.
func NamespaceEQ(v string) predicate.Grant {
	return predicate.Grant(sql.FieldEQ(FieldNamespace, v))
}

// NamespaceNEQ applies the NEQ predicate on the "namespace" field.
func NamespaceNEQ(v string) predicate.Grant {
	return predicate.Grant(sql.FieldNEQ(FieldNamespace, v))
}

// NamespaceIn applies the In predicate on the "namespace" field.
func NamespaceIn(vs ...string) predicate.Grant {
	return predicate.Grant(sql.FieldIn(FieldNamespace, vs...))
}

// NamespaceNotIn applies the NotIn predicate on the "namespace" field.
func NamespaceNotIn(vs ...string) predicate.Grant {
	return predicate.Grant(sql.FieldNotIn(FieldNamespace, vs...))
}

// NamespaceGT applies the GT predicate on the "namespace" field.
func NamespaceGT(v string) predicate.Grant {
	return predicate.Grant(sql.FieldGT(FieldNamespace, v))
}

// NamespaceGTE applies the GTE predicate on the "namespace" field.
func NamespaceGTE(v string) predicate.Grant {
	return predicate.Grant(sql.FieldGTE(FieldNamespace, v))
}

// NamespaceLT applies the LT predicate on the "namespace" field.
func NamespaceLT(v string) predicate.Grant {
	return predicate.Grant(sql.FieldLT(FieldNamespace, v))
}

// NamespaceLTE applies the LTE predicate on the "namespace" field.
func NamespaceLTE(v string) predicate.Grant {
	return predicate.Grant(sql.FieldLTE(FieldNamespace, v))
}

// NamespaceContains applies the Contains predicate on the "namespace" field.
func NamespaceContains(v string) predicate.Grant {
	return predicate.Grant(sql.FieldContains(FieldNamespace, v))
}

// NamespaceHasPrefix applies the HasPrefix predicate on the "namespace" field.
func NamespaceHasPrefix(v string) predicate.Grant {
	return predicate.Grant(sql.FieldHasPrefix(FieldNamespace, v))
}

// NamespaceHasSuffix applies the HasSuffix predicate on the "namespace" field.
func NamespaceHasSuffix(v string) predicate.Grant {
	return predicate.Grant(sql.FieldHasSuffix(FieldNamespace, v))
}

// NamespaceEqualFold applies the EqualFold predicate on the "namespace" field.
func NamespaceEqualFold(v string) predicate.Grant {
	return predicate.Grant(sql.FieldEqualFold(FieldNamespace, v))
}

// NamespaceContainsFold applies the ContainsFold predicate on the "namespace" field.
func NamespaceContainsFold(v string) predicate.Grant {
	return predicate.Grant(sql.FieldContainsFold(FieldNamespace, v))
}

// MetadataIsNil applies the IsNil predicate on the "metadata" field.
func MetadataIsNil() predicate.Grant {
	return predicate.Grant(sql.FieldIsNull(FieldMetadata))
}

// MetadataNotNil applies the NotNil predicate on the "metadata" field.
func MetadataNotNil() predicate.Grant {
	return predicate.Grant(sql.FieldNotNull(FieldMetadata))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Grant {
	return predicate.Grant(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Grant {
	return predicate.Grant(sql.FieldNotNull(FieldDeletedAt))
}

// OwnerIDEQ applies the EQ predicate on the "owner_id" field.
func OwnerIDEQ(v credit.GrantOwner) predicate.Grant {
	vc := string(v)
	return predicate.Grant(sql.FieldEQ(FieldOwnerID, vc))
}

// OwnerIDNEQ applies the NEQ predicate on the "owner_id" field.
func OwnerIDNEQ(v credit.GrantOwner) predicate.Grant {
	vc := string(v)
	return predicate.Grant(sql.FieldNEQ(FieldOwnerID, vc))
}

// OwnerIDIn applies the In predicate on the "owner_id" field.
func OwnerIDIn(vs ...credit.GrantOwner) predicate.Grant {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = string(vs[i])
	}
	return predicate.Grant(sql.FieldIn(FieldOwnerID, v...))
}

// OwnerIDNotIn applies the NotIn predicate on the "owner_id" field.
func OwnerIDNotIn(vs ...credit.GrantOwner) predicate.Grant {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = string(vs[i])
	}
	return predicate.Grant(sql.FieldNotIn(FieldOwnerID, v...))
}

// OwnerIDGT applies the GT predicate on the "owner_id" field.
func OwnerIDGT(v credit.GrantOwner) predicate.Grant {
	vc := string(v)
	return predicate.Grant(sql.FieldGT(FieldOwnerID, vc))
}

// OwnerIDGTE applies the GTE predicate on the "owner_id" field.
func OwnerIDGTE(v credit.GrantOwner) predicate.Grant {
	vc := string(v)
	return predicate.Grant(sql.FieldGTE(FieldOwnerID, vc))
}

// OwnerIDLT applies the LT predicate on the "owner_id" field.
func OwnerIDLT(v credit.GrantOwner) predicate.Grant {
	vc := string(v)
	return predicate.Grant(sql.FieldLT(FieldOwnerID, vc))
}

// OwnerIDLTE applies the LTE predicate on the "owner_id" field.
func OwnerIDLTE(v credit.GrantOwner) predicate.Grant {
	vc := string(v)
	return predicate.Grant(sql.FieldLTE(FieldOwnerID, vc))
}

// OwnerIDContains applies the Contains predicate on the "owner_id" field.
func OwnerIDContains(v credit.GrantOwner) predicate.Grant {
	vc := string(v)
	return predicate.Grant(sql.FieldContains(FieldOwnerID, vc))
}

// OwnerIDHasPrefix applies the HasPrefix predicate on the "owner_id" field.
func OwnerIDHasPrefix(v credit.GrantOwner) predicate.Grant {
	vc := string(v)
	return predicate.Grant(sql.FieldHasPrefix(FieldOwnerID, vc))
}

// OwnerIDHasSuffix applies the HasSuffix predicate on the "owner_id" field.
func OwnerIDHasSuffix(v credit.GrantOwner) predicate.Grant {
	vc := string(v)
	return predicate.Grant(sql.FieldHasSuffix(FieldOwnerID, vc))
}

// OwnerIDEqualFold applies the EqualFold predicate on the "owner_id" field.
func OwnerIDEqualFold(v credit.GrantOwner) predicate.Grant {
	vc := string(v)
	return predicate.Grant(sql.FieldEqualFold(FieldOwnerID, vc))
}

// OwnerIDContainsFold applies the ContainsFold predicate on the "owner_id" field.
func OwnerIDContainsFold(v credit.GrantOwner) predicate.Grant {
	vc := string(v)
	return predicate.Grant(sql.FieldContainsFold(FieldOwnerID, vc))
}

// AmountEQ applies the EQ predicate on the "amount" field.
func AmountEQ(v float64) predicate.Grant {
	return predicate.Grant(sql.FieldEQ(FieldAmount, v))
}

// AmountNEQ applies the NEQ predicate on the "amount" field.
func AmountNEQ(v float64) predicate.Grant {
	return predicate.Grant(sql.FieldNEQ(FieldAmount, v))
}

// AmountIn applies the In predicate on the "amount" field.
func AmountIn(vs ...float64) predicate.Grant {
	return predicate.Grant(sql.FieldIn(FieldAmount, vs...))
}

// AmountNotIn applies the NotIn predicate on the "amount" field.
func AmountNotIn(vs ...float64) predicate.Grant {
	return predicate.Grant(sql.FieldNotIn(FieldAmount, vs...))
}

// AmountGT applies the GT predicate on the "amount" field.
func AmountGT(v float64) predicate.Grant {
	return predicate.Grant(sql.FieldGT(FieldAmount, v))
}

// AmountGTE applies the GTE predicate on the "amount" field.
func AmountGTE(v float64) predicate.Grant {
	return predicate.Grant(sql.FieldGTE(FieldAmount, v))
}

// AmountLT applies the LT predicate on the "amount" field.
func AmountLT(v float64) predicate.Grant {
	return predicate.Grant(sql.FieldLT(FieldAmount, v))
}

// AmountLTE applies the LTE predicate on the "amount" field.
func AmountLTE(v float64) predicate.Grant {
	return predicate.Grant(sql.FieldLTE(FieldAmount, v))
}

// PriorityEQ applies the EQ predicate on the "priority" field.
func PriorityEQ(v uint8) predicate.Grant {
	return predicate.Grant(sql.FieldEQ(FieldPriority, v))
}

// PriorityNEQ applies the NEQ predicate on the "priority" field.
func PriorityNEQ(v uint8) predicate.Grant {
	return predicate.Grant(sql.FieldNEQ(FieldPriority, v))
}

// PriorityIn applies the In predicate on the "priority" field.
func PriorityIn(vs ...uint8) predicate.Grant {
	return predicate.Grant(sql.FieldIn(FieldPriority, vs...))
}

// PriorityNotIn applies the NotIn predicate on the "priority" field.
func PriorityNotIn(vs ...uint8) predicate.Grant {
	return predicate.Grant(sql.FieldNotIn(FieldPriority, vs...))
}

// PriorityGT applies the GT predicate on the "priority" field.
func PriorityGT(v uint8) predicate.Grant {
	return predicate.Grant(sql.FieldGT(FieldPriority, v))
}

// PriorityGTE applies the GTE predicate on the "priority" field.
func PriorityGTE(v uint8) predicate.Grant {
	return predicate.Grant(sql.FieldGTE(FieldPriority, v))
}

// PriorityLT applies the LT predicate on the "priority" field.
func PriorityLT(v uint8) predicate.Grant {
	return predicate.Grant(sql.FieldLT(FieldPriority, v))
}

// PriorityLTE applies the LTE predicate on the "priority" field.
func PriorityLTE(v uint8) predicate.Grant {
	return predicate.Grant(sql.FieldLTE(FieldPriority, v))
}

// EffectiveAtEQ applies the EQ predicate on the "effective_at" field.
func EffectiveAtEQ(v time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldEQ(FieldEffectiveAt, v))
}

// EffectiveAtNEQ applies the NEQ predicate on the "effective_at" field.
func EffectiveAtNEQ(v time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldNEQ(FieldEffectiveAt, v))
}

// EffectiveAtIn applies the In predicate on the "effective_at" field.
func EffectiveAtIn(vs ...time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldIn(FieldEffectiveAt, vs...))
}

// EffectiveAtNotIn applies the NotIn predicate on the "effective_at" field.
func EffectiveAtNotIn(vs ...time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldNotIn(FieldEffectiveAt, vs...))
}

// EffectiveAtGT applies the GT predicate on the "effective_at" field.
func EffectiveAtGT(v time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldGT(FieldEffectiveAt, v))
}

// EffectiveAtGTE applies the GTE predicate on the "effective_at" field.
func EffectiveAtGTE(v time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldGTE(FieldEffectiveAt, v))
}

// EffectiveAtLT applies the LT predicate on the "effective_at" field.
func EffectiveAtLT(v time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldLT(FieldEffectiveAt, v))
}

// EffectiveAtLTE applies the LTE predicate on the "effective_at" field.
func EffectiveAtLTE(v time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldLTE(FieldEffectiveAt, v))
}

// ExpiresAtEQ applies the EQ predicate on the "expires_at" field.
func ExpiresAtEQ(v time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldEQ(FieldExpiresAt, v))
}

// ExpiresAtNEQ applies the NEQ predicate on the "expires_at" field.
func ExpiresAtNEQ(v time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldNEQ(FieldExpiresAt, v))
}

// ExpiresAtIn applies the In predicate on the "expires_at" field.
func ExpiresAtIn(vs ...time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldIn(FieldExpiresAt, vs...))
}

// ExpiresAtNotIn applies the NotIn predicate on the "expires_at" field.
func ExpiresAtNotIn(vs ...time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldNotIn(FieldExpiresAt, vs...))
}

// ExpiresAtGT applies the GT predicate on the "expires_at" field.
func ExpiresAtGT(v time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldGT(FieldExpiresAt, v))
}

// ExpiresAtGTE applies the GTE predicate on the "expires_at" field.
func ExpiresAtGTE(v time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldGTE(FieldExpiresAt, v))
}

// ExpiresAtLT applies the LT predicate on the "expires_at" field.
func ExpiresAtLT(v time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldLT(FieldExpiresAt, v))
}

// ExpiresAtLTE applies the LTE predicate on the "expires_at" field.
func ExpiresAtLTE(v time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldLTE(FieldExpiresAt, v))
}

// VoidedAtEQ applies the EQ predicate on the "voided_at" field.
func VoidedAtEQ(v time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldEQ(FieldVoidedAt, v))
}

// VoidedAtNEQ applies the NEQ predicate on the "voided_at" field.
func VoidedAtNEQ(v time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldNEQ(FieldVoidedAt, v))
}

// VoidedAtIn applies the In predicate on the "voided_at" field.
func VoidedAtIn(vs ...time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldIn(FieldVoidedAt, vs...))
}

// VoidedAtNotIn applies the NotIn predicate on the "voided_at" field.
func VoidedAtNotIn(vs ...time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldNotIn(FieldVoidedAt, vs...))
}

// VoidedAtGT applies the GT predicate on the "voided_at" field.
func VoidedAtGT(v time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldGT(FieldVoidedAt, v))
}

// VoidedAtGTE applies the GTE predicate on the "voided_at" field.
func VoidedAtGTE(v time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldGTE(FieldVoidedAt, v))
}

// VoidedAtLT applies the LT predicate on the "voided_at" field.
func VoidedAtLT(v time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldLT(FieldVoidedAt, v))
}

// VoidedAtLTE applies the LTE predicate on the "voided_at" field.
func VoidedAtLTE(v time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldLTE(FieldVoidedAt, v))
}

// VoidedAtIsNil applies the IsNil predicate on the "voided_at" field.
func VoidedAtIsNil() predicate.Grant {
	return predicate.Grant(sql.FieldIsNull(FieldVoidedAt))
}

// VoidedAtNotNil applies the NotNil predicate on the "voided_at" field.
func VoidedAtNotNil() predicate.Grant {
	return predicate.Grant(sql.FieldNotNull(FieldVoidedAt))
}

// ResetMaxRolloverEQ applies the EQ predicate on the "reset_max_rollover" field.
func ResetMaxRolloverEQ(v float64) predicate.Grant {
	return predicate.Grant(sql.FieldEQ(FieldResetMaxRollover, v))
}

// ResetMaxRolloverNEQ applies the NEQ predicate on the "reset_max_rollover" field.
func ResetMaxRolloverNEQ(v float64) predicate.Grant {
	return predicate.Grant(sql.FieldNEQ(FieldResetMaxRollover, v))
}

// ResetMaxRolloverIn applies the In predicate on the "reset_max_rollover" field.
func ResetMaxRolloverIn(vs ...float64) predicate.Grant {
	return predicate.Grant(sql.FieldIn(FieldResetMaxRollover, vs...))
}

// ResetMaxRolloverNotIn applies the NotIn predicate on the "reset_max_rollover" field.
func ResetMaxRolloverNotIn(vs ...float64) predicate.Grant {
	return predicate.Grant(sql.FieldNotIn(FieldResetMaxRollover, vs...))
}

// ResetMaxRolloverGT applies the GT predicate on the "reset_max_rollover" field.
func ResetMaxRolloverGT(v float64) predicate.Grant {
	return predicate.Grant(sql.FieldGT(FieldResetMaxRollover, v))
}

// ResetMaxRolloverGTE applies the GTE predicate on the "reset_max_rollover" field.
func ResetMaxRolloverGTE(v float64) predicate.Grant {
	return predicate.Grant(sql.FieldGTE(FieldResetMaxRollover, v))
}

// ResetMaxRolloverLT applies the LT predicate on the "reset_max_rollover" field.
func ResetMaxRolloverLT(v float64) predicate.Grant {
	return predicate.Grant(sql.FieldLT(FieldResetMaxRollover, v))
}

// ResetMaxRolloverLTE applies the LTE predicate on the "reset_max_rollover" field.
func ResetMaxRolloverLTE(v float64) predicate.Grant {
	return predicate.Grant(sql.FieldLTE(FieldResetMaxRollover, v))
}

// ResetMinRolloverEQ applies the EQ predicate on the "reset_min_rollover" field.
func ResetMinRolloverEQ(v float64) predicate.Grant {
	return predicate.Grant(sql.FieldEQ(FieldResetMinRollover, v))
}

// ResetMinRolloverNEQ applies the NEQ predicate on the "reset_min_rollover" field.
func ResetMinRolloverNEQ(v float64) predicate.Grant {
	return predicate.Grant(sql.FieldNEQ(FieldResetMinRollover, v))
}

// ResetMinRolloverIn applies the In predicate on the "reset_min_rollover" field.
func ResetMinRolloverIn(vs ...float64) predicate.Grant {
	return predicate.Grant(sql.FieldIn(FieldResetMinRollover, vs...))
}

// ResetMinRolloverNotIn applies the NotIn predicate on the "reset_min_rollover" field.
func ResetMinRolloverNotIn(vs ...float64) predicate.Grant {
	return predicate.Grant(sql.FieldNotIn(FieldResetMinRollover, vs...))
}

// ResetMinRolloverGT applies the GT predicate on the "reset_min_rollover" field.
func ResetMinRolloverGT(v float64) predicate.Grant {
	return predicate.Grant(sql.FieldGT(FieldResetMinRollover, v))
}

// ResetMinRolloverGTE applies the GTE predicate on the "reset_min_rollover" field.
func ResetMinRolloverGTE(v float64) predicate.Grant {
	return predicate.Grant(sql.FieldGTE(FieldResetMinRollover, v))
}

// ResetMinRolloverLT applies the LT predicate on the "reset_min_rollover" field.
func ResetMinRolloverLT(v float64) predicate.Grant {
	return predicate.Grant(sql.FieldLT(FieldResetMinRollover, v))
}

// ResetMinRolloverLTE applies the LTE predicate on the "reset_min_rollover" field.
func ResetMinRolloverLTE(v float64) predicate.Grant {
	return predicate.Grant(sql.FieldLTE(FieldResetMinRollover, v))
}

// RecurrencePeriodEQ applies the EQ predicate on the "recurrence_period" field.
func RecurrencePeriodEQ(v recurrence.RecurrenceInterval) predicate.Grant {
	vc := v
	return predicate.Grant(sql.FieldEQ(FieldRecurrencePeriod, vc))
}

// RecurrencePeriodNEQ applies the NEQ predicate on the "recurrence_period" field.
func RecurrencePeriodNEQ(v recurrence.RecurrenceInterval) predicate.Grant {
	vc := v
	return predicate.Grant(sql.FieldNEQ(FieldRecurrencePeriod, vc))
}

// RecurrencePeriodIn applies the In predicate on the "recurrence_period" field.
func RecurrencePeriodIn(vs ...recurrence.RecurrenceInterval) predicate.Grant {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Grant(sql.FieldIn(FieldRecurrencePeriod, v...))
}

// RecurrencePeriodNotIn applies the NotIn predicate on the "recurrence_period" field.
func RecurrencePeriodNotIn(vs ...recurrence.RecurrenceInterval) predicate.Grant {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Grant(sql.FieldNotIn(FieldRecurrencePeriod, v...))
}

// RecurrencePeriodIsNil applies the IsNil predicate on the "recurrence_period" field.
func RecurrencePeriodIsNil() predicate.Grant {
	return predicate.Grant(sql.FieldIsNull(FieldRecurrencePeriod))
}

// RecurrencePeriodNotNil applies the NotNil predicate on the "recurrence_period" field.
func RecurrencePeriodNotNil() predicate.Grant {
	return predicate.Grant(sql.FieldNotNull(FieldRecurrencePeriod))
}

// RecurrenceAnchorEQ applies the EQ predicate on the "recurrence_anchor" field.
func RecurrenceAnchorEQ(v time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldEQ(FieldRecurrenceAnchor, v))
}

// RecurrenceAnchorNEQ applies the NEQ predicate on the "recurrence_anchor" field.
func RecurrenceAnchorNEQ(v time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldNEQ(FieldRecurrenceAnchor, v))
}

// RecurrenceAnchorIn applies the In predicate on the "recurrence_anchor" field.
func RecurrenceAnchorIn(vs ...time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldIn(FieldRecurrenceAnchor, vs...))
}

// RecurrenceAnchorNotIn applies the NotIn predicate on the "recurrence_anchor" field.
func RecurrenceAnchorNotIn(vs ...time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldNotIn(FieldRecurrenceAnchor, vs...))
}

// RecurrenceAnchorGT applies the GT predicate on the "recurrence_anchor" field.
func RecurrenceAnchorGT(v time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldGT(FieldRecurrenceAnchor, v))
}

// RecurrenceAnchorGTE applies the GTE predicate on the "recurrence_anchor" field.
func RecurrenceAnchorGTE(v time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldGTE(FieldRecurrenceAnchor, v))
}

// RecurrenceAnchorLT applies the LT predicate on the "recurrence_anchor" field.
func RecurrenceAnchorLT(v time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldLT(FieldRecurrenceAnchor, v))
}

// RecurrenceAnchorLTE applies the LTE predicate on the "recurrence_anchor" field.
func RecurrenceAnchorLTE(v time.Time) predicate.Grant {
	return predicate.Grant(sql.FieldLTE(FieldRecurrenceAnchor, v))
}

// RecurrenceAnchorIsNil applies the IsNil predicate on the "recurrence_anchor" field.
func RecurrenceAnchorIsNil() predicate.Grant {
	return predicate.Grant(sql.FieldIsNull(FieldRecurrenceAnchor))
}

// RecurrenceAnchorNotNil applies the NotNil predicate on the "recurrence_anchor" field.
func RecurrenceAnchorNotNil() predicate.Grant {
	return predicate.Grant(sql.FieldNotNull(FieldRecurrenceAnchor))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Grant) predicate.Grant {
	return predicate.Grant(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Grant) predicate.Grant {
	return predicate.Grant(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Grant) predicate.Grant {
	return predicate.Grant(sql.NotPredicates(p))
}
