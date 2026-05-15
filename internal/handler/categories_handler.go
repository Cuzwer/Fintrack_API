package handler

import (
	"github.com/cuzwer/fintrack/internal/service"
	"gorm.io/gorm"
)


func CreateCategory_handler(db *gorm.DB) {
	service.CreatedCategory_Service(db)
}
