package repository

import (
	"fmt"
	"log"

	"github.com/cuzwer/fintrack/internal/domain"
	"gorm.io/gorm"
)



func CreateTableUser_Repo(db *gorm.DB) {
  
	user := &domain.User{}
	query := `CREATE TABLE users (
  id_user SERIAL PRIMARY KEY,
  email_user VARCHAR(100) NOT NULL,
  password_hash VARCHAR(100) NOT NULL,
  created_at DATE DEFAULT CURRENT_DATE
);`
  
  checkTable := db.Migrator().HasTable(user)
  
	if checkTable != false { 
		fmt.Printf("user already created [ %v ] ",  checkTable)
		return
	}
  
	err := db.Exec(query).Error
	if err != nil { log.Fatalf("Cant Create Database : %v", err) }

  fmt.Printf("Success fully Create User Table \n")
}
