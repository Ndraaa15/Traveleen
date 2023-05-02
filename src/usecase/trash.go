package usecase

import (
	"context"
	"gin/sdk/currency"
	"gin/src/entity"
	"gin/src/enum"
	"gin/src/model"
	"gin/src/repository"
	"strings"

	"github.com/google/uuid"
)

type TrashInterface interface {
	ExchangeTrash(ctx context.Context, newTrash model.NewExchangeTrash, userID uint) (entity.Trash, error)
	GetExchangeHistory(ctx context.Context, userID uint) ([]entity.Trash, error)
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

func (uc *Trash) ExchangeTrash(ctx context.Context, newTrash model.NewExchangeTrash, userID uint) (entity.Trash, error) {

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

	trash.Location = newTrash.Location
	trash.Mass = newTrash.Mass
	trash.Code = uniqueCode
	trash.Status = enum.Menunggu
	trash.ExchangeTotal = exchangeTotal
	trash.UserID = userID

	user, err1 := uc.userRepo.GetUserByID(ctx, userID)

	if err1 != nil {
		return trash, err1
	}

	walletNow := currency.ConvertRupiahIntoFloat(user.Wallet)
	totalWallet := walletNow + exchangeTotal
	user.Wallet = currency.FormatRupiah(totalWallet)
	userUpdated, err2 := uc.userRepo.UpdateUser(ctx, user)

	if err2 != nil {
		return trash, err2
	}
	trash.User = userUpdated
	trash, err3 := uc.trashRepo.CreateExchange(ctx, trash)

	if err3 != nil {
		return trash, err3
	}

	return trash, nil
}

func (uc *Trash) GetExchangeHistory(ctx context.Context, userID uint) ([]entity.Trash, error) {
	trashes, err := uc.trashRepo.GetHistory(ctx, userID)

	if err != nil {
		return trashes, err
	}

	return trashes, nil
}
