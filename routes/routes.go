package routes

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"api-recipe-keeper/config"
	"api-recipe-keeper/controllers"
	"api-recipe-keeper/middlewares"
)

func Routes(db *gorm.DB) *echo.Echo {

	e := echo.New()
	env := config.Get("APP_ENV").String()
	if env == "production" || env == "development" {
		e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
			StackSize: 1 << 10, // 1 KB
		}))
	}
	e.HTTPErrorHandler = middlewares.ErrorHandler
	e.Use(middlewares.TransactionHandler(db))

	e.GET("/api/setting", controllers.SettingGetApiHandle)
	e.POST("/api/setting", controllers.SettingUpdateApiHandle)

	e.GET("/api/ingridients/:id", controllers.IngridientGetByIdApiHandle)
	e.GET("/api/ingridients", controllers.IngridientGetPaginatedApiHandle)
	e.POST("/api/ingridients", controllers.IngridientCreateApiHandle)
	e.PUT("/api/ingridients/:id", controllers.IngridientUpdateByIdApiHandle)
	e.DELETE("/api/ingridients/:id", controllers.IngridientDeleteByIdApiHandle)

	e.GET("/api/food_categories/:id", controllers.FoodCategoryGetByIdApiHandle)
	e.GET("/api/food_categories", controllers.FoodCategoryGetPaginatedApiHandle)
	e.POST("/api/food_categories", controllers.FoodCategoryCreateApiHandle)
	e.PUT("/api/food_categories/:id", controllers.FoodCategoryUpdateByIdApiHandle)
	e.DELETE("/api/food_categories/:id", controllers.FoodCategoryDeleteByIdApiHandle)

	e.GET("/api/recipes/:id", controllers.RecipeGetByIdApiHandle)
	e.GET("/api/recipes", controllers.RecipeGetPaginatedApiHandle)
	e.POST("/api/recipes", controllers.RecipeCreateApiHandle)
	e.PUT("/api/recipes/:id", controllers.RecipeUpdateByIdApiHandle)
	e.DELETE("/api/recipes/:id", controllers.RecipeDeleteByIdApiHandle)

	return e
}
