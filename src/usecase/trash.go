package usecase

import (
	"context"
	"errors"
	"gin/sdk/token"
	"gin/src/entity"
	"gin/src/enum"
	"gin/src/model"
	"gin/src/repository"
)

type TrashInterface interface {
	Exchange(ctx context.Context, newTrash model.ExchangeTrash, userID uint) (entity.Trash, error)
	GetHistory(ctx context.Context, userID uint) ([]entity.Trash, error)
}

type Trash struct {
	trashRepo repository.TrashInterface
	userRepo  repository.UserInterface
}

func InitTrash(trashRepo repository.TrashInterface, userRepo repository.UserInterface) TrashInterface {
	return &Trash{
		trashRepo: trashRepo,
		userRepo:  userRepo,
	}
}

func (uc *Trash) Exchange(ctx context.Context, newTrash model.ExchangeTrash, userID uint) (entity.Trash, error) {
	var trash entity.Trash

	exchangeTotal := newTrash.Mass

	if newTrash.Category == "Plastik" {
		trash.Category = enum.Plastik
		exchangeTotal *= 2000
	} else if newTrash.Category == "Kaca" {
		trash.Category = enum.Kaca
		exchangeTotal *= 5000
	} else if newTrash.Category == "Kertas" {
		trash.Category = enum.Kertas
		exchangeTotal *= 1000
	} else if newTrash.Category == "Eletronik" {
		trash.Category = enum.Elektronik
		exchangeTotal *= 25000
	} else if newTrash.Category == "Metal" {
		trash.Category = enum.Metal
		exchangeTotal *= 15000
	} else if newTrash.Category == "Kardus" {
		trash.Category = enum.Kardus
		exchangeTotal *= 7000
	} else if newTrash.Category == "Organik" {
		trash.Category = enum.Organik
		exchangeTotal *= 2000
	}

	trash = entity.Trash{
		Location:      "-",
		Mass:          newTrash.Mass,
		Code:          token.GenerateToken(),
		Status:        enum.Menunggu,
		UserID:        userID,
		ExchangeTotal: exchangeTotal,
	}

	trash, err := uc.trashRepo.Exchange(ctx, trash)

	if err != nil {
		return trash, errors.New("FAILED TO EXCHANGE TRASH")
	}

	return trash, nil
}

func (uc *Trash) GetHistory(ctx context.Context, userID uint) ([]entity.Trash, error) {
	trashes, err := uc.trashRepo.GetHistory(ctx, userID)

	if err != nil {
		return trashes, errors.New("FAILED TO GET EXCHANGE HISTORY")
	}

	return trashes, nil
}
