package service

import (
	"github.com/cuzwer/fintrack/internal/repository"
	"gorm.io/gorm"
)


func CreateUser_Service (db *gorm.DB) { 
	repository.CreateTableUser_Repo(db)
}
