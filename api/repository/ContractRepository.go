package repository

import (
	"fmt"

	"github.com/Mpetrato/goledger-challenge-besu/model"
	"gorm.io/gorm"
)

type ContractRepository struct {
	db *gorm.DB
}

func NewContractRepository(db *gorm.DB) *ContractRepository {
	return &ContractRepository{db}
}

func (r *ContractRepository) GetContractValue() (*model.ContractModel, error) {
	var contractModel *model.ContractModel

	err := r.db.First(&contractModel).Error
	if err != nil {
		fmt.Println("Error on GetContractValue -> ", err)
		return nil, err
	}

	return contractModel, nil
}

func (r *ContractRepository) SetContractValue(contract *model.ContractModel) error {
	err := r.db.Save(&contract).Error
	if err != nil {
		fmt.Println("Error on SetContractValue -> ", err)
		return err
	}

	return nil
}
