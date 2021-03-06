// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/bluepsm/app/ent/food"
	"github.com/bluepsm/app/ent/user"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Food is the model entity for the Food schema.
type Food struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Ing holds the value of the "ing" field.
	Ing string `json:"ing,omitempty"`
	// Calories holds the value of the "calories" field.
	Calories int `json:"calories,omitempty"`
	// Added holds the value of the "added" field.
	Added string `json:"added,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FoodQuery when eager-loading is set.
	Edges   FoodEdges `json:"edges"`
	user_id *int
}

// FoodEdges holds the relations/edges for other nodes in the graph.
type FoodEdges struct {
	// Mealplans holds the value of the mealplans edge.
	Mealplans []*Mealplan
	// Owner holds the value of the owner edge.
	Owner *User
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// MealplansOrErr returns the Mealplans value or an error if the edge
// was not loaded in eager-loading.
func (e FoodEdges) MealplansOrErr() ([]*Mealplan, error) {
	if e.loadedTypes[0] {
		return e.Mealplans, nil
	}
	return nil, &NotLoadedError{edge: "mealplans"}
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FoodEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.Owner == nil {
			// The edge owner was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Food) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // name
		&sql.NullString{}, // ing
		&sql.NullInt64{},  // calories
		&sql.NullString{}, // added
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Food) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // user_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Food fields.
func (f *Food) assignValues(values ...interface{}) error {
	if m, n := len(values), len(food.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	f.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		f.Name = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field ing", values[1])
	} else if value.Valid {
		f.Ing = value.String
	}
	if value, ok := values[2].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field calories", values[2])
	} else if value.Valid {
		f.Calories = int(value.Int64)
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field added", values[3])
	} else if value.Valid {
		f.Added = value.String
	}
	values = values[4:]
	if len(values) == len(food.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field user_id", value)
		} else if value.Valid {
			f.user_id = new(int)
			*f.user_id = int(value.Int64)
		}
	}
	return nil
}

// QueryMealplans queries the mealplans edge of the Food.
func (f *Food) QueryMealplans() *MealplanQuery {
	return (&FoodClient{config: f.config}).QueryMealplans(f)
}

// QueryOwner queries the owner edge of the Food.
func (f *Food) QueryOwner() *UserQuery {
	return (&FoodClient{config: f.config}).QueryOwner(f)
}

// Update returns a builder for updating this Food.
// Note that, you need to call Food.Unwrap() before calling this method, if this Food
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *Food) Update() *FoodUpdateOne {
	return (&FoodClient{config: f.config}).UpdateOne(f)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (f *Food) Unwrap() *Food {
	tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: Food is not a transactional entity")
	}
	f.config.driver = tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *Food) String() string {
	var builder strings.Builder
	builder.WriteString("Food(")
	builder.WriteString(fmt.Sprintf("id=%v", f.ID))
	builder.WriteString(", name=")
	builder.WriteString(f.Name)
	builder.WriteString(", ing=")
	builder.WriteString(f.Ing)
	builder.WriteString(", calories=")
	builder.WriteString(fmt.Sprintf("%v", f.Calories))
	builder.WriteString(", added=")
	builder.WriteString(f.Added)
	builder.WriteByte(')')
	return builder.String()
}

// Foods is a parsable slice of Food.
type Foods []*Food

func (f Foods) config(cfg config) {
	for _i := range f {
		f[_i].config = cfg
	}
}
