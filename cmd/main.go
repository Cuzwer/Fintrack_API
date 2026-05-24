package main

import (
	"fmt"
	_ "fmt"

	_ "github.com/cuzwer/fintrack/internal/handler"
	"github.com/cuzwer/fintrack/internal/middleware"
	"github.com/cuzwer/fintrack/internal/routes"
	"github.com/cuzwer/fintrack/pkg/config"
	"github.com/cuzwer/fintrack/pkg/database"
	"github.com/gofiber/fiber/v2"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
)

func main() {

	// section Prepare

	app := fiber.New()
	cfg := config.LoadConfig()
	PORT := ":" + cfg.PORT

	db := database.ConnectDB(cfg)

	if db.Error != nil {
		fmt.Printf("Cant ConnectDB  %v", db.Error)
	}

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.DBUSER,
		cfg.DBPASSWORD,
		cfg.DBHOST,
		cfg.DBPORT,
		cfg.DBNAME,
		cfg.DBSSLMODE,
	)
	database.RunDatabaseMigrations(dbURL)
	// Middle ware
	app.Use(middleware.NewCorsMiddleware())

	// section API
	routes.SetUpRoutest(app, db)
	fmt.Printf("Server is running on port %v 🙏💀", PORT)
	app.Listen(PORT)
}
