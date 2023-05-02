package main

import (
	"gin/database/mysql"
	"gin/database/supabase"
	"gin/src/handler"
	"gin/src/repository"
	"gin/src/usecase"
	"log"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func main() {

	db, err := mysql.SqlInit()

	if err != nil {
		log.Fatal("Failed to initialize MySQL connection")
	}

	supabase := supabase.SupabaseInit()

	db.RunMigration()

	repository := repository.InitRepository(*db, supabase)
	usecase := usecase.InitUsecase(repository)
	handler := handler.InitHandler(usecase)

	handler.Run()
}
