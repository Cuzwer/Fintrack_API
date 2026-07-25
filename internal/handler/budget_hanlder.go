package handler

import (
	"strconv"

	"github.com/cuzwer/fintrack/internal/domain"
	"github.com/cuzwer/fintrack/internal/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Budget_obj struct {
	DB *gorm.DB
}


func BudgetOBJ_handler (db *gorm.DB) *Budget_obj{ 
	var buff Budget_obj;
	budget := &buff

	budget.DB = db
	return budget
}

func(h *Budget_obj) PostBudget_Handler() fiber.Handler {
	return func(c *fiber.Ctx) error { 
		val := c.Locals("userID")
		usrId , ok := val.(uint)

		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message" : "Invalid user ID",
			})
		}
		var newBudget_obj domain.Budgets
		newBudget := &newBudget_obj
		if err := c.BodyParser(newBudget) ; err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message" : "Invalid request body",
				"error": err.Error(),
			})
		}
		
		newBudget.ID_user = int(usrId)

		err := service.PostBudget_service(newBudget,h.DB); 
		if err != nil { 
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message" : "Failed to create budget" ,
				"error": err.Error(),
			})
		}
	
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message" : "Budget created successfully",
		})
	}
}


func(h *Budget_obj) GetBudget_Handler() fiber.Handler{
	return func(c *fiber.Ctx) error { 
		val := c.Locals("userID"); 
		usrId , ok := val.(uint)

		if !ok { 
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message" : "Invalid user ID",
			})
		}

		var allBudg_obj []domain.Budgets_detail;
		allBudg := &allBudg_obj;
		err := service.GetBudget_Service(int(usrId) , allBudg , h.DB);
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message" : "Failed to retrieve budgets",
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message" : "Budgets retrieved successfully",
			"result" : allBudg,
		})
	}
}

func(h *Budget_obj) UpdateBudget_Handler() fiber.Handler{
	return func(c *fiber.Ctx) error {
		var budg_obj domain.Budgets_Edit
		budg := &budg_obj

		if err := c.BodyParser(budg) ; err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message" : "Invalid request body",
				"error": err.Error(),
			})	
		}
  	
		val := c.Locals("userID");
		userID ,  ok := val.(uint);
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message" : "Invalid user ID",
			})
		}
  	
		budg.ID_user = int(userID)
		
		err := service.UpdateBudg_Service(budg , h.DB)
		
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message" : "Failed to update budget",
				"error": err.Error(),
			})
		}
		
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message" : "Budget updated successfully",
		})
	}
}

func(h *Budget_obj) DeleteBudget() fiber.Handler{
	return  func(c *fiber.Ctx) error { 
		val := c.Locals("userID");
  	userID , ok := val.(uint);

		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message" : "Invalid user ID",
			})
		}

		id_budget , err := strconv.Atoi(c.Params("id"));

		if err != nil { 
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message" : "Invalid budget ID format",
				"error": err.Error(),
			})
		}
		
		var detail_obj domain.Budget_Delete;
		detail := &detail_obj;

		detail.ID_user = int(userID)
		detail.Id_bug = id_budget
		
		err =  service.DeleteBudge_Service(detail , h.DB);
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message" : "Failed to delete budget" ,
				"error": err.Error(),
			})
		}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "Budget deleted successfully",
	})
	}
}
