package service

import (
	"github.com/cuzwer/fintrack/internal/repository"
	"gorm.io/gorm"
)


func AccountTable_Service(db *gorm.DB) {
	repository.CreateAccountTable_Repo(db)
}
