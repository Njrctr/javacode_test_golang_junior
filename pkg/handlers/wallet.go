package handler

import (
	"net/http"

	models "github.com/Njrctr/javacode_test_golang_junior/models"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

// @Summary Create todo list
// @Security ApiKeyAuth
// @Tags lists
// @Description create todo list
// @ID create-list
// @Accept  json
// @Produce  json
// @Param input body models.TodoListCreateUpdate true "list info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists [post]
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

// @Summary Get All Lists
// @Security ApiKeyAuth
// @Tags lists
// @Description get all lists
// @ID get-all-lists
// @Accept  json
// @Produce  json
// @Success 200 {object} getAllListsResponce
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists [get]
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

// @Summary Get List By Id
// @Security ApiKeyAuth
// @Tags lists
// @Description get list by id
// @ID get-list-by-id
// @Accept  json
// @Produce  json
// @Param list_id path int true "List Id"
// @Success 200 {object} models.TodoList
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists/{list_id} [get]
func (h *Handler) getWalletById(c *gin.Context) {

	walletId, err := uuid.FromString(c.Param("list_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid list_id param")
		return
	}

	wallet, err := h.services.Wallet.GetById(walletId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, wallet)
}

// @Summary Update todo List
// @Security ApiKeyAuth
// @Tags lists
// @Description update todo list
// @ID update-list
// @Accept  json
// @Produce  json
// @Param list_id path int true "List Id"
// @Param input body models.TodoListCreateUpdate true "list info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists/{list_id} [put]
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

// @Summary Delete List
// @Security ApiKeyAuth
// @Tags lists
// @Description delete list
// @ID delete-list
// @Accept  json
// @Produce  json
// @Param list_id path int true "List Id"
// @Success 200 {string} string ok
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists/{list_id} [delete]
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
