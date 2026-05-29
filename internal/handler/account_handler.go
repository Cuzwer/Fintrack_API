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
			return c.Status(fiber.StatusUnsupportedMediaType).SendString(err.Error())	}

		err = service.GetAccountAll_Service(id_account, AllAccount , h.DB)
		if err != nil { 
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		
		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
			"message " : "success fully get Account",
			"data " : AllAccount,
		})
}
}


func (h *AccountDB_handler) PostAccount_Handler() fiber.Handler{
	return func (c *fiber.Ctx ) error {
		
		var Newaccount = &domain.Detail_Account{}
		if err := c.BodyParser(Newaccount) ; err != nil {
			return  c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		err := service.PostAccount_Service(Newaccount,h.DB)

		if err != nil { 
			return  c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
	
		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
			"message" : fiber.StatusAccepted,
		})
	}
}
