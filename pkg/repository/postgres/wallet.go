package pg_rep

import (
	"fmt"

	models "github.com/Njrctr/javacode_test_golang_junior/models"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
)

type WalletPostgres struct {
	db *sqlx.DB
}

func NewWalletPostgres(db *sqlx.DB) *WalletPostgres {
	return &WalletPostgres{db: db}
}

func (r *WalletPostgres) Create(userId int) (uuid.UUID, error) {
	tr, err := r.db.Begin() //* Старт транзакции
	if err != nil {
		return uuid.Nil, err
	}

	var walletId uuid.UUID
	createListQuery := fmt.Sprintf("INSERT INTO %s (uuid) values ($1) RETURNING id", walletsTable)
	row := tr.QueryRow(createListQuery)
	if err := row.Scan(&walletId); err != nil {
		tr.Rollback()
		return uuid.Nil, err
	}
	createUsersWalletQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) values ($1, $2)", usersWalletsTable)
	_, err = tr.Exec(createUsersWalletQuery, userId, walletId)
	if err != nil {
		tr.Rollback()
		return uuid.Nil, err
	}

	return walletId, tr.Commit()
}

func (r *WalletPostgres) GetAll(userId int) ([]models.Wallet, error) {
	var lists []models.Wallet

	query := fmt.Sprintf("SELECT uuid, balance FROM %s WHERE user_id=$1", walletsTable)
	err := r.db.Select(&lists, query, userId)

	return lists, err
}

func (r *WalletPostgres) GetByUUID(walletId uuid.UUID) (models.Wallet, error) {
	var wallet models.Wallet

	query := fmt.Sprintf("SELECT uuid, balance FROM %s WHERE uuid=$1", walletsTable)
	err := r.db.Get(&wallet, query, walletId)

	return wallet, err
}

func (r *WalletPostgres) Update(input models.WalletUpdate) error {
	var setQuery string

	if input.OperationType == "DEPOSIT" {
		setQuery = fmt.Sprintf("balance=balance+$%d", input.Amount)
	}
	if input.OperationType == "WITHDRAW" {
		setQuery = fmt.Sprintf("balance=balance-$%d", input.Amount)
	}

	query := fmt.Sprintf("UPDATE %s SET %s WHERE uuid=$1", walletsTable, setQuery)

	_, err := r.db.Exec(query, input.WalletUUID)
	return err
}

func (r *WalletPostgres) Delete(walleId uuid.UUID) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE uuid=$1", walletsTable)
	_, err := r.db.Exec(query, walleId)
	return err
}
