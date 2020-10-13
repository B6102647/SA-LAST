// Code generated by entc, DO NOT EDIT.

package role

const (
	// Label holds the string label denoting the role type in the database.
	Label = "role"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldROLEID holds the string denoting the role_id field in the database.
	FieldROLEID = "role_id"
	// FieldROLENAME holds the string denoting the role_name field in the database.
	FieldROLENAME = "role_name"

	// EdgeRole holds the string denoting the role edge name in mutations.
	EdgeRole = "Role"

	// Table holds the table name of the role in the database.
	Table = "roles"
	// RoleTable is the table the holds the Role relation/edge.
	RoleTable = "users"
	// RoleInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	RoleInverseTable = "users"
	// RoleColumn is the table column denoting the Role relation/edge.
	RoleColumn = "role_role"
)

// Columns holds all SQL columns for role fields.
var Columns = []string{
	FieldID,
	FieldROLEID,
	FieldROLENAME,
}

var (
	// ROLEIDValidator is a validator for the "ROLE_ID" field. It is called by the builders before save.
	ROLEIDValidator func(int) error
	// ROLENAMEValidator is a validator for the "ROLE_NAME" field. It is called by the builders before save.
	ROLENAMEValidator func(string) error
)
