package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Textures holds the schema definition for the Textures entity.
type Textures struct {
	ent.Schema
}

// Fields of the Textures.
func (Textures) Fields() []ent.Field {
	return []ent.Field{
		field.Time("updated_at").Default(0),
		field.Time("created_at").Default(0),
		field.String("uuid").Default(""),
		field.String("name").Default(""),
		field.String("hash").Default(""),
		field.String("type").Default(""), // Skin Or Cape
		field.String("model").Default(""), // Default Or Slim
		field.String("url").Default(""),
	}
}

// Edges of the Textures.
func (Textures) Edges() []ent.Edge {
	return nil
}
