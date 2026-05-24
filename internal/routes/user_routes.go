package routes

import (
	"github.com/cuzwer/fintrack/internal/handler"
	"github.com/cuzwer/fintrack/internal/middleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)


func SettupUserRoutes(router fiber.Router ,db *gorm.DB) { 
		userHandler := handler.NewUserHandler(db);


 	  UserGroup := router.Group("/user")

		UserGroup.Get("/", handler.GetallUser_handler(db))
		UserGroup.Post("/register",userHandler.RegisterUser_handler())
		UserGroup.Post("/login", middleware.Login_Limitter() ,userHandler.LoginUser_handler())
		UserGroup.Delete("/:id", middleware.VerifySession() ,userHandler.DeleteUser_handler())
}
