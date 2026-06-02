package routes

import (
	"github.com/cuzwer/fintrack/internal/handler"
	"github.com/cuzwer/fintrack/internal/middleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)


func SettupTransection_Routes(router fiber.Router , db  *gorm.DB ) { 
	  
	  TransHandler := handler.TransHanler_context(db)

		TransGroup := router.Group("/transection");

		TransGroup.Post("/", middleware.VerifySession() , TransHandler.PostTrans_Handler() )
		TransGroup.Get("/", middleware.VerifySession(), TransHandler.GetAllTrans_Handler())

}
