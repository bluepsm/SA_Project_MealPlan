// Code generated by entc, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"

	// EdgeMealplans holds the string denoting the mealplans edge name in mutations.
	EdgeMealplans = "mealplans"
	// EdgeFoods holds the string denoting the foods edge name in mutations.
	EdgeFoods = "foods"

	// Table holds the table name of the user in the database.
	Table = "users"
	// MealplansTable is the table the holds the mealplans relation/edge.
	MealplansTable = "mealplans"
	// MealplansInverseTable is the table name for the Mealplan entity.
	// It exists in this package in order to avoid circular dependency with the "mealplan" package.
	MealplansInverseTable = "mealplans"
	// MealplansColumn is the table column denoting the mealplans relation/edge.
	MealplansColumn = "user_id"
	// FoodsTable is the table the holds the foods relation/edge.
	FoodsTable = "foods"
	// FoodsInverseTable is the table name for the Food entity.
	// It exists in this package in order to avoid circular dependency with the "food" package.
	FoodsInverseTable = "foods"
	// FoodsColumn is the table column denoting the foods relation/edge.
	FoodsColumn = "user_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldEmail,
}
