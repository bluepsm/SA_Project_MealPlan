// Code generated by entc, DO NOT EDIT.

package mealplan

import (
	"time"

	"github.com/bluepsm/app/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Mealplan {
	return predicate.Mealplan(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Mealplan {
	return predicate.Mealplan(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Mealplan {
	return predicate.Mealplan(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Mealplan {
	return predicate.Mealplan(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Mealplan {
	return predicate.Mealplan(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Mealplan {
	return predicate.Mealplan(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Mealplan {
	return predicate.Mealplan(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Mealplan {
	return predicate.Mealplan(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Mealplan {
	return predicate.Mealplan(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Date applies equality check predicate on the "date" field. It's identical to DateEQ.
func Date(v time.Time) predicate.Mealplan {
	return predicate.Mealplan(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDate), v))
	})
}

// DateEQ applies the EQ predicate on the "date" field.
func DateEQ(v time.Time) predicate.Mealplan {
	return predicate.Mealplan(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDate), v))
	})
}

// DateNEQ applies the NEQ predicate on the "date" field.
func DateNEQ(v time.Time) predicate.Mealplan {
	return predicate.Mealplan(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDate), v))
	})
}

// DateIn applies the In predicate on the "date" field.
func DateIn(vs ...time.Time) predicate.Mealplan {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Mealplan(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDate), v...))
	})
}

// DateNotIn applies the NotIn predicate on the "date" field.
func DateNotIn(vs ...time.Time) predicate.Mealplan {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Mealplan(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDate), v...))
	})
}

// DateGT applies the GT predicate on the "date" field.
func DateGT(v time.Time) predicate.Mealplan {
	return predicate.Mealplan(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDate), v))
	})
}

// DateGTE applies the GTE predicate on the "date" field.
func DateGTE(v time.Time) predicate.Mealplan {
	return predicate.Mealplan(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDate), v))
	})
}

// DateLT applies the LT predicate on the "date" field.
func DateLT(v time.Time) predicate.Mealplan {
	return predicate.Mealplan(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDate), v))
	})
}

// DateLTE applies the LTE predicate on the "date" field.
func DateLTE(v time.Time) predicate.Mealplan {
	return predicate.Mealplan(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDate), v))
	})
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Mealplan {
	return predicate.Mealplan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.Mealplan {
	return predicate.Mealplan(func(s *sql.Selector) {
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

// HasFoods applies the HasEdge predicate on the "foods" edge.
func HasFoods() predicate.Mealplan {
	return predicate.Mealplan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FoodsTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, FoodsTable, FoodsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFoodsWith applies the HasEdge predicate on the "foods" edge with a given conditions (other predicates).
func HasFoodsWith(preds ...predicate.Food) predicate.Mealplan {
	return predicate.Mealplan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FoodsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, FoodsTable, FoodsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMeals applies the HasEdge predicate on the "meals" edge.
func HasMeals() predicate.Mealplan {
	return predicate.Mealplan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MealsTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, MealsTable, MealsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMealsWith applies the HasEdge predicate on the "meals" edge with a given conditions (other predicates).
func HasMealsWith(preds ...predicate.Meal) predicate.Mealplan {
	return predicate.Mealplan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MealsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, MealsTable, MealsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Mealplan) predicate.Mealplan {
	return predicate.Mealplan(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Mealplan) predicate.Mealplan {
	return predicate.Mealplan(func(s *sql.Selector) {
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
func Not(p predicate.Mealplan) predicate.Mealplan {
	return predicate.Mealplan(func(s *sql.Selector) {
		p(s.Not())
	})
}
