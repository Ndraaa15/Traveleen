package handler

import (
	"gin/sdk/message"
	"gin/src/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Payment(ctx *gin.Context) {
	var paymentType model.PaymentType

	if err := ctx.ShouldBindJSON(&paymentType); err != nil {
		message.ErrorResponse(ctx, http.StatusUnprocessableEntity, "Failed to bind JSON", err.Error())
		return
	}

	user, exist := ctx.Get("user")

	if !exist {
		message.ErrorResponse(ctx, http.StatusUnauthorized, "Failed to get JWT token!", nil)
		return
	}

	userID := user.(uint)

	purchases, err := h.uc.Purchase.Payment(ctx.Request.Context(), userID, paymentType)

	if err != nil {
		message.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to pay", err.Error())
		return
	}

	message.SuccessResponse(ctx, http.StatusOK, "Success pay", purchases)
}

func (h *Handler) PurchasesHistory(ctx *gin.Context) {
	user, exist := ctx.Get("user")

	if !exist {
		message.ErrorResponse(ctx, http.StatusUnauthorized, "Failed to get JWT token!", nil)
		return
	}

	userID := user.(uint)

	purchases, err := h.uc.Purchase.PurchasesHistory(ctx.Request.Context(), userID)

	if err != nil {
		message.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to get purchase history!", err.Error())
		return
	}

	message.SuccessResponse(ctx, http.StatusOK, "Purchase history found!", purchases)

}
