package main

import (
	"fmt"

	"github.com/Mpetrato/goledger-challenge-besu/database"
	"github.com/Mpetrato/goledger-challenge-besu/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"
)

func main() {
	db, err := database.InitDatabase()
	if err != nil {
		fmt.Printf("error on initDatabase -> %v", err)
		panic(err)
	}

	app := fiber.New()
	app.Use(logger.New())

	setupRoutes(app, db)

	app.Listen(":3000")
}

func setupRoutes(app *fiber.App, db *gorm.DB) {
	api := app.Group("/api/v1")

	contractGroup := api.Group("/contract")
	router.InitContractRouter(contractGroup, db)
}
