package model

type ContractModel struct {
	ID    uint   `gorm:"primaryKey"`
	Value uint64 `json:"contract_value"`
}
