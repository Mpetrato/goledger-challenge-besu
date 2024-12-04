package main

import (
	"fmt"

	"github.com/Mpetrato/goledger-challenge-besu/database"
	"github.com/Mpetrato/goledger-challenge-besu/helpers"
	"github.com/Mpetrato/goledger-challenge-besu/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"
)

func main() {
	loadEnvs()
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

func loadEnvs() {
	_, err := helpers.GetOSEnv("CONTRACT_ADDRESS")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	_, err = helpers.GetOSEnv("PRIVATE_KEY")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
