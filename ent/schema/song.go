package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Song holds the schema definition for the Song entity.
type Song struct {
	ent.Schema
}

// Fields of the Song.
func (Song) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("title"),
		field.Float("duration").SchemaType(map[string]string{
			dialect.MySQL: "decimal(6,2)",
		}),
		field.Bool("lyricsExits").Default(false),
	}
}

// Edges of the Song.
func (Song) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("artists", Artist.Type),
	}
}
