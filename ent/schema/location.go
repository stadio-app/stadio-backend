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
		field.Int64("location_id").
			Unique(),
		field.String("name"),
		field.String("type"), // Making it a string temporarily until we define an Enum or a different way of defining a type.
	}
}

// Edges of the Location.
func (Location) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("address", Address.Type),
		edge.From("reviews", Review.Type).
			Ref("location"),
		edge.From("owner", Owner.Type).
			Ref("locations").
			Unique(),
	}
}
