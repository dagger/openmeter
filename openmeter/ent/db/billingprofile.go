// Code generated by ent, DO NOT EDIT.

package db

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/openmeterio/openmeter/openmeter/billing/provider"
	"github.com/openmeterio/openmeter/openmeter/ent/db/billingprofile"
	"github.com/openmeterio/openmeter/openmeter/ent/db/billingworkflowconfig"
)

// BillingProfile is the model entity for the BillingProfile schema.
type BillingProfile struct {
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
	// Key holds the value of the "key" field.
	Key string `json:"key,omitempty"`
	// ProviderConfig holds the value of the "provider_config" field.
	ProviderConfig provider.Configuration `json:"provider_config,omitempty"`
	// WorkflowConfigID holds the value of the "workflow_config_id" field.
	WorkflowConfigID string `json:"workflow_config_id,omitempty"`
	// Default holds the value of the "default" field.
	Default bool `json:"default,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BillingProfileQuery when eager-loading is set.
	Edges        BillingProfileEdges `json:"edges"`
	selectValues sql.SelectValues
}

// BillingProfileEdges holds the relations/edges for other nodes in the graph.
type BillingProfileEdges struct {
	// BillingInvoices holds the value of the billing_invoices edge.
	BillingInvoices []*BillingInvoice `json:"billing_invoices,omitempty"`
	// BillingWorkflowConfig holds the value of the billing_workflow_config edge.
	BillingWorkflowConfig *BillingWorkflowConfig `json:"billing_workflow_config,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// BillingInvoicesOrErr returns the BillingInvoices value or an error if the edge
// was not loaded in eager-loading.
func (e BillingProfileEdges) BillingInvoicesOrErr() ([]*BillingInvoice, error) {
	if e.loadedTypes[0] {
		return e.BillingInvoices, nil
	}
	return nil, &NotLoadedError{edge: "billing_invoices"}
}

// BillingWorkflowConfigOrErr returns the BillingWorkflowConfig value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BillingProfileEdges) BillingWorkflowConfigOrErr() (*BillingWorkflowConfig, error) {
	if e.BillingWorkflowConfig != nil {
		return e.BillingWorkflowConfig, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: billingworkflowconfig.Label}
	}
	return nil, &NotLoadedError{edge: "billing_workflow_config"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*BillingProfile) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case billingprofile.FieldDefault:
			values[i] = new(sql.NullBool)
		case billingprofile.FieldID, billingprofile.FieldNamespace, billingprofile.FieldKey, billingprofile.FieldWorkflowConfigID:
			values[i] = new(sql.NullString)
		case billingprofile.FieldCreatedAt, billingprofile.FieldUpdatedAt, billingprofile.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case billingprofile.FieldProviderConfig:
			values[i] = billingprofile.ValueScanner.ProviderConfig.ScanValue()
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BillingProfile fields.
func (bp *BillingProfile) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case billingprofile.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				bp.ID = value.String
			}
		case billingprofile.FieldNamespace:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field namespace", values[i])
			} else if value.Valid {
				bp.Namespace = value.String
			}
		case billingprofile.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				bp.CreatedAt = value.Time
			}
		case billingprofile.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				bp.UpdatedAt = value.Time
			}
		case billingprofile.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				bp.DeletedAt = new(time.Time)
				*bp.DeletedAt = value.Time
			}
		case billingprofile.FieldKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field key", values[i])
			} else if value.Valid {
				bp.Key = value.String
			}
		case billingprofile.FieldProviderConfig:
			if value, err := billingprofile.ValueScanner.ProviderConfig.FromValue(values[i]); err != nil {
				return err
			} else {
				bp.ProviderConfig = value
			}
		case billingprofile.FieldWorkflowConfigID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field workflow_config_id", values[i])
			} else if value.Valid {
				bp.WorkflowConfigID = value.String
			}
		case billingprofile.FieldDefault:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field default", values[i])
			} else if value.Valid {
				bp.Default = value.Bool
			}
		default:
			bp.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the BillingProfile.
// This includes values selected through modifiers, order, etc.
func (bp *BillingProfile) Value(name string) (ent.Value, error) {
	return bp.selectValues.Get(name)
}

// QueryBillingInvoices queries the "billing_invoices" edge of the BillingProfile entity.
func (bp *BillingProfile) QueryBillingInvoices() *BillingInvoiceQuery {
	return NewBillingProfileClient(bp.config).QueryBillingInvoices(bp)
}

// QueryBillingWorkflowConfig queries the "billing_workflow_config" edge of the BillingProfile entity.
func (bp *BillingProfile) QueryBillingWorkflowConfig() *BillingWorkflowConfigQuery {
	return NewBillingProfileClient(bp.config).QueryBillingWorkflowConfig(bp)
}

// Update returns a builder for updating this BillingProfile.
// Note that you need to call BillingProfile.Unwrap() before calling this method if this BillingProfile
// was returned from a transaction, and the transaction was committed or rolled back.
func (bp *BillingProfile) Update() *BillingProfileUpdateOne {
	return NewBillingProfileClient(bp.config).UpdateOne(bp)
}

// Unwrap unwraps the BillingProfile entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (bp *BillingProfile) Unwrap() *BillingProfile {
	_tx, ok := bp.config.driver.(*txDriver)
	if !ok {
		panic("db: BillingProfile is not a transactional entity")
	}
	bp.config.driver = _tx.drv
	return bp
}

// String implements the fmt.Stringer.
func (bp *BillingProfile) String() string {
	var builder strings.Builder
	builder.WriteString("BillingProfile(")
	builder.WriteString(fmt.Sprintf("id=%v, ", bp.ID))
	builder.WriteString("namespace=")
	builder.WriteString(bp.Namespace)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(bp.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(bp.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := bp.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("key=")
	builder.WriteString(bp.Key)
	builder.WriteString(", ")
	builder.WriteString("provider_config=")
	builder.WriteString(fmt.Sprintf("%v", bp.ProviderConfig))
	builder.WriteString(", ")
	builder.WriteString("workflow_config_id=")
	builder.WriteString(bp.WorkflowConfigID)
	builder.WriteString(", ")
	builder.WriteString("default=")
	builder.WriteString(fmt.Sprintf("%v", bp.Default))
	builder.WriteByte(')')
	return builder.String()
}

// BillingProfiles is a parsable slice of BillingProfile.
type BillingProfiles []*BillingProfile
