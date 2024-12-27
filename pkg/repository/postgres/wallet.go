package pg_rep

import (
	"errors"
	"fmt"

	models "github.com/Njrctr/javacode_test_golang_junior/models"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
)

const (
	errorIncorrectUserId = "pq: insert or update on table \"users_wallets\" violates foreign key constraint \"users_wallets_user_id_fkey\""
	errorBalance         = "pq: new row for relation \"wallets\" violates check constraint \"wallets_balance_check\""
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
		if err.Error() == errorIncorrectUserId {
			return uuid.Nil, errors.New("недопустимое значение поля user_id")
		}
		return uuid.Nil, err
	}

	return walletUUID, tr.Commit()
}

func (r *WalletPostgres) GetAll(userId int) ([]models.Wallet, error) {
	var lists []models.Wallet

	query := fmt.Sprintf("SELECT w.uuid, w.balance, w.blocked FROM %s w INNER JOIN %s uw on w.id=uw.wallet_id WHERE uw.user_id=$1", walletsTable, usersWalletsTable)
	err := r.db.Select(&lists, query, userId)

	return lists, err
}

func (r *WalletPostgres) GetByUUID(walletUUID uuid.UUID) (models.Wallet, error) {
	var wallet models.Wallet

	query := fmt.Sprintf("SELECT uuid, balance, blocked FROM %s WHERE uuid=$1", walletsTable)
	err := r.db.Get(&wallet, query, walletUUID)

	return wallet, err
}

func (r *WalletPostgres) GetBalanceByUUID(walletUUID uuid.UUID) (int, error) {
	var wallet models.Wallet

	query := fmt.Sprintf("SELECT uuid, balance FROM %s WHERE uuid=$1", walletsTable)
	err := r.db.Get(&wallet, query, walletUUID)

	return wallet.Balance, err
}

func (r *WalletPostgres) Update(input models.WalletUpdate) error {
	var setQuery string
	setQuery = fmt.Sprintf("balance=balance+%d", input.Amount)

	if input.OperationType == "WITHDRAW" {
		setQuery = fmt.Sprintf("balance=balance-%d", input.Amount)
	}

	query := fmt.Sprintf("UPDATE %s SET %s WHERE uuid=$1", walletsTable, setQuery)

	_, err := r.db.Exec(query, input.WalletUUID)
	if err != nil && err.Error() == errorBalance {
		return errors.New("недостаточно средств на счете")
	}
	return err
}

func (r *WalletPostgres) Delete(userId int, walletUUID uuid.UUID) error {
	walletIdQuery := fmt.Sprintf("SELECT * FROM %s WHERE uuid=$1", walletsTable)
	var wallet models.Wallet
	err := r.db.Get(&wallet, walletIdQuery, walletUUID)
	if err != nil {
		return err
	}

	if wallet.Balance != 0 {
		return errors.New("невозможно удалить не пустой кошелёк")
	}
	if wallet.Blocked {
		return errors.New("невозможно удалить заблокированный колешёк. Обратитесь к администрации")
	}

	query := fmt.Sprintf("DELETE FROM %s w USING %s uw WHERE uw.wallet_id=w.id AND uw.user_id=$1 AND uw.wallet_id=$2", walletsTable, usersWalletsTable)
	_, err = r.db.Exec(query, userId, wallet.Id)
	return err
}
