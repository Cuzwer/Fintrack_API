package handler

import (
	"strconv"

	"github.com/cuzwer/fintrack/internal/domain"
	"github.com/cuzwer/fintrack/internal/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AccountDB_handler struct { 
	DB *gorm.DB
}


func AccountConts_handler(db *gorm.DB) *AccountDB_handler {
	Account_Interface := &AccountDB_handler{}
	Account_Interface.DB = db

	return Account_Interface
}

func (h *AccountDB_handler) GetAllAccount() fiber.Handler{
	return  func(c *fiber.Ctx) error { 
		AllAccount := new([]domain.Detail_Account)
		id_account,err := strconv.Atoi(c.Params("id"))
		
		if err != nil { 
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid account ID format",
				"error": err.Error(),
			})
		}

		err = service.GetAccountAll_Service(id_account, AllAccount , h.DB)
		if err != nil { 
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to retrieve accounts",
				"error": err.Error(),
			})
		}
		
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Accounts retrieved successfully",
			"data": AllAccount,
		})
}
}

func (h *AccountDB_handler) PostAccount_Handler() fiber.Handler{
	return func (c *fiber.Ctx ) error {
		
		var Newaccount = &domain.Detail_Account{}
		if err := c.BodyParser(Newaccount) ; err != nil {
			return  c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid request body",
				"error": err.Error(),
			})
		}
		err := service.PostAccount_Service(Newaccount,h.DB)

		if err != nil { 
			return  c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to create account",
				"error": err.Error(),
			})
		}
	
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message" : "Account created successfully",
		})
	}
}
