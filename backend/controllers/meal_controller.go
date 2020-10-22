package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/bluepsm/app/ent"
	"github.com/bluepsm/app/ent/meal"
	"github.com/gin-gonic/gin"
)

// MealController defines the struct for the meal controller
type MealController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateMeal handles POST requests for adding meal entities
// @Summary Create meal
// @Description Create meal
// @ID create-meal
// @Accept   json
// @Produce  json
// @Param meal body ent.Meal true "Meal entity"
// @Success 200 {object} ent.Meal
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /meals [post]
func (ctl *MealController) CreateMeal(c *gin.Context) {
	obj := ent.Meal{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "meal binding failed",
		})
		return
	}

	m, err := ctl.client.Meal.
		Create().
		SetName(obj.Name).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, m)
}

// GetMeal handles GET requests to retrieve a meal entity
// @Summary Get a meal entity by ID
// @Description get meal by ID
// @ID get-meal
// @Produce  json
// @Param id path int true "Meal ID"
// @Success 200 {object} ent.Meal
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /meals/{id} [get]
func (ctl *MealController) GetMeal(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	m, err := ctl.client.Meal.
		Query().
		Where(meal.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, m)
}

// ListMeal handles request to get a list of meal entities
// @Summary List meal entities
// @Description list meal entities
// @ID list-meal
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Meal
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /meals [get]
func (ctl *MealController) ListMeal(c *gin.Context) {
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

	meals, err := ctl.client.Meal.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, meals)
}

// DeleteMeal handles DELETE requests to delete a meal entity
// @Summary Delete a meal entity by ID
// @Description get meal by ID
// @ID delete-meal
// @Produce  json
// @Param id path int true "Meal ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /meals/{id} [delete]
func (ctl *MealController) DeleteMeal(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Meal.
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

// UpdateMeal handles PUT requests to update a meal entity
// @Summary Update a meal entity by ID
// @Description update meal by ID
// @ID update-meal
// @Accept   json
// @Produce  json
// @Param id path int true "Meal ID"
// @Param meal body ent.Meal true "Meal entity"
// @Success 200 {object} ent.Meal
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /foodmenus/{id} [put]
func (ctl *MealController) UpdateMeal(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Meal{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "meal binding failed",
		})
		return
	}

	f, err := ctl.client.Meal.
		UpdateOneID(int(id)).
		SetName(obj.Name).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "update meal failed",
		})
		return
	}

	c.JSON(200, f)
}

// NewMealController creates and registers handles for the meal controller
func NewMealController(router gin.IRouter, client *ent.Client) *MealController {
	mc := &MealController{
		client: client,
		router: router,
	}
	mc.register()
	return mc
}

// InitMealController registers routes to the main engine
func (ctl *MealController) register() {
	meals := ctl.router.Group("/meals")

	meals.GET("", ctl.ListMeal)

	// CRUD
	meals.POST("", ctl.CreateMeal)
	meals.GET(":id", ctl.GetMeal)
	meals.PUT(":id", ctl.UpdateMeal)
	meals.DELETE(":id", ctl.DeleteMeal)
}
