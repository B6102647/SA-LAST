package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Book holds the schema definition for the Book entity.
type Book struct {
	ent.Schema
}

// Fields of the Book.
func (Book) Fields() []ent.Field {
	return []ent.Field{
		field.String("BOOK_NAME").NotEmpty().Unique(),
		field.String("USER_NAME"),
		field.String("CATEGORY"),
		field.String("Author"),
	}
}

// Edges of the Book.
func (Book) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("Booklist", BookBorrow.Type).StorageKey(edge.Column("BOOK_ID")),
		edge.From("Status", Status.Type).Ref("status").Unique(),
	}
}
