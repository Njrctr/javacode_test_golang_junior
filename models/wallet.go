package models

import (
	"errors"
	"math"

	"github.com/gofrs/uuid"
)

type Wallet struct {
	Id      int       `json:"-" db:"id"`
	UUID    uuid.UUID `json:"uuid" db:"uuid"`
	Balance int       `json:"balance" db:"balance"`
	Blocked bool      `json:"blocked" db:"blocked"`
}

type WalletUpdate struct {
	WalletUUID    uuid.UUID `json:"walletUUID" binding:"required"`
	OperationType string    `json:"operationType" binding:"required"` // DEPOSIT or WITHDRAW
	Amount        int       `json:"amount" binding:"required"`
}

type BlockWallet struct {
	WalletUUID uuid.UUID `json:"walletUUID" binding:"required"`
	Block      *bool     `json:"block" binding:"required"`
}

func (i *WalletUpdate) Validate() error {
	if i.OperationType != "DEPOSIT" && i.OperationType != "WITHDRAW" {
		return errors.New("неизвестный тип операции")
	}
	if i.Amount < 0 {
		i.Amount = int(math.Abs(float64(i.Amount)))
	}
	return nil
}
