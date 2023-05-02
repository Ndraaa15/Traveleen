package repository

import (
	"gin/database/mysql"

	supabasestorageuploader "github.com/adityarizkyramadhan/supabase-storage-uploader"
)

type Repository struct {
	User       UserInterface
	EcoTourism EcoTourismInterface
	Trash      TrashInterface
	Article    ArticleInterface
}

func InitRepository(sql mysql.DB, supabase supabasestorageuploader.SupabaseClientService) *Repository {
	return &Repository{
		User:       InitUser(sql, supabase),
		EcoTourism: InitEcoTourism(sql, supabase),
		Trash:      InitTrash(sql, supabase),
		Article:    InitArticle(sql, supabase),
	}
}
