package handler

import (
	"gin/src/middleware"
	"gin/src/usecase"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

var once = sync.Once{}

type Handler struct {
	http *gin.Engine
	uc   *usecase.Usecase
}

func InitHandler(uc *usecase.Usecase) *Handler {
	r := Handler{}
	once.Do(func() {
		r.http = gin.Default()
		r.uc = uc
		r.RoutesAndMiddleware()

	})
	return &r
}

func (h *Handler) RoutesAndMiddleware() {
	h.http.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world",
		})
	})

	v1 := h.http.Group("/api/v1")

	h.http.Use(middleware.CORS())

	user := h.http.Group(v1.BasePath() + "/user")
	user.POST("/signup", h.UserRegister)
	user.POST("/login", h.UserLogin)
	user.Use(middleware.IsUserLoggedIn).
		GET("/profile", h.GetProfile).
		DELETE("/delete", h.DeleteAccount).
		PUT("/update", h.UserUpdate).
		POST("/upload/photo", h.UploadPhotoProfile).
		GET("/cart", h.GetCart).
		GET("/history", h.PurchasesHistory).
		POST("/payment", h.Payment)

	eco := h.http.Group(v1.BasePath() + "/tourism")
	eco.Use(middleware.IsUserLoggedIn).
		GET("/", h.GetAllTourism).
		GET("/:id", h.GetTourismByID).
		POST("/:id/comment", h.UploadComment).
		POST("/add/:id/cart", h.AddCart).
		DELETE("/del/:id/cart", h.DeleteCartContent)

	trash := h.http.Group(v1.BasePath() + "/trash")
	trash.Use(middleware.IsUserLoggedIn).
		POST("/exchange", h.ExchangeTrash).
		GET("/exchange/history", h.ExchangeTrashHistory)

	article := h.http.Group(v1.BasePath() + "/article")
	article.Use(middleware.IsUserLoggedIn).
		GET("/", h.GetAllArticles).
		GET(":id", h.GetArticleByID).
		POST("/create", h.CreateArticle)
}

func (h *Handler) Run() {
	h.http.Run()
}
