package service

import (
	models "github.com/Njrctr/javacode_test_golang_junior/models"
	"github.com/Njrctr/javacode_test_golang_junior/pkg/repository"
	"github.com/gofrs/uuid"
)

type AdminService struct {
	repo repository.Admin
}

func NewAdminService(repo repository.Admin) *AdminService {
	return &AdminService{repo: repo}
}

func (c *AdminService) GetByUUID(walletUUID uuid.UUID) (models.Wallet, error) {
	return c.repo.GetByUUID(walletUUID)
}
func (c *AdminService) Update(input models.WalletUpdate) error {
	return c.repo.Update(input)
}
func (c *AdminService) BlockWallet(input models.BlockWallet) error {
	return c.repo.BlockWallet(input)
}
