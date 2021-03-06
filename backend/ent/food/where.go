// Code generated by entc, DO NOT EDIT.

package food

import (
	"github.com/bluepsm/app/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Ing applies equality check predicate on the "ing" field. It's identical to IngEQ.
func Ing(v string) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIng), v))
	})
}

// Calories applies equality check predicate on the "calories" field. It's identical to CaloriesEQ.
func Calories(v int) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCalories), v))
	})
}

// Added applies equality check predicate on the "added" field. It's identical to AddedEQ.
func Added(v string) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAdded), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Food {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Food(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Food {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Food(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// IngEQ applies the EQ predicate on the "ing" field.
func IngEQ(v string) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIng), v))
	})
}

// IngNEQ applies the NEQ predicate on the "ing" field.
func IngNEQ(v string) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIng), v))
	})
}

// IngIn applies the In predicate on the "ing" field.
func IngIn(vs ...string) predicate.Food {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Food(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldIng), v...))
	})
}

// IngNotIn applies the NotIn predicate on the "ing" field.
func IngNotIn(vs ...string) predicate.Food {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Food(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldIng), v...))
	})
}

// IngGT applies the GT predicate on the "ing" field.
func IngGT(v string) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldIng), v))
	})
}

// IngGTE applies the GTE predicate on the "ing" field.
func IngGTE(v string) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldIng), v))
	})
}

// IngLT applies the LT predicate on the "ing" field.
func IngLT(v string) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldIng), v))
	})
}

// IngLTE applies the LTE predicate on the "ing" field.
func IngLTE(v string) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldIng), v))
	})
}

// IngContains applies the Contains predicate on the "ing" field.
func IngContains(v string) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldIng), v))
	})
}

// IngHasPrefix applies the HasPrefix predicate on the "ing" field.
func IngHasPrefix(v string) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldIng), v))
	})
}

// IngHasSuffix applies the HasSuffix predicate on the "ing" field.
func IngHasSuffix(v string) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldIng), v))
	})
}

// IngEqualFold applies the EqualFold predicate on the "ing" field.
func IngEqualFold(v string) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldIng), v))
	})
}

// IngContainsFold applies the ContainsFold predicate on the "ing" field.
func IngContainsFold(v string) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldIng), v))
	})
}

// CaloriesEQ applies the EQ predicate on the "calories" field.
func CaloriesEQ(v int) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCalories), v))
	})
}

// CaloriesNEQ applies the NEQ predicate on the "calories" field.
func CaloriesNEQ(v int) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCalories), v))
	})
}

// CaloriesIn applies the In predicate on the "calories" field.
func CaloriesIn(vs ...int) predicate.Food {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Food(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCalories), v...))
	})
}

// CaloriesNotIn applies the NotIn predicate on the "calories" field.
func CaloriesNotIn(vs ...int) predicate.Food {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Food(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCalories), v...))
	})
}

// CaloriesGT applies the GT predicate on the "calories" field.
func CaloriesGT(v int) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCalories), v))
	})
}

// CaloriesGTE applies the GTE predicate on the "calories" field.
func CaloriesGTE(v int) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCalories), v))
	})
}

// CaloriesLT applies the LT predicate on the "calories" field.
func CaloriesLT(v int) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCalories), v))
	})
}

// CaloriesLTE applies the LTE predicate on the "calories" field.
func CaloriesLTE(v int) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCalories), v))
	})
}

// AddedEQ applies the EQ predicate on the "added" field.
func AddedEQ(v string) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAdded), v))
	})
}

// AddedNEQ applies the NEQ predicate on the "added" field.
func AddedNEQ(v string) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAdded), v))
	})
}

// AddedIn applies the In predicate on the "added" field.
func AddedIn(vs ...string) predicate.Food {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Food(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAdded), v...))
	})
}

// AddedNotIn applies the NotIn predicate on the "added" field.
func AddedNotIn(vs ...string) predicate.Food {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Food(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAdded), v...))
	})
}

// AddedGT applies the GT predicate on the "added" field.
func AddedGT(v string) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAdded), v))
	})
}

// AddedGTE applies the GTE predicate on the "added" field.
func AddedGTE(v string) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAdded), v))
	})
}

// AddedLT applies the LT predicate on the "added" field.
func AddedLT(v string) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAdded), v))
	})
}

// AddedLTE applies the LTE predicate on the "added" field.
func AddedLTE(v string) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAdded), v))
	})
}

// AddedContains applies the Contains predicate on the "added" field.
func AddedContains(v string) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAdded), v))
	})
}

// AddedHasPrefix applies the HasPrefix predicate on the "added" field.
func AddedHasPrefix(v string) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAdded), v))
	})
}

// AddedHasSuffix applies the HasSuffix predicate on the "added" field.
func AddedHasSuffix(v string) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAdded), v))
	})
}

// AddedEqualFold applies the EqualFold predicate on the "added" field.
func AddedEqualFold(v string) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAdded), v))
	})
}

// AddedContainsFold applies the ContainsFold predicate on the "added" field.
func AddedContainsFold(v string) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAdded), v))
	})
}

// HasMealplans applies the HasEdge predicate on the "mealplans" edge.
func HasMealplans() predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MealplansTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MealplansTable, MealplansColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMealplansWith applies the HasEdge predicate on the "mealplans" edge with a given conditions (other predicates).
func HasMealplansWith(preds ...predicate.Mealplan) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MealplansInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MealplansTable, MealplansColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Food) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Food) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Food) predicate.Food {
	return predicate.Food(func(s *sql.Selector) {
		p(s.Not())
	})
}
