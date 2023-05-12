package usecase

import "gin/src/repository"

type Usecase struct {
	User       UserInterface
	EcoTourism EcoTourismInterface
	Trash      TrashInterface
	Article    ArticleInterface
	Cart       CartInterface
	Purchase   PurchaseInterface
}

func InitUsecase(repo *repository.Repository) *Usecase {
	return &Usecase{
		User:       InitUser(repo.User, repo.EcoTourism),
		EcoTourism: InitEcoTourism(repo.EcoTourism),
		Trash:      InitTrash(repo.Trash, repo.User),
		Article:    InitArticle(repo.Article),
		Cart:       InitCart(repo.Cart, repo.EcoTourism),
		Purchase:   InitPurchase(repo.Purchase, repo.User, repo.Cart),
	}
}
