package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Location holds the schema definition for the Location entity.
type Location struct {
	ent.Schema
}

// Fields of the Location.
func (Location) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("type"), // Making it a string temporarily until we define an Enum or a different way of defining a type.
	}
}

// Edges of the Location.
func (Location) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("address", Address.Type).Unique(),
		edge.To("reviews", Review.Type),
		edge.From("owner", Owner.Type).
			Ref("locations").
			Unique(),
		edge.To("events", Event.Type),
	}
}
