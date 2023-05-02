package repository

import (
	"context"
	"gin/database/mysql"
	"gin/src/entity"
	"mime/multipart"

	supabasestorageuploader "github.com/adityarizkyramadhan/supabase-storage-uploader"
)

type ArticleInterface interface {
	UploadThumbnail(ctx context.Context, thumbnail *multipart.FileHeader) (string, error)
	CreateArticle(ctx context.Context, article entity.Article) (entity.Article, error)
	GetAllArticle(ctx context.Context) ([]entity.Article, error)
	GetArticleByID(ctx context.Context, articleID uint) (entity.Article, error)
}

type Article struct {
	sql      mysql.DB
	supabase supabasestorageuploader.SupabaseClientService
}

func InitArticle(sql mysql.DB, supabase supabasestorageuploader.SupabaseClientService) ArticleInterface {
	return &Article{
		sql:      sql,
		supabase: supabase,
	}
}

func (r *Article) UploadThumbnail(ctx context.Context, thumbnail *multipart.FileHeader) (string, error) {
	link, err := r.supabase.Upload(thumbnail)
	if err != nil {
		return "", err
	}

	return link, nil
}

func (r *Article) CreateArticle(ctx context.Context, article entity.Article) (entity.Article, error) {
	if err := r.sql.Debug().WithContext(ctx).Create(&article).Error; err != nil {
		return article, err
	}

	return article, nil
}

func (r *Article) GetAllArticle(ctx context.Context) ([]entity.Article, error) {
	var articles []entity.Article
	if err := r.sql.Debug().WithContext(ctx).Find(&articles).Error; err != nil {
		return articles, err
	}

	return articles, nil
}

func (r *Article) GetArticleByID(ctx context.Context, articleID uint) (entity.Article, error) {
	var article entity.Article
	if err := r.sql.Debug().WithContext(ctx).Where("id = ?", articleID).First(&article).Error; err != nil {
		return article, err
	}

	return article, nil
}
