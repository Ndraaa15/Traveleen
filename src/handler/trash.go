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

	trash, err := h.uc.Trash.Exchange(ctx, newTrash, userID)

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

	trashes, err := h.uc.Trash.GetHistory(ctx, userID)

	if err != nil {
		message.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to load history", err.Error())
		return
	}

	message.SuccessResponse(ctx, http.StatusOK, "History Found", trashes)
}

func (h *Handler) ValidateCode(ctx *gin.Context) {
	code := model.ValidateCode{}

	if err := ctx.ShouldBindJSON(&code); err != nil {
		message.ErrorResponse(ctx, http.StatusBadRequest, "Failed to bind JSON", err.Error())
		return
	}

	trash, err := h.uc.Trash.ValidateCode(ctx, code)

	if err != nil {
		message.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to validate code", err.Error())
		return
	}

	message.SuccessResponse(ctx, http.StatusOK, "Validate success", trash)
}
