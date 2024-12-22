package pg_rep

import (
	"errors"
	"fmt"

	models "github.com/Njrctr/javacode_test_golang_junior/models"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
)

type AdminPostgres struct {
	db *sqlx.DB
}

func NewAdminPostgres(db *sqlx.DB) *AdminPostgres {
	return &AdminPostgres{db: db}
}

func (r *AdminPostgres) GetByUUID(walletUUID uuid.UUID) (models.Wallet, error) {
	var wallet models.Wallet

	query := fmt.Sprintf("SELECT uuid, balance FROM %s WHERE uuid=$1", walletsTable)
	err := r.db.Get(&wallet, query, walletUUID)

	return wallet, err
}

func (r *AdminPostgres) Update(input models.WalletUpdate) error {
	var setQuery string
	switch input.OperationType {
	case "DEPOSIT":
		setQuery = fmt.Sprintf("balance=balance+%d", input.Amount)
	case "WITHDRAW":
		setQuery = fmt.Sprintf("balance=balance-%d", input.Amount)
	}

	query := fmt.Sprintf("UPDATE %s SET %s WHERE uuid=$1", walletsTable, setQuery)

	_, err := r.db.Exec(query, input.WalletUUID)
	if err != nil && err.Error() == "pq: new row for relation \"wallets\" violates check constraint \"wallets_balance_check\"" {
		return errors.New("недостаточно средств на счете")
	}
	return err
}

func (r *AdminPostgres) BlockWallet(input models.BlockWallet) error {

	query := fmt.Sprintf("UPDATE %s SET blocked=$1 WHERE uuid=$2", walletsTable)
	_, err := r.db.Exec(query, input.Block, input.WalletUUID)
	return err
}
