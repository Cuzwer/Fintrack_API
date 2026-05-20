package routes

import (
	"github.com/cuzwer/fintrack/internal/handler"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)


func SettupUserRoutes(router fiber.Router ,db *gorm.DB) { 
		userHandler := handler.NewUserHandler(db);


 	  UserGroup := router.Group("/user")

		UserGroup.Get("/", handler.GetallUser_handler(db))
		UserGroup.Post("/register", userHandler.RegisterUser_handler())
		UserGroup.Post("/login", userHandler.LoginUser_handler())
}
