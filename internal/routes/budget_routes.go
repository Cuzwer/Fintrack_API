package routes

import (
	"github.com/cuzwer/fintrack/internal/handler"
	"github.com/cuzwer/fintrack/internal/middleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)


func Settup_Budgets(router fiber.Router , db *gorm.DB ) {
		bugetGroup := router.Group("/budget")
		handler_budg  :=  handler.BudgetOBJ_handler(db)

		bugetGroup.Post("/", middleware.VerifySession() , handler_budg.PostBudget_Handler() )
		bugetGroup.Get("/", middleware.VerifySession(), handler_budg.GetBudget_Handler())
		bugetGroup.Put("/", middleware.VerifySession(),  handler_budg.UpdateBudget_Handler())
		bugetGroup.Delete("/:id", middleware.VerifySession() ,  handler_budg.DeleteBudget())
}
