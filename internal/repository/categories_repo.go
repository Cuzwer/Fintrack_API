package repository

import (
	"fmt"
	"log"

	"github.com/cuzwer/fintrack/internal/domain"
	"gorm.io/gorm"
)


func CreateCategories_Repo(db *gorm.DB){
	var category_object domain.Categories
	category := &category_object


	if db.Migrator().HasTable(category) != false {
		fmt.Println("Categories Table Already Created")
		return
	}

	query := `CREATE TABLE categories (
  id_category   SERIAL PRIMARY KEY,
  id_user       INT NOT NULL REFERENCES users(id_user) ON DELETE CASCADE,
  name_category VARCHAR(100) NOT NULL,
  type_category VARCHAR(20) NOT NULL CHECK (type_category IN ('income', 'expense'))
);`
	
	err := db.Exec(query).Error
	if err != nil { 
		log.Fatalf("failed to created Categories Table %v" , err)
	}

	fmt.Printf("Success fully Created Categories Table [%v]",err)
}
