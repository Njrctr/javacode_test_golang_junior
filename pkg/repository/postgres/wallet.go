package pg_rep

import (
	"errors"
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

	var walletId int
	var walletUUID uuid.UUID
	createListQuery := fmt.Sprintf("INSERT INTO %s (balance) VALUES (default) RETURNING id, uuid", walletsTable)
	row := tr.QueryRow(createListQuery)
	if err := row.Scan(&walletId, &walletUUID); err != nil {
		tr.Rollback()
		return uuid.Nil, err
	}
	createUsersWalletQuery := fmt.Sprintf("INSERT INTO %s (user_id, wallet_id) values ($1, $2)", usersWalletsTable)
	_, err = tr.Exec(createUsersWalletQuery, userId, walletId)
	if err != nil {
		tr.Rollback()
		return uuid.Nil, err
	}

	return walletUUID, tr.Commit()
}

func (r *WalletPostgres) GetAll(userId int) ([]models.Wallet, error) {
	var lists []models.Wallet

	query := fmt.Sprintf("SELECT w.uuid, w.balance FROM %s w INNER JOIN %s uw on w.id=uw.wallet_id WHERE uw.user_id=$1", walletsTable, usersWalletsTable)
	err := r.db.Select(&lists, query, userId)

	return lists, err
}

func (r *WalletPostgres) GetByUUID(walletId uuid.UUID) (models.Wallet, error) {
	var wallet models.Wallet

	query := fmt.Sprintf("SELECT uuid, balance FROM %s WHERE uuid=$1", walletsTable)
	err := r.db.Get(&wallet, query, walletId)

	return wallet, err
}

func (r *WalletPostgres) GetBalanceByUUID(walletId uuid.UUID) (int, error) {
	var wallet models.Wallet

	query := fmt.Sprintf("SELECT uuid, balance FROM %s WHERE uuid=$1", walletsTable)
	err := r.db.Get(&wallet, query, walletId)

	return wallet.Balance, err
}

func (r *WalletPostgres) Update(input models.WalletUpdate) error {
	var setQuery string
	switch input.OperationType {
	case "DEPOSIT":
		setQuery = fmt.Sprintf("balance=balance+%d", input.Amount)
	case "WITHDRAW":
		setQuery = fmt.Sprintf("balance=balance-%d", input.Amount)
	default:
		return errors.New("неизвестный тип операции OperationType")
	}

	query := fmt.Sprintf("UPDATE %s SET %s WHERE uuid=$1", walletsTable, setQuery)

	_, err := r.db.Exec(query, input.WalletUUID)
	if err.Error() == "pq: new row for relation \"wallets\" violates check constraint \"wallets_balance_check\"" {
		return errors.New("недостаточно средств на счете")
	}
	return err
}

func (r *WalletPostgres) Delete(walleId uuid.UUID) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE uuid=$1", walletsTable)
	_, err := r.db.Exec(query, walleId)
	return err
}
