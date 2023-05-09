package handler

import (
	"gin/sdk/message"
	"gin/src/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) PostTourism(ctx *gin.Context) {
	newEcotourism := model.PostEcotourisms{}

	if err := ctx.ShouldBindJSON(&newEcotourism); err != nil {
		message.ErrorResponse(ctx, http.StatusUnprocessableEntity, "Failed to bind JSON", err.Error())
		return
	}

	ecotourism, err := h.uc.EcoTourism.Create(ctx.Request.Context(), newEcotourism)

	if err != nil {
		message.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to create tourism!", err.Error())
		return
	}

	message.SuccessResponse(ctx, http.StatusOK, "Success to create ecotourism!", &ecotourism)
}

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
