package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("USER_EMAIL").NotEmpty().Unique(),
		field.String("USER_NAME").NotEmpty(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("Borrow", BookBorrow.Type).StorageKey(edge.Column("User_ID")),
		edge.From("RolePlay", Role.Type).Ref("role").Unique(),
	}
}
