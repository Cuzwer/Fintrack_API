package main

import (
	"fmt"
	_ "fmt"

	"github.com/cuzwer/fintrack/internal/handler"
	"github.com/cuzwer/fintrack/pkg/config"
	"github.com/cuzwer/fintrack/pkg/database"
	"github.com/gofiber/fiber/v2"
)

func main () { 
	app := fiber.New()
	cfg := config.LoadConfig()
	PORT := ":"+cfg.PORT
  
	db := database.ConnectDB(cfg)
  
	if db.Error != nil {fmt.Printf("Cant ConnectDB  %v", db.Error)}
	
	handler.CreateUsertable_Handler(db) 
	handler.AccountTable_handler(db)
  app.Get("/api/hello", func (c *fiber.Ctx) error {
		return c.Status(fiber.StatusAccepted).SendString("GoodBye")
	})



	app.Listen(PORT)
}
