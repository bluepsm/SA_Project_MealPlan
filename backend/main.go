package main

import (
	"context"
	"fmt"
	"log"

	"github.com/bluepsm/app/controllers"
	_ "github.com/bluepsm/app/docs"
	"github.com/bluepsm/app/ent"
	"github.com/bluepsm/app/ent/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// User struct
type User struct {
	Name  string
	Email string
}

// Users struct
type Users struct {
	User []User
}

// Food struct
type Food struct {
	Name     string
	Ing      string
	Calories int
	Added    string
	Owner    int
}

// Foods struct
type Foods struct {
	Food []Food
}

// Meal struct
type Meal struct {
	Name string
}

// Meals struct
type Meals struct {
	Meal []Meal
}

// @title SUT SA Example API
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
	router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:ent.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewUserController(v1, client)
	controllers.NewFoodController(v1, client)
	controllers.NewMealController(v1, client)
	controllers.NewMealplanController(v1, client)

	users := Users{
		User: []User{
			User{"Burapa Pusamart", "burapa9826@gmail.com"},
			User{"Name Surname", "me@example.com"},
		},
	}

	for _, u := range users.User {
		client.User.
			Create().
			SetEmail(u.Email).
			SetName(u.Name).
			Save(context.Background())
	}

	foods := Foods{
		Food: []Food{
			Food{"ข้าวมันไก่", "ข้าว, ไก่, แตงกวา", 350, "2020 - 07 - 12", 01},
			Food{"ข้าวผัดกุ้ง", "ข้าว, กุ้ง, ไข่", 380, "2020 - 07 - 12", 02},
		},
	}

	for _, f := range foods.Food {

		u, err := client.User.
			Query().
			Where(user.IDEQ(int(f.Owner))).
			Only(context.Background())
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		client.Food.
			Create().
			SetName(f.Name).
			SetIng(f.Ing).
			SetCalories(f.Calories).
			SetAdded(f.Added).
			SetOwner(u).
			Save(context.Background())
	}

	meals := Meals{
		Meal: []Meal{
			Meal{"มื้อเช้า"},
			Meal{"มื้อกลางวัน"},
			Meal{"มื้อเย็น"},
		},
	}

	for _, m := range meals.Meal {
		client.Meal.
			Create().
			SetName(m.Name).
			Save(context.Background())
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
