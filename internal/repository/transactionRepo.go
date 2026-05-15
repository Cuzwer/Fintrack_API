package repository

import (
	"fmt"
	"log"

	"github.com/cuzwer/fintrack/internal/domain"
	"gorm.io/gorm"
)

func Transaction_Repo (db *gorm.DB) { 
	transaction := new(domain.Transaction);
	
	if db.Migrator().HasTable(transaction) != false {
		fmt.Println("Transaction Already Craeted long before")
		return
	}

	query := `CREATE TABLE transactions (
  id_trans         SERIAL PRIMARY KEY,
  id_account       INT NOT NULL REFERENCES accounts(id_account) ON DELETE CASCADE,
  id_category      INT REFERENCES categories(id_category),
  amount_trans     DECIMAL(15, 2) NOT NULL,
  type_trans       VARCHAR(20) NOT NULL CHECK (type_trans IN ('income', 'expense')),
  descrip_trans    VARCHAR(255),
  transaction_date DATE NOT NULL DEFAULT CURRENT_DATE,
  created_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`
  if err := db.Exec(query).Error ; err != nil { 
		log.Fatalf("Faild to create a transaction table [%v]" ,  err)
	}


	fmt.Println("create transaction table succesfullty 🔥")
}
