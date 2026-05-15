package repository

import (
	"fmt"
	"log"

	"github.com/cuzwer/fintrack/internal/domain"
	"gorm.io/gorm"
)




func CreateAccountTable_Repo(db *gorm.DB) { 
	account := &domain.Account{}
	if db.Migrator().HasTable(account)  != false { 
		fmt.Println("Table Account Already create")
		return
	}
  query := `CREATE TABLE accounts (
  id_account   SERIAL PRIMARY KEY,
  id_user      INT NOT NULL REFERENCES users(id_user) ON DELETE CASCADE,
  name         VARCHAR(100) NOT NULL,
  type_account VARCHAR(50) NOT NULL,
  balance      DECIMAL(15, 2) DEFAULT 0.00,
  currency     VARCHAR(10) DEFAULT 'THB'
);`


	if err := db.Exec(query).Error ; err != nil { 
		log.Fatalf("Some thing went wrong [%v]",err )	
	}
	
	fmt.Print("Table Account success to create")
}
