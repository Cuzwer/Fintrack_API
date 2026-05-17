package database

import (
	"fmt"
	"log"

	"github.com/cuzwer/fintrack/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)



func ConnectDB(cfg *config.DBConfig) (*gorm.DB){
   dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
   cfg.DBHOST,
   cfg.DBUSER,
	 cfg.DBPASSWORD,
	 cfg.DBNAME,
	 cfg.DBPORT,
	 cfg.DBSSLMODE,
 )
   db , err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	 if err != nil { 
		log.Fatalf("fail to connect to database %v", err)
	 }
  	fmt.Printf("success fully ConnectDB %v", err)
	
  return  db
}
