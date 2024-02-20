package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()),
		field.String("nickname"),
		field.String("email"),
		field.String("password"),
		field.Time("created_at").
			Default(func() time.Time {
				return time.Now()
			}).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
		field.Time("edited_at").
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
