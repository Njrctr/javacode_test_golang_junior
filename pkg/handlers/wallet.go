package handler

import (
	"net/http"

	models "github.com/Njrctr/javacode_test_golang_junior/models"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

// @Summary Create Wallet
// @Security ApiKeyAuth
// @Tags Wallets
// @Description create wallet
// @ID create-wallet
// @Accept  json
// @Produce  json
// @Success 200 {uuid} uuid 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/wallet/new [post]
func (h *Handler) createWallet(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	uuid, err := h.services.Wallet.Create(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"uuid": uuid,
	})
}

type getAllWalletsResponce struct {
	Data []models.Wallet `json:"data"`
}

// @Summary Get All Wallets
// @Security ApiKeyAuth
// @Tags Wallets
// @Description get all wallets
// @ID get-all-wallets
// @Accept  json
// @Produce  json
// @Success 200 {object} getAllWalletsResponce
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/wallet [get]
func (h *Handler) getAllWallets(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	wallets, err := h.services.Wallet.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllWalletsResponce{
		Data: wallets,
	})
}

// @Summary Get Wallet By UUID
// @Security ApiKeyAuth
// @Tags Wallets
// @Description get wallet by uuid
// @ID get-wallet-by-uuid
// @Accept  json
// @Produce  json
// @Param wallet_uuid path string true "Wallet uuid"
// @Success 200 {object} models.Wallet
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/wallets/{wallet_uuid} [get]
func (h *Handler) getWalletById(c *gin.Context) {

	walletId, err := uuid.FromString(c.Param("wallet_uuid"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid wallet_uuid param")
		return
	}

	wallet, err := h.services.Wallet.GetByUUID(walletId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, wallet)
}

func (h *Handler) GetBalanceByUUID(c *gin.Context) {

	walletId, err := uuid.FromString(c.Param("wallet_uuid"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid wallet_uuid param")
		return
	}

	walletBalance, err := h.services.Wallet.GetBalanceByUUID(walletId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, walletBalance)
}

// @Summary Update Wallet
// @Security ApiKeyAuth
// @Tags Wallets
// @Description update wallet
// @ID update-wallet
// @Accept  json
// @Produce  json
// @Param input body models.WalletUpdate true "Wallet query"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/wallet [post]
func (h *Handler) updateWallet(c *gin.Context) {

	var input models.WalletUpdate
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Wallet.Update(input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponce{"ok"})
}

// @Summary Delete Wallet
// @Security ApiKeyAuth
// @Tags Wallets
// @Description delete wallet
// @ID delete-wallet
// @Accept  json
// @Produce  json
// @Param wallet_uuid path string true "Wallet UUID"
// @Success 200 {string} string ok
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/wallet/{wallet_uuid} [delete]
func (h *Handler) deleteWallet(c *gin.Context) {

	walletId, err := uuid.FromString(c.Param("wallet_uuid"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid wallet_uuid param")
		return
	}

	err = h.services.Wallet.Delete(walletId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponce{
		Status: "ok",
	})
}
