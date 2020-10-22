package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/bluepsm/app/ent"
	"github.com/bluepsm/app/ent/food"
	"github.com/bluepsm/app/ent/meal"
	"github.com/bluepsm/app/ent/mealplan"
	"github.com/bluepsm/app/ent/user"
	"github.com/gin-gonic/gin"
)

// MealplanController defines the struct for the mealplan controller
type MealplanController struct {
	client *ent.Client
	router gin.IRouter
}

// Mealplan defines the struct for the mealplan
type Mealplan struct {
	Owner int
	Food  int
	Meal  int
	Date  string
}

// CreateMealplan handles POST requests for adding mealplan entities
// @Summary Create mealplan
// @Description Create mealplan
// @ID create-mealplan
// @Accept   json
// @Produce  json
// @Param mealplan body Mealplan true "Mealplan entity"
// @Success 200 {object} ent.Mealplan
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /mealplans [post]
func (ctl *MealplanController) CreateMealplan(c *gin.Context) {
	obj := Mealplan{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "meal plan binding failed",
		})
		return
	}

	u, err := ctl.client.User.
		Query().
		Where(user.IDEQ(int(obj.Owner))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "user not found",
		})
		return
	}

	m, err := ctl.client.Meal.
		Query().
		Where(meal.IDEQ(int(obj.Meal))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "meal not found",
		})
		return
	}

	f, err := ctl.client.Food.
		Query().
		Where(food.IDEQ(int(obj.Food))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "food not found",
		})
		return
	}

	time, err := time.Parse(time.RFC3339, obj.Date)

	mp, err := ctl.client.Mealplan.
		Create().
		SetOwner(u).
		SetFoods(f).
		SetMeals(m).
		SetDate(time).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, mp)
}

// GetMealplan handles GET requests to retrieve a mealplan entity
// @Summary Get a mealplan entity by ID
// @Description get mealplan by ID
// @ID get-mealplan
// @Produce  json
// @Param id path int true "Mealplan ID"
// @Success 200 {object} ent.Mealplan
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /mealplans/{id} [get]
func (ctl *MealplanController) GetMealplan(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	mp, err := ctl.client.Mealplan.
		Query().
		Where(mealplan.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, mp)
}

// ListMealplan handles request to get a list of mealplan entities
// @Summary List mealplan entities
// @Description list mealplan entities
// @ID list-mealplan
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Mealplan
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /mealplans [get]
func (ctl *MealplanController) ListMealplan(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	mealplans, err := ctl.client.Mealplan.
		Query().
		WithOwner().
		WithFoods().
		WithMeals().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, mealplans)
}

// DeleteMealplan handles DELETE requests to delete a mealplan entity
// @Summary Delete a mealplan entity by ID
// @Description get mealplan by ID
// @ID delete-mealplan
// @Produce  json
// @Param id path int true "Mealplan ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /mealplans/{id} [delete]
func (ctl *MealplanController) DeleteMealplan(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Mealplan.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// UpdateMealplan handles PUT requests to update a mealplan entity
// @Summary Update a mealplan entity by ID
// @Description update mealplan by ID
// @ID update-mealplan
// @Accept   json
// @Produce  json
// @Param id path int true "Mealplan ID"
// @Param mealplan body ent.Mealplan true "Mealplan entity"
// @Success 200 {object} ent.Mealplan
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /mealplans/{id} [put]
func (ctl *MealplanController) UpdateMealplan(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Mealplan{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "meal plan binding failed",
		})
		return
	}
	obj.ID = int(id)
	mp, err := ctl.client.Mealplan.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, mp)
}

// NewMealplanController creates and registers handles for the mealplan controller
func NewMealplanController(router gin.IRouter, client *ent.Client) *MealplanController {
	mpc := &MealplanController{
		client: client,
		router: router,
	}
	mpc.register()
	return mpc
}

// InitMealplanController registers routes to the main engine
func (ctl *MealplanController) register() {
	mealplans := ctl.router.Group("/mealplans")

	mealplans.GET("", ctl.ListMealplan)

	// CRUD
	mealplans.POST("", ctl.CreateMealplan)
	mealplans.GET(":id", ctl.GetMealplan)
	mealplans.PUT(":id", ctl.UpdateMealplan)
	mealplans.DELETE(":id", ctl.DeleteMealplan)
}
