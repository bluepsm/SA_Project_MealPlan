package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/bluepsm/app/ent"
	"github.com/bluepsm/app/ent/food"
	"github.com/bluepsm/app/ent/user"
	"github.com/gin-gonic/gin"
)

// FoodController defines the struct for the food controller
type FoodController struct {
	client *ent.Client
	router gin.IRouter
}

//Food struct
type Food struct {
	Name     string
	Ing      string
	Calories int
	Added    string
	Owner    int
}

// CreateFood handles POST requests for adding food entities
// @Summary Create food
// @Description Create food
// @ID create-food
// @Accept   json
// @Produce  json
// @Param food body ent.Food true "Food entity"
// @Success 200 {object} ent.Food
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /foods [post]
func (ctl *FoodController) CreateFood(c *gin.Context) {
	obj := Food{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "food binding failed",
		})
		return
	}

	u, err := ctl.client.User.
		Query().
		Where(user.IDEQ(int(obj.Owner))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "owner not found",
		})
		return
	}

	f, err := ctl.client.Food.
		Create().
		SetOwner(u).
		SetName(obj.Name).
		SetCalories(obj.Calories).
		SetAdded(obj.Added).
		Save(context.Background())

	fmt.Println(err)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, f)
}

// GetFood handles GET requests to retrieve a food entity
// @Summary Get a food entity by ID
// @Description get food by ID
// @ID get-food
// @Produce  json
// @Param id path int true "Food ID"
// @Success 200 {object} ent.Food
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /foods/{id} [get]
func (ctl *FoodController) GetFood(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	f, err := ctl.client.Food.
		Query().
		WithOwner().
		Where(food.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, f)
}

// ListFood handles request to get a list of food entities
// @Summary List food entities
// @Description list food entities
// @ID list-food
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Food
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /foods [get]
func (ctl *FoodController) ListFood(c *gin.Context) {
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

	foods, err := ctl.client.Food.
		Query().
		WithOwner().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, foods)
}

// UpdateFood handles PUT requests to update a food entity
// @Summary Update a food entity by ID
// @Description update food by ID
// @ID update-food
// @Accept   json
// @Produce  json
// @Param id path int true "Food ID"
// @Param food body ent.Food true "Food entity"
// @Success 200 {object} ent.Food
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /foods/{id} [put]
func (ctl *FoodController) UpdateFood(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := Food{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "food binding failed",
		})
		return
	}
	u, err := ctl.client.User.
		Query().
		Where(user.IDEQ(int(obj.Owner))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	f, err := ctl.client.Food.
		UpdateOneID(int(id)).
		SetOwner(u).
		SetName(obj.Name).
		SetCalories(obj.Calories).
		SetAdded(obj.Added).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "update foodmenu failed",
		})
		return
	}

	c.JSON(200, f)
}

// DeleteFood handles DELETE requests to delete a food entity
// @Summary Delete a food entity by ID
// @Description get food by ID
// @ID delete-food
// @Produce  json
// @Param id path int true "Food ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /foods/{id} [delete]
func (ctl *FoodController) DeleteFood(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Food.
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

// NewFoodController creates and registers handles for the food controller
func NewFoodController(router gin.IRouter, client *ent.Client) *FoodController {
	fc := &FoodController{
		client: client,
		router: router,
	}
	fc.register()
	return fc
}

// InitFoodController registers routes to the main engine
func (ctl *FoodController) register() {
	foods := ctl.router.Group("/foods")

	foods.GET("", ctl.ListFood)

	// CRUD
	foods.POST("", ctl.CreateFood)
	foods.GET(":id", ctl.GetFood)
	foods.PUT(":id", ctl.UpdateFood)
	foods.DELETE(":id", ctl.DeleteFood)
}
