package router

import (
	"fmt"

	"github.com/Mpetrato/goledger-challenge-besu/controller"
	"github.com/Mpetrato/goledger-challenge-besu/repository"
	"github.com/Mpetrato/goledger-challenge-besu/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InitContractRouter(storeGroup fiber.Router, db *gorm.DB) {
	repository := repository.NewContractRepository(db)
	besuService := service.NewBesuService()
	service := service.NewContractService(repository, besuService)
	controller := controller.NewContractController(service)

	storeGroup.Post("/", controller.SetContractValue)

	storeGroup.Get("/", controller.GetContractValue)

	storeGroup.Post("/sync", controller.SyncContractValue)

	storeGroup.Get("/check", controller.CheckContractValue)

	fmt.Println("Contract routes live")
}
