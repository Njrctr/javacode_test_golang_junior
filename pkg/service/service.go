package service

import (
	models "github.com/Njrctr/javacode_test_golang_junior/models"
	"github.com/Njrctr/javacode_test_golang_junior/pkg/repository"
	"github.com/gofrs/uuid"
)

//go:generate mockgen -source=service.go -destination=mock/mock.go

type Autorization interface {
	CreateUser(user models.User) (int, error)
	GenerateJWTToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Wallet interface {
	Create(userId int) (uuid.UUID, error)
	GetAll(userId int) ([]models.Wallet, error)
	GetById(walletId uuid.UUID) (models.Wallet, error)
	Update(input models.WalletUpdate) error
	Delete(walletId uuid.UUID) error
}

type Service struct {
	Autorization
	Wallet
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Autorization: NewAuthService(repos.Autorization),
		Wallet:       NewWalletService(repos.Wallet),
	}
}
