package database

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)


func RunDatabaseMigrations(dbURL string) { 

  m , err := migrate.New("file://./migrations", dbURL)
 	
	if err != nil { 
	 		log.Fatalf("\nFailed to create migrate instance : %v", err)
	}
 	
	if err := m.Up() ; err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Faile to run migrateup : %v", err)
	}
  fmt.Println("\nData migrated succesfully ! 💀🙏😭 ") 
}
