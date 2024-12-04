package main

import (
	"github.com/Mpetrato/goledger-challenge-besu/database"
	"github.com/Mpetrato/goledger-challenge-besu/router"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func main() {
	db, err := database.InitDatabase()
	if err != nil {
		panic(err)
	}

	app := fiber.New()

	setupRoutes(app, db)

	app.Listen(":3000")
}

func setupRoutes(app *fiber.App, db *gorm.DB) {
	api := app.Group("/api/v1")

	contractGroup := api.Group("/contract")
	router.InitContractRouter(contractGroup, db)
}
