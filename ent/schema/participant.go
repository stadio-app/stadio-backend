package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Participant holds the schema definition for the Participant entity.
type Participant struct {
	ent.Schema
}

// Fields of the Participant.
func (Participant) Fields() []ent.Field {
	return []ent.Field{
		field.String("nickname").
			Optional(),
		field.Bool("admin").
			Default(false),
		field.Bool("participates").
			Default(true),
		field.String("skill_level").
			Optional(),
	}
}

// Edges of the Participant.
func (Participant) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("participants").
			Unique(),
		edge.To("reviews", Review.Type),
		edge.From("event", Event.Type).
			Ref("participants").
			Unique(),
	}
}
