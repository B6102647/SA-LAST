package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Status holds the schema definition for the Status entity.
type Status struct {
	ent.Schema
}

// Fields of the Book.
func (Status) Fields() []ent.Field {
	return []ent.Field{
		field.String("STATUS_NAME"),
	}
}

// Edges of the Book.
func (Status) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("status", Book.Type).
            StorageKey(edge.Column("STATUS_ID")),
	}
}
