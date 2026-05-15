package service

import (
	"github.com/cuzwer/fintrack/internal/repository"
	"gorm.io/gorm"
)

func  CreatedCategory_Service(db *gorm.DB) {
	repository.CreateCategories_Repo(db)
}
