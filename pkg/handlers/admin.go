package handler

import (
	"log"
	"net/http"
	"strconv"

	models "github.com/Njrctr/javacode_test_golang_junior/models"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

type CreateWallet struct {
	UserId int `json:"user_id" binding:"required"`
}

// @Summary Create Wallet
// @Security ApiKeyAuth
// @Tags ADMIN
// @Description create wallet
// @ID create-wallet-admin
// @Accept  json
// @Produce  json
// @Param input body CreateWallet true "User ID"
// @Success 200 {uuid} uuid 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/admin/wallet/new [post]
func (h *Handler) createWalletToUser(c *gin.Context) {

	var userId CreateWallet
	if err := c.BindJSON(&userId); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	uuid, err := h.services.Wallet.Create(userId.UserId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"uuid": uuid,
	})
}

type getAllWalletsAdminResponce struct {
	Data []models.Wallet `json:"data"`
}

// @Summary Get All Wallets By User ID
// @Security ApiKeyAuth
// @Tags ADMIN
// @Description get all wallets by user id
// @ID get-all-wallets-admin
// @Accept  json
// @Produce  json
// @Param user_id path string true "User ID"
// @Success 200 {object} getAllWalletsAdminResponce
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/admin/wallet/{user_id} [get]
func (h *Handler) getAllWalletsByUser(c *gin.Context) {

	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid user_id param")
		return
	}

	wallets, err := h.services.Wallet.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllWalletsAdminResponce{
		Data: wallets,
	})
}

// @Summary Get Wallet By UUID
// @Security ApiKeyAuth
// @Tags ADMIN
// @Description get wallet by uuid
// @ID get-wallet-by-uuid-admin
// @Accept  json
// @Produce  json
// @Param wallet_uuid path string true "Wallet uuid"
// @Success 200 {object} models.Wallet
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/admin/wallets/{wallet_uuid} [get]
func (h *Handler) getWalletByUUID(c *gin.Context) {

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

// @Summary Update Wallet
// @Security ApiKeyAuth
// @Tags ADMIN
// @Description update wallet
// @ID update-wallet-admin
// @Accept  json
// @Produce  json
// @Param input body models.WalletUpdate true "Wallet query"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/admin/wallet [post]
func (h *Handler) updateWalletAdmin(c *gin.Context) {

	var input models.WalletUpdate
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if input.OperationType != "DEPOSIT" && input.OperationType != "WITHDRAW" {
		newErrorResponse(c, http.StatusBadRequest, "Неизвестный тип операции OperationType")
		return

	}

	if err := h.services.Wallet.Update(input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponce{"ok"})
}

// @Summary Block Wallet
// @Security ApiKeyAuth
// @Tags ADMIN
// @Description block wallet
// @ID block-wallet-admin
// @Accept  json
// @Produce  json
// @Param input body models.BlockWallet true "Block Wallet"
// @Success 200 {string} string ok
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/v1/admin/wallet/block [put]
func (h *Handler) blockWallet(c *gin.Context) {

	adminId, err := getUserId(c)
	if err != nil {
		return
	}

	var input models.BlockWallet
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if input.Block == nil {
		newErrorResponse(c, http.StatusInternalServerError, "Поле block обязательно!")
		return
	}
	err = h.services.Admin.BlockWallet(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponce{
		Status: "ok",
	})
	if *input.Block {
		log.Println("Admin", adminId, "block wallet", input.WalletUUID)
	} else if !*input.Block {
		log.Println("Admin", adminId, "UNblock wallet", input.WalletUUID)
	}
}
