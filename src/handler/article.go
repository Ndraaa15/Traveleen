package handler

import (
	"gin/sdk/message"
	"gin/src/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateArticle(ctx *gin.Context) {
	thumbnail, err := ctx.FormFile("thumbnail")

	if err != nil {
		message.ErrorResponse(ctx, http.StatusBadRequest, "Failed get data", err.Error())
		return
	}

	newArticle := model.UploadArticle{}

	article, err := h.uc.Article.CreateArticle(ctx.Request.Context(), thumbnail, newArticle)

	if err != nil {
		message.ErrorResponse(ctx, http.StatusInternalServerError, "Failed upload article", err.Error())
		return
	}

	message.SuccessResponse(ctx, http.StatusOK, "Success upload article", article)
}

func (h *Handler) GetAllArticles(ctx *gin.Context) {

}

func (h *Handler) GetArticleByID(ctx *gin.Context) {
	idStr := ctx.Param("id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		message.ErrorResponse(ctx, http.StatusBadRequest, "Failed to parse string", err.Error())
		return
	}

	article, err := h.uc.Article.GetArticleByID(ctx.Request.Context(), uint(id))

	if err != nil {
		message.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to load article", err.Error())
		return
	}

	message.SuccessResponse(ctx, http.StatusOK, "Success upload article", article)
}

func (h *Handler) GetAllArticle(ctx *gin.Context) {
	articles, err := h.uc.Article.GetAllArticle(ctx.Request.Context())

	if err != nil {
		message.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to load article", err.Error())
		return
	}

	message.SuccessResponse(ctx, http.StatusOK, "Success upload article", articles)
}
