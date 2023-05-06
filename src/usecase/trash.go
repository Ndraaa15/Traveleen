package usecase

import (
	"context"
	"gin/src/entity"
	"gin/src/enum"
	"gin/src/model"
	"gin/src/repository"
	"strings"

	"github.com/google/uuid"
)

type TrashInterface interface {
	Exchange(ctx context.Context, newTrash model.NewExchangeTrash, userID uint) (entity.Trash, error)
	GetHistory(ctx context.Context, userID uint) ([]entity.Trash, error)
	ValidateCode(ctx context.Context, inputCode model.ValidateCode) (entity.Trash, error)
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

func (uc *Trash) Exchange(ctx context.Context, newTrash model.NewExchangeTrash, userID uint) (entity.Trash, error) {

	var trash entity.Trash

	id := uuid.New()
	uniqueCode := strings.ToUpper(id.String()[:8])

	exchangeTotal := newTrash.Mass

	if newTrash.Category == "Plastik" {
		trash.Category = enum.Plastik
		exchangeTotal *= 20000
	} else if newTrash.Category == "Kaca" {
		trash.Category = enum.Kaca
		exchangeTotal *= 50000
	} else if newTrash.Category == "Kertas" {
		trash.Category = enum.Kertas
		exchangeTotal *= 10000
	} else if newTrash.Category == "Eletronik" {
		trash.Category = enum.Elektronik
		exchangeTotal *= 100000
	} else if newTrash.Category == "Metal" {
		trash.Category = enum.Metal
		exchangeTotal *= 75000
	} else if newTrash.Category == "Kardus" {
		trash.Category = enum.Kardus
		exchangeTotal *= 20000
	} else if newTrash.Category == "Organik" {
		trash.Category = enum.Organik
		exchangeTotal *= 5000
	}

	user, err1 := uc.userRepo.GetByID(ctx, userID)

	trash.Location = "-"
	trash.Mass = newTrash.Mass
	trash.Code = uniqueCode
	trash.Status = enum.Menunggu
	trash.ExchangeTotal = exchangeTotal
	trash.UserID = userID

	if err1 != nil {
		return trash, err1
	}

	totalWallet := user.Wallet + exchangeTotal
	user.Wallet = totalWallet
	userUpdated, err2 := uc.userRepo.Update(ctx, user)
	trash.User = userUpdated

	if err2 != nil {
		return trash, err2
	}
	trash, err3 := uc.trashRepo.Exchange(ctx, trash)

	if err3 != nil {
		return trash, err3
	}

	return trash, nil
}

func (uc *Trash) GetHistory(ctx context.Context, userID uint) ([]entity.Trash, error) {
	trashes, err := uc.trashRepo.GetHistory(ctx, userID)

	if err != nil {
		return trashes, err
	}

	return trashes, nil
}

func (uc *Trash) ValidateCode(ctx context.Context, inputCode model.ValidateCode) (entity.Trash, error) {
	trash, err := uc.trashRepo.GetByCode(ctx, inputCode.Code)

	if err != nil {
		return trash, err
	}

	if inputCode.IsSuccess {
		trash.Status = enum.Berhasil
	} else {
		trash.Status = enum.Gagal
	}

	trash, err = uc.trashRepo.Update(ctx, trash)

	if err != nil {
		return trash, err
	}

	return trash, nil
}
