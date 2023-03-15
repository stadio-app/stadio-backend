package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Event holds the schema definition for the Event entity.
type Event struct {
	ent.Schema
}

// Fields of the Event.
func (Event) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("type").
			Optional(),
		field.Time("start_date").Default(time.Now),
		field.Time("end_date").
			Default(time.Now().Add(time.Hour * 24 * 7)),
	}
}

// Edges of the Event.
func (Event) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("participants", Participant.Type),
		edge.From("location", Location.Type).
			Ref("events").
			Unique(),
	}
}
