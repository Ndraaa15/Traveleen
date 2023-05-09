package handler

import (
	"gin/sdk/message"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateArticle(ctx *gin.Context) {
	user, exist := ctx.Get("user")

	if !exist {
		message.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to load JWT token, please try again!", nil)
		return
	}

	userID := user.(uint)

	thumbnail, err := ctx.FormFile("thumbnail")

	if err != nil {
		message.ErrorResponse(ctx, http.StatusBadRequest, "Failed get data", err.Error())
		return
	}

	title := ctx.PostForm("title")
	body := ctx.PostForm("body")
	dataArticle := []string{title, body}

	article, err := h.uc.Article.Create(ctx.Request.Context(), thumbnail, dataArticle, userID)

	if err != nil {
		message.ErrorResponse(ctx, http.StatusInternalServerError, "Failed upload article", err.Error())
		return
	}

	message.SuccessResponse(ctx, http.StatusOK, "Success upload article", &article)
}

func (h *Handler) GetAllArticles(ctx *gin.Context) {
	articles, err := h.uc.Article.GetAll(ctx.Request.Context())

	if err != nil {
		message.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to load articles!", err.Error())
		return
	}

	message.SuccessResponse(ctx, http.StatusOK, "Success upload article", &articles)
}

func (h *Handler) GetArticleByID(ctx *gin.Context) {
	idStr := ctx.Param("id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		message.ErrorResponse(ctx, http.StatusBadRequest, "Failed to parse string", err.Error())
		return
	}

	article, err := h.uc.Article.GetByID(ctx.Request.Context(), uint(id))

	if err != nil {
		message.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to load article", err.Error())
		return
	}

	message.SuccessResponse(ctx, http.StatusOK, "Success upload article", &article)
}
