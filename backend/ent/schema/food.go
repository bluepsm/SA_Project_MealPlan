package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Food holds the schema definition for the Food entity.
type Food struct {
	ent.Schema
}

// Fields of the Food.
func (Food) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Unique(),
		field.String("ing"),
		field.Int("calories"),
		field.String("added"),
	}
}

// Edges of the Food.
func (Food) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("mealplans", Mealplan.Type).StorageKey(edge.Column("menu_id")),
		edge.From("owner", User.Type).Ref("foods").Unique(),
	}
}
