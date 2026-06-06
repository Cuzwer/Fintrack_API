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
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message" : "Invalid format user id ",
			})
		}
		var newBudget_obj domain.Budgets
		newBudget := &newBudget_obj
		if err := c.BodyParser(newBudget) ; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message" : "Invalid format Post Budgets",
			})
		}
		
		newBudget.ID_user = int(usrId)

		err := service.PostBudget_service(newBudget,h.DB); 
		if err != nil { 
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message" : "failed to post budget" ,
				"status"	: fiber.StatusInternalServerError,
			})
		}
	
		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
			"message" : "Successfully post budget",
			"status" : fiber.StatusAccepted,
		})
	}
}


func(h *Budget_obj) GetBudget_Handler() fiber.Handler{
	return func(c *fiber.Ctx) error { 
		val := c.Locals("userID"); 
		usrId , ok := val.(uint)

		if !ok { 
			c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message" : "Invalid userID format",
			})
		}

		var allBudg_obj []domain.Budgets_detail;
		allBudg := &allBudg_obj;
		err := service.GetBudget_Service(int(usrId) , allBudg , h.DB);
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message" : "failed to get budget something went wrong",
			})
		}

		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
			"message" : "Ok",
			"result" : allBudg,
		})
	}
}

func(h *Budget_obj) UpdateBudget_Handler() fiber.Handler{
	return func(c *fiber.Ctx) error {
		var budg_obj domain.Budgets_Edit
		budg := &budg_obj

		if err := c.BodyParser(budg) ; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message" : "Invalid Input",
			})	
		}
  	
		val := c.Locals("userID");
		userID ,  ok := val.(uint);
		if !ok {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message" : "Invalid user id " ,
			})
		}
 		
		budg.ID_user = int(userID)
		
		err := service.UpdateBudg_Service(budg , h.DB)
		
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message" : "failed to Update something went wrong",
			})
		}
		
		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
			"message" : "Successfully Update",
		})
	}
}

func(h *Budget_obj) DeleteBudget() fiber.Handler{
	return  func(c *fiber.Ctx) error { 
		val := c.Locals("userID");
  	userID , ok := val.(uint);

		if !ok {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message" : "Invalid userID please login respectfully",
			})
		}

		id_budget , err := strconv.Atoi(c.Params("id"));

		if err != nil { 
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message" : "Invalid Id Budget",
			})
		}
		
		var detail_obj domain.Budget_Delete;
		detail := &detail_obj;

		detail.ID_user = int(userID)
		detail.Id_bug = id_budget
		
		err =  service.DeleteBudge_Service(detail , h.DB);
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message" : "failed to Deleted Budgets" ,
			})
		}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "Successfully Delete Budget",
	})
	}
}
