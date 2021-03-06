// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// BooksColumns holds the columns for the "books" table.
	BooksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "book_name", Type: field.TypeString, Unique: true},
		{Name: "user_name", Type: field.TypeString},
		{Name: "category", Type: field.TypeString},
		{Name: "author", Type: field.TypeString},
		{Name: "STATUS_ID", Type: field.TypeInt, Nullable: true},
	}
	// BooksTable holds the schema information for the "books" table.
	BooksTable = &schema.Table{
		Name:       "books",
		Columns:    BooksColumns,
		PrimaryKey: []*schema.Column{BooksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "books_status_status",
				Columns: []*schema.Column{BooksColumns[5]},

				RefColumns: []*schema.Column{StatusColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// BookBorrowsColumns holds the columns for the "book_borrows" table.
	BookBorrowsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "added_time", Type: field.TypeTime},
		{Name: "BOOK_ID", Type: field.TypeInt, Nullable: true},
		{Name: "PURPOSE_ID", Type: field.TypeInt, Nullable: true},
		{Name: "User_ID", Type: field.TypeInt, Nullable: true},
	}
	// BookBorrowsTable holds the schema information for the "book_borrows" table.
	BookBorrowsTable = &schema.Table{
		Name:       "book_borrows",
		Columns:    BookBorrowsColumns,
		PrimaryKey: []*schema.Column{BookBorrowsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "book_borrows_books_Booklist",
				Columns: []*schema.Column{BookBorrowsColumns[2]},

				RefColumns: []*schema.Column{BooksColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "book_borrows_purposes_for",
				Columns: []*schema.Column{BookBorrowsColumns[3]},

				RefColumns: []*schema.Column{PurposesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "book_borrows_users_Borrow",
				Columns: []*schema.Column{BookBorrowsColumns[4]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PurposesColumns holds the columns for the "purposes" table.
	PurposesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "purpose_name", Type: field.TypeString, Unique: true},
	}
	// PurposesTable holds the schema information for the "purposes" table.
	PurposesTable = &schema.Table{
		Name:        "purposes",
		Columns:     PurposesColumns,
		PrimaryKey:  []*schema.Column{PurposesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// RolesColumns holds the columns for the "roles" table.
	RolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "role_name", Type: field.TypeString, Unique: true},
	}
	// RolesTable holds the schema information for the "roles" table.
	RolesTable = &schema.Table{
		Name:        "roles",
		Columns:     RolesColumns,
		PrimaryKey:  []*schema.Column{RolesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// StatusColumns holds the columns for the "status" table.
	StatusColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "status_name", Type: field.TypeString},
	}
	// StatusTable holds the schema information for the "status" table.
	StatusTable = &schema.Table{
		Name:        "status",
		Columns:     StatusColumns,
		PrimaryKey:  []*schema.Column{StatusColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_email", Type: field.TypeString, Unique: true},
		{Name: "user_name", Type: field.TypeString},
		{Name: "ROLE_ID", Type: field.TypeInt, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "users_roles_role",
				Columns: []*schema.Column{UsersColumns[3]},

				RefColumns: []*schema.Column{RolesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BooksTable,
		BookBorrowsTable,
		PurposesTable,
		RolesTable,
		StatusTable,
		UsersTable,
	}
)

func init() {
	BooksTable.ForeignKeys[0].RefTable = StatusTable
	BookBorrowsTable.ForeignKeys[0].RefTable = BooksTable
	BookBorrowsTable.ForeignKeys[1].RefTable = PurposesTable
	BookBorrowsTable.ForeignKeys[2].RefTable = UsersTable
	UsersTable.ForeignKeys[0].RefTable = RolesTable
}
