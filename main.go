package main

import (
	"gin/database/mysql"
	"gin/database/supabase"
	"gin/src/entity"
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

	var totalEcotourism int64
	if err := db.Model(&entity.Ecotourism{}).Count(&totalEcotourism).Error; err != nil {
		log.Fatal(err.Error())
	}

	if totalEcotourism == 0 {
		usecase.GenerateEcotourismDummy(db)
	}

	var totalComment int64
	if err := db.Model(&entity.Comment{}).Count(&totalComment).Error; err != nil {
		log.Fatal(err.Error())
	}

	if totalComment == 0 {
		usecase.GenerateCommentDummy(db)
	}

	var totalUser int64
	if err := db.Model(&entity.User{}).Count(&totalUser).Error; err != nil {
		log.Fatal(err.Error())
	}

	if totalUser == 0 {
		usecase.GenerateUserDummy(db)
	}

	handler := handler.InitHandler(usecase)

	handler.Run()
}
