package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Mealplan holds the schema definition for the Mealplan entity.
type Mealplan struct {
	ent.Schema
}

// Fields of the Mealplan.
func (Mealplan) Fields() []ent.Field {
	return []ent.Field{
		field.Time("date"),
	}
}

// Edges of the Mealplan.
func (Mealplan) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("mealplans").Unique(),
		edge.From("foods", Food.Type).Ref("mealplans").Unique(),
		edge.From("meals", Meal.Type).Ref("mealplans").Unique(),
	}
}
