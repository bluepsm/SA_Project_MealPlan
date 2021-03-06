// Code generated by entc, DO NOT EDIT.

package mealplan

const (
	// Label holds the string label denoting the mealplan type in the database.
	Label = "mealplan"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDate holds the string denoting the date field in the database.
	FieldDate = "date"

	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgeFoods holds the string denoting the foods edge name in mutations.
	EdgeFoods = "foods"
	// EdgeMeals holds the string denoting the meals edge name in mutations.
	EdgeMeals = "meals"

	// Table holds the table name of the mealplan in the database.
	Table = "mealplans"
	// OwnerTable is the table the holds the owner relation/edge.
	OwnerTable = "mealplans"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "user_id"
	// FoodsTable is the table the holds the foods relation/edge.
	FoodsTable = "mealplans"
	// FoodsInverseTable is the table name for the Food entity.
	// It exists in this package in order to avoid circular dependency with the "food" package.
	FoodsInverseTable = "foods"
	// FoodsColumn is the table column denoting the foods relation/edge.
	FoodsColumn = "menu_id"
	// MealsTable is the table the holds the meals relation/edge.
	MealsTable = "mealplans"
	// MealsInverseTable is the table name for the Meal entity.
	// It exists in this package in order to avoid circular dependency with the "meal" package.
	MealsInverseTable = "meals"
	// MealsColumn is the table column denoting the meals relation/edge.
	MealsColumn = "meal_id"
)

// Columns holds all SQL columns for mealplan fields.
var Columns = []string{
	FieldID,
	FieldDate,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Mealplan type.
var ForeignKeys = []string{
	"menu_id",
	"meal_id",
	"user_id",
}
