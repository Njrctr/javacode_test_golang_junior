package service

import (
	models "github.com/Njrctr/javacode_test_golang_junior/models"
	"github.com/Njrctr/javacode_test_golang_junior/pkg/repository"
	"github.com/gofrs/uuid"
)

//go:generate mockgen -source=service.go -destination=mock/mock.go

type Autorization interface {
	CreateUser(user models.SignUpInput) (int, error)
	GenerateJWTToken(username, password string) (string, error)
	ParseToken(token string) (int, bool, error)
}

type Wallet interface {
	Create(userId int) (uuid.UUID, error)
	GetAll(userId int) ([]models.Wallet, error)
	GetByUUID(walletUUID uuid.UUID) (models.Wallet, error)
	GetBalanceByUUID(walletUUID uuid.UUID) (int, error)
	Update(input models.WalletUpdate) error
	Delete(userId int, walletUUID uuid.UUID) error //todo: Сделать Удаление возможным только при нулевом балансе
}

type Admin interface {
	Update(input models.WalletUpdate) error
	BlockWallet(input models.BlockWallet) error
}

type Service struct {
	Autorization
	Wallet
	Admin
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Autorization: NewAuthService(repos.Autorization),
		Wallet:       NewWalletService(repos.Wallet),
		Admin:        NewAdminService(repos.Admin),
	}
}
