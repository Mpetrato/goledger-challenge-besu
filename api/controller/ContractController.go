package controller

import (
	"fmt"

	"github.com/Mpetrato/goledger-challenge-besu/model"
	"github.com/Mpetrato/goledger-challenge-besu/service"
	"github.com/Mpetrato/goledger-challenge-besu/types"
	"github.com/gofiber/fiber/v2"
)

type ContractController struct {
	contractService *service.ContractService
}

func NewContractController(contractService *service.ContractService) *ContractController {
	return &ContractController{contractService}
}

func (c *ContractController) GetContractValue(ctx *fiber.Ctx) error {
	value, err := c.contractService.GetContractValue()
	if err != nil {
		fmt.Println("error on get contract value", err)
		return ctx.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"error": types.ErrorOnGetContractValue})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"contractValue": value})
}

func (c *ContractController) SetContractValue(ctx *fiber.Ctx) error {
	var contract *model.ContractModel

	err := ctx.BodyParser(&contract)
	if err != nil {
		fmt.Println("Error on parse contract", err)
		return ctx.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"error": types.ErrorOnParseContract})
	}

	err = c.contractService.SetContractValue(contract)
	if err != nil {
		fmt.Println("error on set contract value", err)
		return ctx.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"error": types.ErrorOnSetContractValue})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success"})
}

func (c *ContractController) SyncContractValue(ctx *fiber.Ctx) error {
	err := c.contractService.SyncContractValue()
	if err != nil {
		fmt.Println("failed to sync contract value", err)
		return ctx.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"error": types.ErrorOnSyncContractValue})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Success sync contract value"})
}

func (c *ContractController) CheckContractValue(ctx *fiber.Ctx) error {
	result, err := c.contractService.CheckContractValue()
	if err != nil {
		fmt.Println("failed to check contract value error ->", err)
		return ctx.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"error": types.ErrorOnCheckContractValue})
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}
