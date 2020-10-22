// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/bluepsm/app/ent/meal"
	"github.com/bluepsm/app/ent/mealplan"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// MealCreate is the builder for creating a Meal entity.
type MealCreate struct {
	config
	mutation *MealMutation
	hooks    []Hook
}

// SetName sets the name field.
func (mc *MealCreate) SetName(s string) *MealCreate {
	mc.mutation.SetName(s)
	return mc
}

// AddMealplanIDs adds the mealplans edge to Mealplan by ids.
func (mc *MealCreate) AddMealplanIDs(ids ...int) *MealCreate {
	mc.mutation.AddMealplanIDs(ids...)
	return mc
}

// AddMealplans adds the mealplans edges to Mealplan.
func (mc *MealCreate) AddMealplans(m ...*Mealplan) *MealCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mc.AddMealplanIDs(ids...)
}

// Mutation returns the MealMutation object of the builder.
func (mc *MealCreate) Mutation() *MealMutation {
	return mc.mutation
}

// Save creates the Meal in the database.
func (mc *MealCreate) Save(ctx context.Context) (*Meal, error) {
	if _, ok := mc.mutation.Name(); !ok {
		return nil, &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	var (
		err  error
		node *Meal
	)
	if len(mc.hooks) == 0 {
		node, err = mc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MealMutation)
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
func (mc *MealCreate) SaveX(ctx context.Context) *Meal {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mc *MealCreate) sqlSave(ctx context.Context) (*Meal, error) {
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

func (mc *MealCreate) createSpec() (*Meal, *sqlgraph.CreateSpec) {
	var (
		m     = &Meal{config: mc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: meal.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: meal.FieldID,
			},
		}
	)
	if value, ok := mc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: meal.FieldName,
		})
		m.Name = value
	}
	if nodes := mc.mutation.MealplansIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   meal.MealplansTable,
			Columns: []string{meal.MealplansColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: mealplan.FieldID,
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
