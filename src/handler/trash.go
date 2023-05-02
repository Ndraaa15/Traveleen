package handler

import (
	"gin/sdk/message"
	"gin/src/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ExchangeTrash(ctx *gin.Context) {
	user, exist := ctx.Get("user")

	if !exist {
		message.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to load JWT token, please try again!", nil)
		return
	}

	userID := user.(uint)

	newTrash := model.NewExchangeTrash{}

	if err := ctx.ShouldBindJSON(&newTrash); err != nil {
		message.ErrorResponse(ctx, http.StatusBadRequest, "Failed to bind JSON", err.Error())
		return
	}

	trash, err := h.uc.Trash.ExchangeTrash(ctx, newTrash, userID)

	if err != nil {
		message.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to exchange trash", err.Error())
		return
	}

	message.SuccessResponse(ctx, http.StatusOK, "Success exchange trash", trash)

}

func (h *Handler) ExchangeTrashHistory(ctx *gin.Context) {
	user, exist := ctx.Get("user")

	if !exist {
		message.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to load JWT token, please try again!", nil)
		return
	}

	userID := user.(uint)

	trashes, err := h.uc.Trash.GetExchangeHistory(ctx, userID)

	if err != nil {
		message.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to load history", err.Error())
		return
	}

	message.SuccessResponse(ctx, http.StatusOK, "History Found", trashes)
}
