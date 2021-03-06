// Code generated by entc, DO NOT EDIT.

package book

const (
	// Label holds the string label denoting the book type in the database.
	Label = "book"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldBOOKNAME holds the string denoting the book_name field in the database.
	FieldBOOKNAME = "book_name"
	// FieldUSERNAME holds the string denoting the user_name field in the database.
	FieldUSERNAME = "user_name"
	// FieldCATEGORY holds the string denoting the category field in the database.
	FieldCATEGORY = "category"
	// FieldAuthor holds the string denoting the author field in the database.
	FieldAuthor = "author"

	// EdgeBooklist holds the string denoting the booklist edge name in mutations.
	EdgeBooklist = "Booklist"
	// EdgeStatus holds the string denoting the status edge name in mutations.
	EdgeStatus = "Status"

	// Table holds the table name of the book in the database.
	Table = "books"
	// BooklistTable is the table the holds the Booklist relation/edge.
	BooklistTable = "book_borrows"
	// BooklistInverseTable is the table name for the BookBorrow entity.
	// It exists in this package in order to avoid circular dependency with the "bookborrow" package.
	BooklistInverseTable = "book_borrows"
	// BooklistColumn is the table column denoting the Booklist relation/edge.
	BooklistColumn = "BOOK_ID"
	// StatusTable is the table the holds the Status relation/edge.
	StatusTable = "books"
	// StatusInverseTable is the table name for the Status entity.
	// It exists in this package in order to avoid circular dependency with the "status" package.
	StatusInverseTable = "status"
	// StatusColumn is the table column denoting the Status relation/edge.
	StatusColumn = "STATUS_ID"
)

// Columns holds all SQL columns for book fields.
var Columns = []string{
	FieldID,
	FieldBOOKNAME,
	FieldUSERNAME,
	FieldCATEGORY,
	FieldAuthor,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Book type.
var ForeignKeys = []string{
	"STATUS_ID",
}

var (
	// BOOKNAMEValidator is a validator for the "BOOK_NAME" field. It is called by the builders before save.
	BOOKNAMEValidator func(string) error
)
