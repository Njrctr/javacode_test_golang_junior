package models

type Wallet struct {
}

type WalletUpdate struct {
	ValletId      int
	OperationType string // DEPOSIT or WITHDRAW
	Amount        int
}
