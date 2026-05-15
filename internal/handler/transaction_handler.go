package handler

import (
	"github.com/cuzwer/fintrack/internal/service"
	"gorm.io/gorm"
)

func CreateTransaction_handler (db *gorm.DB) {
	service.CreateTransaction(db)
}
