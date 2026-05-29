package routes

import (
	"github.com/cuzwer/fintrack/internal/handler"
	"github.com/cuzwer/fintrack/internal/middleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)



func SettupCat_route(router fiber.Router, db *gorm.DB) {
	
	catHandler := handler.CategoryHandler(db)

	catGroup := router.Group("/categories")
	

	catGroup.Get("/", middleware.VerifySession(), catHandler.GetCategory())
	catGroup.Delete("/:id", middleware.VerifySession(), catHandler.DeleteCat_handler())
	catGroup.Post("/", middleware.VerifySession() , catHandler.PostCat_hanler())
	catGroup.Put("/:id", middleware.VerifySession() , catHandler.UpdateCat_handler())
}
