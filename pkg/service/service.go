package service

import (
	models "github.com/Njrctr/javacode_test_golang_junior/models"
	"github.com/Njrctr/javacode_test_golang_junior/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mock/mock.go

type Autorization interface {
	CreateUser(user models.User) (int, error)
	GenerateJWTToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Wallet interface {
	Create(models.Wallet) (int, error)
	Update(valletId int, operationType string, amount int) error
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
