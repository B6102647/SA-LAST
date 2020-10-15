package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// BookBorrow holds the schema definition for the BookBorrow entity.
type BookBorrow struct {
	ent.Schema
}

// Fields of the BookBorrow.
func (BookBorrow) Fields() []ent.Field {
	return []ent.Field{
		field.Time("ADDED_TIME"),
	}
}

// Edges of the BookBorrow.
func (BookBorrow) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("USER", User.Type).Ref("Borrow").Unique(),
		edge.From("BOOK", Book.Type).Ref("Booklist").Unique(),
		edge.From("PURPOSE", Purpose.Type).Ref("for").Unique(),
	}
}
