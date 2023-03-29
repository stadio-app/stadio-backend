package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Address holds the schema definition for the Address entity.
type Address struct {
	ent.Schema
}

func (Address) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the Address.
func (Address) Fields() []ent.Field {
	return []ent.Field{
		field.Float("latitude"),
		field.Float("longitude"),
		field.Text("maps_link"),
		field.String("full_address"),
	}
}

// Edges of the Address.
func (Address) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("location", Location.Type).
			Ref("address").
			Unique().
			Required(),
	}
}
