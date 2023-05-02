package main

import (
	"gin/database/mysql"
	"gin/database/supabase"
	"gin/src/handler"
	"gin/src/middleware"
	"gin/src/repository"
	"gin/src/usecase"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var Router *gin.Engine

func main() {
	r := gin.Default()
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load .env")
	}

	db, err := mysql.SqlInit()

	if err != nil {
		log.Fatal("Failed to initialize MySQL connection")
	}

	supabase := supabase.SupabaseInit()

	db.RunMigration()

	repo := repository.InitRepository(*db, supabase)
	uc := usecase.InitUsecase(repo)
	h := handler.InitHandler(uc)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})

	v1 := r.Group("/api/v1")

	user := r.Group(v1.BasePath() + "/user")
	user.POST("/signup", h.UserRegister) //new user signup
	user.POST("/login", h.UserLogin)     //user login
	user.Use(middleware.IsUserLoggedIn).
		PUT("/update", h.UserUpdate).               //user update profile without photo profile
		POST("/upload/photo", h.UploadPhotoProfile) //user upload photo profile

	eco := r.Group(v1.BasePath() + "/tourism")
	eco.POST("/post", h.PostTourism) //post ecotourism
	eco.Use(middleware.IsUserLoggedIn).
		GET("/", h.GetAllTourism).                                       //get all eco tourism
		GET("/:id", h.GetTourismByID).                                   //get eco tourism by id
		GET("/filter/category/:category", h.GetTourismByCategory).       //get filtered eco tourism by category
		GET("/filter/price/:startPrice/:endPrice", h.GetTourismByPrice). //get filtered eco tourism by price
		GET("/filter/region/:region", h.GetTourismByRegion).             //get filtered
		POST("/:id/comment", h.UploadComment)                            //comment on current eco tourism, id in here is id ecotourism

	trash := r.Group(v1.BasePath() + "/trash")
	trash.Use(middleware.IsUserLoggedIn).
		POST("/exchange", h.ExchangeTrash).              //exchange trash into coin
		GET("/exchange/history", h.ExchangeTrashHistory) //get exchange trash history

	article := r.Group(v1.BasePath() + "/article")
	article.Use(middleware.IsUserLoggedIn).
		GET("/", h.GetAllArticles).      //get all articles
		GET(":id", h.GetArticleByID).    //get article by id
		POST("/create", h.CreateArticle) //create a article

	r.Run()
}
