package repository

import (
	models "github.com/Njrctr/javacode_test_golang_junior/models"
	pg_rep "github.com/Njrctr/javacode_test_golang_junior/pkg/repository/postgres"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
)

type Autorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(email, password string) (models.User, error)
}

type Wallet interface {
	Create(userId int) (uuid.UUID, error)
	GetAll(userId int) ([]models.Wallet, error)
	GetByUUID(walletId uuid.UUID) (models.Wallet, error)
	GetBalanceByUUID(walletId uuid.UUID) (int, error)
	Update(input models.WalletUpdate) error
	Delete(userId int, walletId uuid.UUID) error
}

type Admin interface {
	GetByUUID(walletId uuid.UUID) (models.Wallet, error)
	Update(input models.WalletUpdate) error
	BlockWallet(input models.BlockWallet) error
}

type Repository struct {
	Autorization
	Wallet
	Admin
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Autorization: pg_rep.NewAuthPostgres(db),
		Wallet:       pg_rep.NewWalletPostgres(db),
		Admin:        pg_rep.NewAdminPostgres(db),
	}
}
