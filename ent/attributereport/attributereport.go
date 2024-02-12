// Code generated by ent, DO NOT EDIT.

package attributereport

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the attributereport type in the database.
	Label = "attribute_report"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// EdgeAttributes holds the string denoting the attributes edge name in mutations.
	EdgeAttributes = "attributes"
	// EdgeStatement holds the string denoting the statement edge name in mutations.
	EdgeStatement = "statement"
	// Table holds the table name of the attributereport in the database.
	Table = "attribute_reports"
	// AttributesTable is the table that holds the attributes relation/edge.
	AttributesTable = "attribute_assertions"
	// AttributesInverseTable is the table name for the AttributeAssertion entity.
	// It exists in this package in order to avoid circular dependency with the "attributeassertion" package.
	AttributesInverseTable = "attribute_assertions"
	// AttributesColumn is the table column denoting the attributes relation/edge.
	AttributesColumn = "attribute_report_attributes"
	// StatementTable is the table that holds the statement relation/edge.
	StatementTable = "attribute_reports"
	// StatementInverseTable is the table name for the Statement entity.
	// It exists in this package in order to avoid circular dependency with the "statement" package.
	StatementInverseTable = "statements"
	// StatementColumn is the table column denoting the statement relation/edge.
	StatementColumn = "statement_attributes_report"
)

// Columns holds all SQL columns for attributereport fields.
var Columns = []string{
	FieldID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "attribute_reports"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"statement_attributes_report",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the AttributeReport queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByAttributesCount orders the results by attributes count.
func ByAttributesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAttributesStep(), opts...)
	}
}

// ByAttributes orders the results by attributes terms.
func ByAttributes(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAttributesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByStatementField orders the results by statement field.
func ByStatementField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newStatementStep(), sql.OrderByField(field, opts...))
	}
}
func newAttributesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AttributesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, AttributesTable, AttributesColumn),
	)
}
func newStatementStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(StatementInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, StatementTable, StatementColumn),
	)
}