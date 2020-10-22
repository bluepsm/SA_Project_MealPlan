package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Meal holds the schema definition for the Meal entity.
type Meal struct {
	ent.Schema
}

// Fields of the Meal.
func (Meal) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Unique(),
	}
}

// Edges of the Meal.
func (Meal) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("mealplans", Mealplan.Type).StorageKey(edge.Column("meal_id")),
	}
}
