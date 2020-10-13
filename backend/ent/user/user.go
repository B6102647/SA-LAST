// Code generated by entc, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUSERID holds the string denoting the user_id field in the database.
	FieldUSERID = "user_id"
	// FieldUSEREMAIL holds the string denoting the user_email field in the database.
	FieldUSEREMAIL = "user_email"
	// FieldUSERNAME holds the string denoting the user_name field in the database.
	FieldUSERNAME = "user_name"

	// EdgeBooklist holds the string denoting the booklist edge name in mutations.
	EdgeBooklist = "Booklist"
	// EdgeRole holds the string denoting the role edge name in mutations.
	EdgeRole = "Role"

	// Table holds the table name of the user in the database.
	Table = "users"
	// BooklistTable is the table the holds the Booklist relation/edge.
	BooklistTable = "book_borrows"
	// BooklistInverseTable is the table name for the BookBorrow entity.
	// It exists in this package in order to avoid circular dependency with the "bookborrow" package.
	BooklistInverseTable = "book_borrows"
	// BooklistColumn is the table column denoting the Booklist relation/edge.
	BooklistColumn = "user_booklist"
	// RoleTable is the table the holds the Role relation/edge.
	RoleTable = "users"
	// RoleInverseTable is the table name for the Role entity.
	// It exists in this package in order to avoid circular dependency with the "role" package.
	RoleInverseTable = "roles"
	// RoleColumn is the table column denoting the Role relation/edge.
	RoleColumn = "role_role"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUSERID,
	FieldUSEREMAIL,
	FieldUSERNAME,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the User type.
var ForeignKeys = []string{
	"role_role",
}

var (
	// USERIDValidator is a validator for the "USER_ID" field. It is called by the builders before save.
	USERIDValidator func(int) error
	// USEREMAILValidator is a validator for the "USER_EMAIL" field. It is called by the builders before save.
	USEREMAILValidator func(string) error
	// USERNAMEValidator is a validator for the "USER_NAME" field. It is called by the builders before save.
	USERNAMEValidator func(string) error
)