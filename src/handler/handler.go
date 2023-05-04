package handler

import (
	"gin/src/middleware"
	"gin/src/usecase"
	"net/http"
	"sync"

	"github.com/gin-contrib/cors"
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

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	h.http.Use(cors.New(config))

	v1 := h.http.Group("/api/v1")

	user := h.http.Group(v1.BasePath() + "/user")
	user.POST("/signup", h.UserRegister) //new user signup
	user.POST("/login", h.UserLogin)     //user login
	user.Use(middleware.IsUserLoggedIn).
		GET("/profile", h.GetProfile).
		PUT("/update", h.UserUpdate).               //user update profile without photo profile
		POST("/upload/photo", h.UploadPhotoProfile) //user upload photo profile

	eco := h.http.Group(v1.BasePath() + "/tourism")
	eco.POST("/post", h.PostTourism) //post ecotourism
	eco.Use(middleware.IsUserLoggedIn).
		GET("/", h.GetAllTourism).                                       //get all eco tourism
		GET("/:id", h.GetTourismByID).                                   //get eco tourism by id
		GET("/filter/category/:category", h.GetTourismByCategory).       //get filtered eco tourism by category
		GET("/filter/price/:startPrice/:endPrice", h.GetTourismByPrice). //get filtered eco tourism by price
		GET("/filter/region/:region", h.GetTourismByRegion).             //get filtered
		POST("/:id/comment", h.UploadComment)                            //comment on current eco tourism, id in here is id ecotourism

	trash := h.http.Group(v1.BasePath() + "/trash")
	trash.Use(middleware.IsUserLoggedIn).
		POST("/exchange", h.ExchangeTrash).              //exchange trash into coin
		GET("/exchange/history", h.ExchangeTrashHistory) //get exchange trash history

	article := h.http.Group(v1.BasePath() + "/article")
	article.Use(middleware.IsUserLoggedIn).
		GET("/", h.GetAllArticles).      //get all articles
		GET(":id", h.GetArticleByID).    //get article by id
		POST("/create", h.CreateArticle) //create a article

}

func (h *Handler) Run() {
	h.http.Run()
}
