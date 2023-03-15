package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

func BaseSchema() []ent.Field {
	return []ent.Field{
		field.Uint64("id").
			Unique(),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now),
	}
}
