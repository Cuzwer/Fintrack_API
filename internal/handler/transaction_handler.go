package handler

import (

	"github.com/cuzwer/fintrack/internal/domain"
	"github.com/cuzwer/fintrack/internal/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type TransHandler_obj struct {
	DB *gorm.DB
}

func TransHanler_context(db *gorm.DB ) *TransHandler_obj {
	var Trans TransHandler_obj;
	val := &Trans
	val.DB = db

	return val
}


func(h *TransHandler_obj) PostTrans_Handler() fiber.Handler{
	return  func(c *fiber.Ctx) error{
		newTrans := new(domain.Transaction)	
			
		if err := c.BodyParser(newTrans) ; err != nil { 
  			return  c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid request body",
				"error": err.Error(),
			})
		}
    val := c.Locals("userID")
		usrID , ok := val.(uint)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid user ID",
			})
		}
		err := service.PostTransaction( usrID ,  newTrans , h.DB );
		if err != nil {
  			return  c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to create transaction",
				"error": err.Error(),
			})
		}
    
		return  c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message" : "Transaction created successfully",
		})

	}
}

func(h *TransHandler_obj) GetAllTrans_Handler() fiber.Handler{
	return  func(c *fiber.Ctx) error {
		var Trans_obj []domain.Transaction_SentBack
		Trans := &Trans_obj
		
		val := c.Locals("userID")
		usrID , ok := val.(uint)

		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid user ID",
			})
		}
  	
		err := service.GetTransaction_Service(int(usrID) , Trans , h.DB );
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to retrieve transactions",
				"error": err.Error(),
			})
		}
		
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message" : "Transactions retrieved successfully",
			"value" : Trans,
		})
	}
}
