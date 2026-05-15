package handler

import (
	"github.com/cuzwer/fintrack/internal/service"
	"gorm.io/gorm"
)




func AccountTable_handler(db *gorm.DB) {
	service.AccountTable_Service(db)
}
