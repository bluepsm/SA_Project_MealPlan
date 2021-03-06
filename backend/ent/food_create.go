// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/bluepsm/app/ent/food"
	"github.com/bluepsm/app/ent/mealplan"
	"github.com/bluepsm/app/ent/user"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// FoodCreate is the builder for creating a Food entity.
type FoodCreate struct {
	config
	mutation *FoodMutation
	hooks    []Hook
}

// SetName sets the name field.
func (fc *FoodCreate) SetName(s string) *FoodCreate {
	fc.mutation.SetName(s)
	return fc
}

// SetIng sets the ing field.
func (fc *FoodCreate) SetIng(s string) *FoodCreate {
	fc.mutation.SetIng(s)
	return fc
}

// SetCalories sets the calories field.
func (fc *FoodCreate) SetCalories(i int) *FoodCreate {
	fc.mutation.SetCalories(i)
	return fc
}

// SetAdded sets the added field.
func (fc *FoodCreate) SetAdded(s string) *FoodCreate {
	fc.mutation.SetAdded(s)
	return fc
}

// AddMealplanIDs adds the mealplans edge to Mealplan by ids.
func (fc *FoodCreate) AddMealplanIDs(ids ...int) *FoodCreate {
	fc.mutation.AddMealplanIDs(ids...)
	return fc
}

// AddMealplans adds the mealplans edges to Mealplan.
func (fc *FoodCreate) AddMealplans(m ...*Mealplan) *FoodCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return fc.AddMealplanIDs(ids...)
}

// SetOwnerID sets the owner edge to User by id.
func (fc *FoodCreate) SetOwnerID(id int) *FoodCreate {
	fc.mutation.SetOwnerID(id)
	return fc
}

// SetNillableOwnerID sets the owner edge to User by id if the given value is not nil.
func (fc *FoodCreate) SetNillableOwnerID(id *int) *FoodCreate {
	if id != nil {
		fc = fc.SetOwnerID(*id)
	}
	return fc
}

// SetOwner sets the owner edge to User.
func (fc *FoodCreate) SetOwner(u *User) *FoodCreate {
	return fc.SetOwnerID(u.ID)
}

// Mutation returns the FoodMutation object of the builder.
func (fc *FoodCreate) Mutation() *FoodMutation {
	return fc.mutation
}

// Save creates the Food in the database.
func (fc *FoodCreate) Save(ctx context.Context) (*Food, error) {
	if _, ok := fc.mutation.Name(); !ok {
		return nil, &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if _, ok := fc.mutation.Ing(); !ok {
		return nil, &ValidationError{Name: "ing", err: errors.New("ent: missing required field \"ing\"")}
	}
	if _, ok := fc.mutation.Calories(); !ok {
		return nil, &ValidationError{Name: "calories", err: errors.New("ent: missing required field \"calories\"")}
	}
	if _, ok := fc.mutation.Added(); !ok {
		return nil, &ValidationError{Name: "added", err: errors.New("ent: missing required field \"added\"")}
	}
	var (
		err  error
		node *Food
	)
	if len(fc.hooks) == 0 {
		node, err = fc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FoodMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fc.mutation = mutation
			node, err = fc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fc.hooks) - 1; i >= 0; i-- {
			mut = fc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FoodCreate) SaveX(ctx context.Context) *Food {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (fc *FoodCreate) sqlSave(ctx context.Context) (*Food, error) {
	f, _spec := fc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	f.ID = int(id)
	return f, nil
}

func (fc *FoodCreate) createSpec() (*Food, *sqlgraph.CreateSpec) {
	var (
		f     = &Food{config: fc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: food.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: food.FieldID,
			},
		}
	)
	if value, ok := fc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: food.FieldName,
		})
		f.Name = value
	}
	if value, ok := fc.mutation.Ing(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: food.FieldIng,
		})
		f.Ing = value
	}
	if value, ok := fc.mutation.Calories(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: food.FieldCalories,
		})
		f.Calories = value
	}
	if value, ok := fc.mutation.Added(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: food.FieldAdded,
		})
		f.Added = value
	}
	if nodes := fc.mutation.MealplansIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   food.MealplansTable,
			Columns: []string{food.MealplansColumn},
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
	if nodes := fc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   food.OwnerTable,
			Columns: []string{food.OwnerColumn},
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
	return f, _spec
}
