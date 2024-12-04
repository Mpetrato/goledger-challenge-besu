package service

import (
	"fmt"

	"github.com/Mpetrato/goledger-challenge-besu/model"
	"github.com/Mpetrato/goledger-challenge-besu/repository"
)

type ContractService struct {
	contractRepository *repository.ContractRepository
	besuService        *BesuService
}

func NewContractService(
	contractRepository *repository.ContractRepository,
	besuService *BesuService,
) *ContractService {
	return &ContractService{contractRepository, besuService}
}

func (s *ContractService) GetContractValue() (*uint64, error) {
	value, err := s.besuService.GetBesuContractValue()
	if err != nil {
		fmt.Println("Error on GetContractValue -> ", err)
		return nil, err
	}

	return value, nil
}

func (s *ContractService) SetContractValue(contract *model.ContractModel) error {
	err := s.besuService.SetBesuContractValue(contract.Value)
	if err != nil {
		fmt.Println("Error on SetContractValue -> ", err)
		return err
	}

	return nil
}

func (s *ContractService) CheckContractValue() (bool, error) {
	contractDbValue, err := s.contractRepository.GetContractValue()
	if err != nil {
		fmt.Println("Error on CheckContractValue -> ", err)
		return false, err
	}

	contractBesuValue, err := s.besuService.GetBesuContractValue()
	if err != nil {
		fmt.Println("Error on CheckContractValue -> ", err)
		return false, err
	}

	fmt.Println(contractDbValue.Value, *contractBesuValue)
	if contractDbValue.Value == *contractBesuValue {
		return true, nil
	}

	return false, nil
}

func (s *ContractService) SyncContractValue() error {
	value, err := s.besuService.GetBesuContractValue()
	if err != nil {
		fmt.Println("Error on SyncContractValue -> ", err)
		return err
	}

	contractDbValue, err := s.contractRepository.GetContractValue()
	if err != nil {
		fmt.Println("Error on SyncContractValue -> ", err)
		return err
	}

	contractDbValue.Value = *value

	err = s.contractRepository.SetContractValue(contractDbValue)
	if err != nil {
		fmt.Println("Error on SyncContractValue -> ", err)
		return err
	}

	return nil
}
