package models

import "github.com/gofrs/uuid"

type Wallet struct {
	Id int `json:"-" db:"id"`
}

type WalletUpdate struct {
	WalletUUID    uuid.UUID `json:"wallet_id" binding:"required"`
	OperationType string    `json:"operation_type" binding:"required"` // DEPOSIT or WITHDRAW
	Amount        int       `json:"amount" binding:"required"`
}
