package usecase

import (
	"context"
	"gin/src/entity"
	"gin/src/model"
	"gin/src/repository"
	"mime/multipart"
)

type ArticleInterface interface {
	CreateArticle(ctx context.Context, thumbnail *multipart.FileHeader, newArticle model.UploadArticle) (entity.Article, error)
	GetAllArticle(ctx context.Context) ([]entity.Article, error)
	GetArticleByID(ctx context.Context, articleID uint) (entity.Article, error)
}

type Article struct {
	articleRepo repository.ArticleInterface
}

func InitArticle(articleRepo repository.ArticleInterface) ArticleInterface {
	return &Article{
		articleRepo: articleRepo,
	}
}

func (uc *Article) CreateArticle(ctx context.Context, thumbnail *multipart.FileHeader, newArticle model.UploadArticle) (entity.Article, error) {
	var article entity.Article

	link, err := uc.articleRepo.UploadThumbnail(ctx, thumbnail)

	if err != nil {
		return article, err
	}

	article = entity.Article{
		Thumbnail: link,
		Title:     newArticle.Title,
		Body:      newArticle.Body,
	}

	article, err = uc.articleRepo.CreateArticle(ctx, article)

	if err != nil {
		return article, err
	}

	return article, nil
}

func (uc *Article) GetAllArticle(ctx context.Context) ([]entity.Article, error) {
	articles, err := uc.articleRepo.GetAllArticle(ctx)

	if err != nil {
		return articles, err
	}

	return articles, nil
}

func (uc *Article) GetArticleByID(ctx context.Context, articleID uint) (entity.Article, error) {
	article, err := uc.articleRepo.GetArticleByID(ctx, articleID)

	if err != nil {
		return article, err
	}

	return article, nil
}
