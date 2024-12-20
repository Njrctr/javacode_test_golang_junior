package models

import "github.com/gofrs/uuid"

type Wallet struct {
	Id      int       `json:"-" db:"id"`
	UUID    uuid.UUID `json:"uuid" db:"uuid"`
	Balance int       `json:"balance" db:"balance"`
}

type WalletUpdate struct {
	WalletUUID    uuid.UUID `json:"walletId" binding:"required"`
	OperationType string    `json:"operationType" binding:"required"` // DEPOSIT or WITHDRAW
	Amount        int       `json:"amount" binding:"required"`
}
