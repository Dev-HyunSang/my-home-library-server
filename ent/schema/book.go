package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Book holds the schema definition for the Book entity.
type Book struct {
	ent.Schema
}

// Fields of the Book.
func (Book) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()),
		field.UUID("user_uuid", uuid.New()),
		field.String("title"),
		field.String("subtitle"),
		field.String("publisher"),
		field.String("publishing_company"),
		field.String("memo"),
		field.Int("total_page").
			Default(0),
		field.Int("current_page").
			Default(0),
		field.Time("created_at").
			Default(time.Now).
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}),
		field.Time("edited_at").
			Default(nil).
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}),
	}
}

// Edges of the Book.
func (Book) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Field("user_uuid").
			Unique().
			Required(),
	}
}
