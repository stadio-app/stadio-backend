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
	return []ent.Field{
		field.Int64("review_id").
			Unique(),
		field.Float("rating"),
		field.Text("message"),
	}
}

// Edges of the Review.
func (Review) Edges() []ent.Edge {
	return []ent.Edge{
		// edge.To("participant", Participant.Type), // edge to Participant
		edge.To("location", Location.Type),
	}
}
