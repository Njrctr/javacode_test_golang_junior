package service

import (
	models "github.com/Njrctr/javacode_test_golang_junior/models"
	"github.com/Njrctr/javacode_test_golang_junior/pkg/repository"
	"github.com/gofrs/uuid"
)

type WalletService struct {
	repo repository.Wallet
}

func NewWalletService(repo repository.Wallet) *WalletService {
	return &WalletService{repo: repo}
}

func (s *WalletService) Create(userId int) (uuid.UUID, error) {
	return s.repo.Create(userId)
}

func (s *WalletService) GetAll(userId int) ([]models.Wallet, error) {
	return s.repo.GetAll(userId)
}

func (s *WalletService) GetByUUID(walletUUID uuid.UUID) (models.Wallet, error) {
	return s.repo.GetByUUID(walletUUID)
}

func (s *WalletService) GetBalanceByUUID(walletUUID uuid.UUID) (int, error) {
	return s.repo.GetBalanceByUUID(walletUUID)
}

func (s *WalletService) Delete(userId int, walletUUID uuid.UUID) error {
	return s.repo.Delete(userId, walletUUID)
}

func (s *WalletService) Update(input models.WalletUpdate) error {
	// if err := input.Validate(); err != nil {
	// 	return err
	// }
	return s.repo.Update(input)
}
