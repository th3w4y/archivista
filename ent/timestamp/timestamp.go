// Code generated by ent, DO NOT EDIT.

package timestamp

const (
	// Label holds the string label denoting the timestamp type in the database.
	Label = "timestamp"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldTimestamp holds the string denoting the timestamp field in the database.
	FieldTimestamp = "timestamp"
	// EdgeSignature holds the string denoting the signature edge name in mutations.
	EdgeSignature = "signature"
	// Table holds the table name of the timestamp in the database.
	Table = "timestamps"
	// SignatureTable is the table that holds the signature relation/edge.
	SignatureTable = "timestamps"
	// SignatureInverseTable is the table name for the Signature entity.
	// It exists in this package in order to avoid circular dependency with the "signature" package.
	SignatureInverseTable = "signatures"
	// SignatureColumn is the table column denoting the signature relation/edge.
	SignatureColumn = "signature_timestamps"
)

// Columns holds all SQL columns for timestamp fields.
var Columns = []string{
	FieldID,
	FieldType,
	FieldTimestamp,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "timestamps"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"signature_timestamps",
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