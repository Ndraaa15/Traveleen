package handler

import (
	"gin/sdk/message"
	"gin/src/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AddCart(ctx *gin.Context) {
	user, exist := ctx.Get("user")

	if !exist {
		message.ErrorResponse(ctx, http.StatusUnauthorized, "Failed to get JWT token!", nil)
		return
	}

	userID := user.(uint)

	ecoIDStr := ctx.Param("id")

	ecoID, err := strconv.Atoi(ecoIDStr)

	if err != nil {
		message.ErrorResponse(ctx, http.StatusBadRequest, "Error while parse string into uint", err.Error())
		return
	}

	newCartProduct := model.CartProduct{}

	if err := ctx.ShouldBindJSON(&newCartProduct); err != nil {
		message.ErrorResponse(ctx, http.StatusUnprocessableEntity, "Failed to bind JSON", err.Error())
		return
	}

	cart, err := h.uc.Cart.AddCart(ctx.Request.Context(), userID, uint(ecoID), newCartProduct)

	if err != nil {
		message.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to add product into cart", err.Error())
		return
	}

	message.SuccessResponse(ctx, http.StatusOK, "Added to cart", &cart)
}

func (h *Handler) GetCart(ctx *gin.Context) {
	user, exist := ctx.Get("user")

	if !exist {
		message.ErrorResponse(ctx, http.StatusUnauthorized, "Failed to get JWT token!", nil)
		return
	}

	userID := user.(uint)

	cart, err := h.uc.Cart.GetCart(ctx.Request.Context(), userID)

	if err != nil {
		message.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to get cart!", err.Error())
		return
	}

	message.SuccessResponse(ctx, http.StatusOK, "Cart found!", cart)
}

func (h *Handler) DeleteCartContent(ctx *gin.Context) {
	user, exist := ctx.Get("user")

	if !exist {
		message.ErrorResponse(ctx, http.StatusUnauthorized, "Failed to get JWT token!", nil)
		return
	}

	userID := user.(uint)

	ecoIDStr := ctx.Param("id")

	ecoID, err := strconv.Atoi(ecoIDStr)

	if err != nil {
		message.ErrorResponse(ctx, http.StatusBadRequest, "Error while parse string into uint", err.Error())
		return
	}

	err = h.uc.Cart.DeleteCartContent(ctx.Request.Context(), userID, uint(ecoID))

	if err != nil {
		message.ErrorResponse(ctx, http.StatusInternalServerError, "Failed delete cart content!", err.Error())
		return
	}

	message.SuccessResponse(ctx, http.StatusOK, "Content deleted!", nil)
}
