package handler

import (
	"gin/sdk/message"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllTourism(ctx *gin.Context) {
	ecotourisms, err := h.uc.EcoTourism.GetAll(ctx.Request.Context())

	if err != nil {
		message.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to get all ecotourisms!", err.Error())
	}

	message.SuccessResponse(ctx, http.StatusOK, "Ecotourisms found!", &ecotourisms)
}

func (h *Handler) GetTourismByID(ctx *gin.Context) {
	ecoIDStr := ctx.Param("id")

	ecoID, err := strconv.Atoi(ecoIDStr)

	if err != nil {
		message.ErrorResponse(ctx, http.StatusBadRequest, "Error while parse string into uint", err.Error())
		return
	}

	ecotourism, err := h.uc.EcoTourism.GetByID(ctx.Request.Context(), uint(ecoID))

	if err != nil {
		message.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to get ecotourism!", err.Error())
		return
	}
	message.SuccessResponse(ctx, http.StatusOK, "Ecotourism found!", &ecotourism)
}
