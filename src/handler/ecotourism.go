package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"gin/sdk/message"
	"gin/src/model"

	"github.com/gin-gonic/gin"
)

func (h *Handler) PostTourism(ctx *gin.Context) {
	newEcotourism := model.PostEcotourisms{}

	if err := ctx.ShouldBindJSON(&newEcotourism); err != nil {
		message.ErrorResponse(ctx, http.StatusBadRequest, "Failed to bind JSON", err.Error())
		return
	}

	ecotourism, err := h.uc.EcoTourism.Create(ctx.Request.Context(), newEcotourism)

	if err != nil {
		message.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to post tourism", err.Error())
		return
	}

	message.SuccessResponse(ctx, http.StatusOK, "Succesfull post tourism", ecotourism)
}

func (h *Handler) GetAllTourism(ctx *gin.Context) {
	ecotourisms, err := h.uc.EcoTourism.GetAll(ctx.Request.Context())
	if err != nil {
		message.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to get all tourisms", err.Error())
	}

	message.SuccessResponse(ctx, http.StatusOK, "Tourism found", ecotourisms)
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
		message.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to load tourism", err.Error())
		return
	}
	message.SuccessResponse(ctx, http.StatusOK, "Tourism found", ecotourism)
}

func (h *Handler) GetTourismByCategory(ctx *gin.Context) {
	category := ctx.Param("category")
	fmt.Println(category)
	ecotourisms, err := h.uc.EcoTourism.GetByCategory(ctx.Request.Context(), category)

	if err != nil {
		message.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to load tourism", err.Error())
		return
	}
	message.SuccessResponse(ctx, http.StatusOK, "Tourism found", ecotourisms)
}

func (h *Handler) GetTourismByPrice(ctx *gin.Context) {
	spStr := ctx.Param("startPrice")
	epStr := ctx.Param("endPrice")

	sp, err := strconv.ParseFloat(spStr, 64)
	if err != nil {
		message.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to parse float", err.Error())
		return
	}

	ep, err := strconv.ParseFloat(epStr, 64)
	if err != nil {
		message.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to parse float", err.Error())
		return
	}

	ecotourisms, err := h.uc.EcoTourism.GetByPrice(ctx.Request.Context(), sp, ep)

	if err != nil {
		message.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to load tourism", err.Error())
		return
	}
	message.SuccessResponse(ctx, http.StatusOK, "Tourism found", ecotourisms)
}

func (h *Handler) GetTourismByRegion(ctx *gin.Context) {
	region := ctx.Param("region")

	ecotourisms, err := h.uc.EcoTourism.GetByRegion(ctx.Request.Context(), region)

	if err != nil {
		message.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to load tourism", err.Error())
		return
	}
	message.SuccessResponse(ctx, http.StatusOK, "Tourism found", ecotourisms)
}
