package routes

import (
	"github.com/cuzwer/fintrack/internal/handler"
	"github.com/cuzwer/fintrack/internal/middleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)



func SettupAccount_Routes(router fiber.Router, db *gorm.DB) { 
	
	DBAccount := handler.AccountConts_handler(db)
	Account_Group :=  router.Group("/account")

	Account_Group.Get("/:id", middleware.VerifySession(), DBAccount.GetAllAccount());
	Account_Group.Post("/", middleware.VerifySession(), DBAccount.PostAccount_Handler())
}
