package usecase

import (
	"context"
	"errors"
	"gin/sdk/time"
	"gin/src/entity"
	"gin/src/enum"
	"gin/src/midtrans"
	"gin/src/model"
	"gin/src/repository"
)

type PurchaseInterface interface {
	Payment(ctx context.Context, userID uint, paymentType model.PaymentType) (model.PurchaseResponse, error)
	PurchasesHistory(ctx context.Context, userID uint) ([]entity.Purchase, error)
}

type Purchase struct {
	purchaseRepo repository.PurchaseInterface
	userRepo     repository.UserInterface
	cartRepo     repository.CartInterface
}

func InitPurchase(purchaseRepo repository.PurchaseInterface, userRepo repository.UserInterface, cartRepo repository.CartInterface) PurchaseInterface {
	return &Purchase{
		purchaseRepo: purchaseRepo,
		userRepo:     userRepo,
		cartRepo:     cartRepo,
	}
}

func (uc *Purchase) Payment(ctx context.Context, userID uint, paymentType model.PaymentType) (model.PurchaseResponse, error) {
	var purchase entity.Purchase
	var purchaseResponse model.PurchaseResponse
	var clock = time.GenerateTime()
	URL := "-"

	user, err := uc.userRepo.GetByID(ctx, userID)

	if err != nil {
		return purchaseResponse, errors.New("USER NOT FOUND")
	}

	cart, err := uc.cartRepo.GetCart(ctx, userID)

	if err != nil {
		return purchaseResponse, errors.New("CART NOT FOUND")
	}

	if paymentType.IsOnlinePayment {
		midtrans.InitializeSnapClient()
		resp, err := midtrans.CreateTransaction(user, cart)
		URL = resp.RedirectURL

		if err != nil {
			return purchaseResponse, errors.New("FAILED TO CREATE ONLINE TRANSACTION")
		}

		cart.TotalPrice = 0

		_, err = uc.cartRepo.UpdateCart(ctx, cart, userID)

		if err != nil {
			return purchaseResponse, errors.New("FAILED TO UPDATE CART")
		}

		for _, product := range cart.CartProduct {
			purchase = entity.Purchase{
				Date:        time.GenerateDate(),
				Time:        clock,
				Place:       product.EcoName + ", " + product.EcoLocation,
				Quantity:    product.Quantity,
				TotalPrice:  product.Price,
				Code:        "-",
				Status:      enum.Menunggu,
				PayCategory: enum.Online,
				EcoName:     product.EcoName,
				EcoCategory: product.EcoCategory,
				EcoLocation: product.EcoLocation,
				EcoImage:    product.EcoImage,
				UserID:      userID,
			}

			if err := uc.cartRepo.DeleteCartContent(ctx, cart.ID, product.EcoID); err != nil {
				return purchaseResponse, errors.New("FAILED TO DELETE CART CONTENT")
			}

			_, err = uc.purchaseRepo.CreatePurchase(ctx, purchase)

			if err != nil {
				return purchaseResponse, errors.New("FAILED TO CREATE PURCHASE")
			}
		}

	} else {

		if user.Wallet < cart.TotalPrice {
			return purchaseResponse, errors.New("INSUFFICIENT BALANCE")
		}

		user.Wallet = user.Wallet - cart.TotalPrice

		_, err = uc.userRepo.Update(ctx, user)

		if err != nil {
			return purchaseResponse, errors.New("FAILED TO UPDATE USER")
		}

		cart.TotalPrice = 0

		_, err = uc.cartRepo.UpdateCart(ctx, cart, userID)

		if err != nil {
			return purchaseResponse, errors.New("FAILED TO UPDATE CART")
		}

		for _, product := range cart.CartProduct {
			purchase = entity.Purchase{
				Date:        time.GenerateDate(),
				Time:        clock,
				Place:       product.EcoName + ", " + product.EcoLocation,
				Quantity:    product.Quantity,
				TotalPrice:  product.Price,
				Code:        "-",
				Status:      enum.Menunggu,
				PayCategory: enum.Coin,
				EcoName:     product.EcoName,
				EcoCategory: product.EcoCategory,
				EcoLocation: product.EcoLocation,
				EcoImage:    product.EcoImage,
				UserID:      userID,
			}

			if err := uc.cartRepo.DeleteCartContent(ctx, cart.ID, product.EcoID); err != nil {
				return purchaseResponse, errors.New("FAILED TO DELETE CART CONTENT")
			}

			_, err = uc.purchaseRepo.CreatePurchase(ctx, purchase)

			if err != nil {
				return purchaseResponse, errors.New("FAILED TO CREATE PURCHASE")
			}
		}
	}

	if paymentType.IsOnlinePayment {
		purchaseResponse = model.PurchaseResponse{
			IsSuccess:   "Payment Success",
			PaymentType: "Online Payment",
			URL:         URL,
		}
	} else {
		purchaseResponse = model.PurchaseResponse{
			IsSuccess:   "Payment Success",
			PaymentType: "Coin Payment",
			URL:         URL,
		}
	}
	return purchaseResponse, nil
}

func (uc *Purchase) PurchasesHistory(ctx context.Context, userID uint) ([]entity.Purchase, error) {
	purchases, err := uc.purchaseRepo.PurchasesHistory(ctx, userID)

	if err != nil {
		return purchases, errors.New("FAILED TO GET PURCHASES HISTORY")
	}

	return purchases, nil
}
