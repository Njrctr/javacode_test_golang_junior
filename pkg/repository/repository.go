package repository

import (
	models "github.com/Njrctr/javacode_test_golang_junior/models"
	pg_rep "github.com/Njrctr/javacode_test_golang_junior/repository/postgres"
	"github.com/jmoiron/sqlx"
)

type Autorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(email, password string) (models.User, error)
}

type Wallet interface {
	Create(userId int) (int, error)
	GetAll(userId int) ([]models.Wallet, error)
	GetById(userId, listId int) (models.Wallet, error)
	Delete(userId, listId int) error
	Update(userId, walletId int, input models.WalletUpdate) error
}

type Repository struct {
	Autorization
	Wallet
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Autorization: pg_rep.NewAuthRepository(db),
		Wallet:       pg_rep.NewWalletRepository(db),
	}
}
