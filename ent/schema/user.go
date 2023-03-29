package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").
			Unique(),
		field.String("phone_number").
			Unique().
			Optional(),
		field.String("name"),
		field.Text("avatar").
			Optional(),
		field.Time("birth_date").
			Optional(),
		field.Text("bio").
			Optional(),
		field.Bool("active").
			Default(false),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("owner", Owner.Type).Unique(),
		edge.To("participants", Participant.Type),
	}
}
