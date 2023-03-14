package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Owner holds the schema definition for the Owner entity.
type Owner struct {
	ent.Schema
}

// Fields of the Owner.
func (Owner) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Unique().
			Default(uuid.New),
		field.String("first_name").
			NotEmpty(),
		field.String("middle_name").
			Optional(),
		field.String("last_name").
			NotEmpty(),
		field.String("full_name").
			NotEmpty(),
		field.String("id_url").
			NotEmpty().
			Unique(),
		field.Bool("verified").
			Default(false),
	}
}

// Edges of the Owner.
func (Owner) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("owner").
			Unique().
			Required(),
		edge.To("locations", Location.Type),
	}
}
