package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Review holds the schema definition for the Review entity.
type Review struct {
	ent.Schema
}

// Fields of the Review.
func (Review) Fields() []ent.Field {
	return append(
		[]ent.Field{
			field.Float("rating"),
			field.Text("message"),
		},
		BaseSchema()...,
	)
}

// Edges of the Review.
func (Review) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("location", Location.Type).
			Ref("reviews").
			Unique(),
		edge.From("participant", Participant.Type).
			Ref("reviews").
			Unique(),
	}
}
