package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

func BaseSchema() []ent.Field {
	return []ent.Field{
		field.Int64("id").
			Unique().
			Immutable(),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now),
	}
}
