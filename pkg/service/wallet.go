package service

import (
	wallet "github.com/Njrctr/javacode_test_golang_junior/models"
	"github.com/Njrctr/javacode_test_golang_junior/pkg/repository"
)

type WalletService struct {
	repo repository.Wallet
}

func NewWalletService(repo repository.Wallet) *WalletService {
	return &WalletService{repo: repo}
}

func (s *WalletService) Create(userId int) (int, error) {
	return s.repo.Create(userId)
}

func (s *WalletService) GetAll(userId int) ([]wallet.Wallet, error) {
	return s.repo.GetAll(userId)
}

func (s *WalletService) Delete(userId, listId int) error {
	return s.repo.Delete(userId, listId)
}

func (s *WalletService) Update(userId, listId int, input wallet.WalletUpdate) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, listId, input)
}
