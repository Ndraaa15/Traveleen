package usecase

import (
	"context"
	"errors"
	"gin/sdk/time"
	"gin/src/entity"
	"gin/src/repository"
	"mime/multipart"
)

type ArticleInterface interface {
	Create(ctx context.Context, thumbnail *multipart.FileHeader, dataArticle []string, userID uint) (entity.Article, error)
	GetAll(ctx context.Context) ([]entity.Article, error)
	GetByID(ctx context.Context, articleID uint) (entity.Article, error)
}

type Article struct {
	articleRepo repository.ArticleInterface
}

func InitArticle(articleRepo repository.ArticleInterface) ArticleInterface {
	return &Article{
		articleRepo: articleRepo,
	}
}

func (uc *Article) Create(ctx context.Context, thumbnail *multipart.FileHeader, dataArticle []string, userID uint) (entity.Article, error) {
	var article entity.Article

	link, err := uc.articleRepo.UploadThumbnail(ctx, thumbnail)

	if err != nil {
		return article, errors.New("FAILED TO UPLOAD THUMBNAIL")
	}

	article = entity.Article{
		Date:      time.GenerateDate(),
		Thumbnail: link,
		Title:     dataArticle[0],
		Body:      dataArticle[1],
		UserID:    userID,
	}

	article, err = uc.articleRepo.Create(ctx, article)

	if err != nil {
		return article, err
	}

	return article, nil
}

func (uc *Article) GetAll(ctx context.Context) ([]entity.Article, error) {
	articles, err := uc.articleRepo.GetAll(ctx)

	if err != nil {
		return articles, errors.New("FAILED TO GET ALL ARTICLES")
	}

	return articles, nil
}

func (uc *Article) GetByID(ctx context.Context, articleID uint) (entity.Article, error) {
	article, err := uc.articleRepo.GetByID(ctx, articleID)

	if err != nil {
		return article, errors.New("FAILED TO GET ARTICLE")
	}

	return article, nil
}
