package service

import (
	"github.com/cuzwer/fintrack/internal/repository"
	"gorm.io/gorm"
)



func CreateTransaction(db *gorm.DB) {
	repository.Transaction_Repo(db)
}
