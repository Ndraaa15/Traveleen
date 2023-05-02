package usecase

import "gin/src/repository"

type Usecase struct {
	User       UserInterface
	EcoTourism EcoTourismInterface
	Trash      TrashInterface
	Article    ArticleInterface
}

func InitUsecase(repo *repository.Repository) *Usecase {
	return &Usecase{
		User:       InitUser(repo.User),
		EcoTourism: InitEcoTourism(repo.EcoTourism),
		Trash:      InitTrash(repo.Trash, repo.User),
		Article:    InitArticle(repo.Article),
	}
}
