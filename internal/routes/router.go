package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)



func SetUpRoutest(app *fiber.App , db *gorm.DB){

	api := app.Group("/api/v1")
	

	SettupUserRoutes(api,db)
}
