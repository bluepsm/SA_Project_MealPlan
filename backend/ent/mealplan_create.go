// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/bluepsm/app/ent/food"
	"github.com/bluepsm/app/ent/meal"
	"github.com/bluepsm/app/ent/mealplan"
	"github.com/bluepsm/app/ent/user"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// MealplanCreate is the builder for creating a Mealplan entity.
type MealplanCreate struct {
	config
	mutation *MealplanMutation
	hooks    []Hook
}

// SetDate sets the date field.
func (mc *MealplanCreate) SetDate(t time.Time) *MealplanCreate {
	mc.mutation.SetDate(t)
	return mc
}

// SetOwnerID sets the owner edge to User by id.
func (mc *MealplanCreate) SetOwnerID(id int) *MealplanCreate {
	mc.mutation.SetOwnerID(id)
	return mc
}

// SetNillableOwnerID sets the owner edge to User by id if the given value is not nil.
func (mc *MealplanCreate) SetNillableOwnerID(id *int) *MealplanCreate {
	if id != nil {
		mc = mc.SetOwnerID(*id)
	}
	return mc
}

// SetOwner sets the owner edge to User.
func (mc *MealplanCreate) SetOwner(u *User) *MealplanCreate {
	return mc.SetOwnerID(u.ID)
}

// SetFoodsID sets the foods edge to Food by id.
func (mc *MealplanCreate) SetFoodsID(id int) *MealplanCreate {
	mc.mutation.SetFoodsID(id)
	return mc
}

// SetNillableFoodsID sets the foods edge to Food by id if the given value is not nil.
func (mc *MealplanCreate) SetNillableFoodsID(id *int) *MealplanCreate {
	if id != nil {
		mc = mc.SetFoodsID(*id)
	}
	return mc
}

// SetFoods sets the foods edge to Food.
func (mc *MealplanCreate) SetFoods(f *Food) *MealplanCreate {
	return mc.SetFoodsID(f.ID)
}

// SetMealsID sets the meals edge to Meal by id.
func (mc *MealplanCreate) SetMealsID(id int) *MealplanCreate {
	mc.mutation.SetMealsID(id)
	return mc
}

// SetNillableMealsID sets the meals edge to Meal by id if the given value is not nil.
func (mc *MealplanCreate) SetNillableMealsID(id *int) *MealplanCreate {
	if id != nil {
		mc = mc.SetMealsID(*id)
	}
	return mc
}

// SetMeals sets the meals edge to Meal.
func (mc *MealplanCreate) SetMeals(m *Meal) *MealplanCreate {
	return mc.SetMealsID(m.ID)
}

// Mutation returns the MealplanMutation object of the builder.
func (mc *MealplanCreate) Mutation() *MealplanMutation {
	return mc.mutation
}

// Save creates the Mealplan in the database.
func (mc *MealplanCreate) Save(ctx context.Context) (*Mealplan, error) {
	if _, ok := mc.mutation.Date(); !ok {
		return nil, &ValidationError{Name: "date", err: errors.New("ent: missing required field \"date\"")}
	}
	var (
		err  error
		node *Mealplan
	)
	if len(mc.hooks) == 0 {
		node, err = mc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MealplanMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mc.mutation = mutation
			node, err = mc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(mc.hooks) - 1; i >= 0; i-- {
			mut = mc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MealplanCreate) SaveX(ctx context.Context) *Mealplan {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mc *MealplanCreate) sqlSave(ctx context.Context) (*Mealplan, error) {
	m, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	m.ID = int(id)
	return m, nil
}

func (mc *MealplanCreate) createSpec() (*Mealplan, *sqlgraph.CreateSpec) {
	var (
		m     = &Mealplan{config: mc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: mealplan.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: mealplan.FieldID,
			},
		}
	)
	if value, ok := mc.mutation.Date(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: mealplan.FieldDate,
		})
		m.Date = value
	}
	if nodes := mc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mealplan.OwnerTable,
			Columns: []string{mealplan.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.FoodsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mealplan.FoodsTable,
			Columns: []string{mealplan.FoodsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: food.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.MealsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mealplan.MealsTable,
			Columns: []string{mealplan.MealsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: meal.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return m, _spec
}
