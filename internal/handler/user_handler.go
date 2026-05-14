package handler

import (

	"github.com/cuzwer/fintrack/internal/service"
	_ "github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)



func CreateUsertable_Handler (db *gorm.DB) {
	service.CreateUser_Service(db)
}
