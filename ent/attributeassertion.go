// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/in-toto/archivista/ent/attributeassertion"
	"github.com/in-toto/archivista/ent/attributereport"
)

// AttributeAssertion is the model entity for the AttributeAssertion schema.
type AttributeAssertion struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Attribute holds the value of the "attribute" field.
	Attribute string `json:"attribute,omitempty"`
	// Target holds the value of the "target" field.
	Target map[string]interface{} `json:"target,omitempty"`
	// Conditions holds the value of the "conditions" field.
	Conditions map[string]interface{} `json:"conditions,omitempty"`
	// Evidence holds the value of the "evidence" field.
	Evidence map[string]interface{} `json:"evidence,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AttributeAssertionQuery when eager-loading is set.
	Edges                       AttributeAssertionEdges `json:"edges"`
	attribute_report_attributes *int
	selectValues                sql.SelectValues
}

// AttributeAssertionEdges holds the relations/edges for other nodes in the graph.
type AttributeAssertionEdges struct {
	// AttributesReport holds the value of the attributes_report edge.
	AttributesReport *AttributeReport `json:"attributes_report,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int
}

// AttributesReportOrErr returns the AttributesReport value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AttributeAssertionEdges) AttributesReportOrErr() (*AttributeReport, error) {
	if e.loadedTypes[0] {
		if e.AttributesReport == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: attributereport.Label}
		}
		return e.AttributesReport, nil
	}
	return nil, &NotLoadedError{edge: "attributes_report"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AttributeAssertion) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case attributeassertion.FieldTarget, attributeassertion.FieldConditions, attributeassertion.FieldEvidence:
			values[i] = new([]byte)
		case attributeassertion.FieldID:
			values[i] = new(sql.NullInt64)
		case attributeassertion.FieldAttribute:
			values[i] = new(sql.NullString)
		case attributeassertion.ForeignKeys[0]: // attribute_report_attributes
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AttributeAssertion fields.
func (aa *AttributeAssertion) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case attributeassertion.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			aa.ID = int(value.Int64)
		case attributeassertion.FieldAttribute:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field attribute", values[i])
			} else if value.Valid {
				aa.Attribute = value.String
			}
		case attributeassertion.FieldTarget:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field target", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &aa.Target); err != nil {
					return fmt.Errorf("unmarshal field target: %w", err)
				}
			}
		case attributeassertion.FieldConditions:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field conditions", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &aa.Conditions); err != nil {
					return fmt.Errorf("unmarshal field conditions: %w", err)
				}
			}
		case attributeassertion.FieldEvidence:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field evidence", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &aa.Evidence); err != nil {
					return fmt.Errorf("unmarshal field evidence: %w", err)
				}
			}
		case attributeassertion.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field attribute_report_attributes", value)
			} else if value.Valid {
				aa.attribute_report_attributes = new(int)
				*aa.attribute_report_attributes = int(value.Int64)
			}
		default:
			aa.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the AttributeAssertion.
// This includes values selected through modifiers, order, etc.
func (aa *AttributeAssertion) Value(name string) (ent.Value, error) {
	return aa.selectValues.Get(name)
}

// QueryAttributesReport queries the "attributes_report" edge of the AttributeAssertion entity.
func (aa *AttributeAssertion) QueryAttributesReport() *AttributeReportQuery {
	return NewAttributeAssertionClient(aa.config).QueryAttributesReport(aa)
}

// Update returns a builder for updating this AttributeAssertion.
// Note that you need to call AttributeAssertion.Unwrap() before calling this method if this AttributeAssertion
// was returned from a transaction, and the transaction was committed or rolled back.
func (aa *AttributeAssertion) Update() *AttributeAssertionUpdateOne {
	return NewAttributeAssertionClient(aa.config).UpdateOne(aa)
}

// Unwrap unwraps the AttributeAssertion entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (aa *AttributeAssertion) Unwrap() *AttributeAssertion {
	_tx, ok := aa.config.driver.(*txDriver)
	if !ok {
		panic("ent: AttributeAssertion is not a transactional entity")
	}
	aa.config.driver = _tx.drv
	return aa
}

// String implements the fmt.Stringer.
func (aa *AttributeAssertion) String() string {
	var builder strings.Builder
	builder.WriteString("AttributeAssertion(")
	builder.WriteString(fmt.Sprintf("id=%v, ", aa.ID))
	builder.WriteString("attribute=")
	builder.WriteString(aa.Attribute)
	builder.WriteString(", ")
	builder.WriteString("target=")
	builder.WriteString(fmt.Sprintf("%v", aa.Target))
	builder.WriteString(", ")
	builder.WriteString("conditions=")
	builder.WriteString(fmt.Sprintf("%v", aa.Conditions))
	builder.WriteString(", ")
	builder.WriteString("evidence=")
	builder.WriteString(fmt.Sprintf("%v", aa.Evidence))
	builder.WriteByte(')')
	return builder.String()
}

// AttributeAssertions is a parsable slice of AttributeAssertion.
type AttributeAssertions []*AttributeAssertion
