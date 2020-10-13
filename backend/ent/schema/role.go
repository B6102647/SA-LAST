package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Role holds the schema definition for the Role entity.
type Role struct {
	ent.Schema
}

// Fields of the Purpose.
func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.Int("ROLE_ID").Positive().Unique(),
		field.String("ROLE_NAME").NotEmpty(),
	}
}

// Edges of the Purpose.
func (Role) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("Role", User.Type),
	}
}
