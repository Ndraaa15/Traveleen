package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"gin/src/entity"
	"gin/src/model"
	"gin/src/repository"
)

type CartInterface interface {
	AddCart(ctx context.Context, userID uint, ecoID uint, newProductCart model.CartProduct) (entity.Cart, error)
	GetCart(ctx context.Context, userID uint) (entity.Cart, error)
	DeleteCartContent(ctx context.Context, userID uint, ecoID uint) error
}

type Cart struct {
	cartRepo repository.CartInterface
	ecoRepo  repository.EcoTourismInterface
}

func InitCart(cartRepo repository.CartInterface, ecoRepo repository.EcoTourismInterface) CartInterface {
	return &Cart{
		cartRepo: cartRepo,
		ecoRepo:  ecoRepo,
	}
}

func (uc *Cart) AddCart(ctx context.Context, userID uint, ecoID uint, newCartProduct model.CartProduct) (entity.Cart, error) {
	var cartProduct entity.CartProduct

	if newCartProduct.Quantity == 0 {
		return entity.Cart{}, errors.New("QUANTITY CANNOT BE 0")
	}

	cart, err := uc.cartRepo.AddCart(ctx, userID)

	if err != nil {
		return cart, errors.New("FAILED TO CREATE CART")
	}

	ecotourism, err := uc.ecoRepo.GetByID(ctx, ecoID)

	if err != nil {
		return cart, errors.New("ECOTOURISM NOT FOUND")
	}

	var images []interface{}
	if err := json.Unmarshal([]byte(ecotourism.Thumbnail), &images); err != nil {
		return cart, errors.New("FAILED TO UNMARSHAL JSON")
	}

	cartProduct = entity.CartProduct{
		CartID:          cart.ID,
		EcoID:           ecoID,
		EcoImage:        images[0].(string),
		EcoName:         ecotourism.Name,
		EcoLocation:     ecotourism.Region,
		EcoCategory:     ecotourism.Category,
		Quantity:        newCartProduct.Quantity,
		Price:           ecotourism.Price * float64(newCartProduct.Quantity),
		PricePerProduct: ecotourism.Price,
	}

	notFound := true

	for i, product := range cart.CartProduct {
		if product.EcoID == ecoID && newCartProduct.Quantity != 0 {
			cart.TotalPrice = cart.TotalPrice - product.Price

			err := uc.cartRepo.DeleteCartContent(ctx, cart.ID, ecoID)

			if err != nil {
				return cart, errors.New("FAILED TO DELETE CART CONTENT")
			}

			cart.CartProduct[i] = cartProduct

			notFound = false

			break
		}
	}

	if notFound {
		cart.CartProduct = append(cart.CartProduct, cartProduct)
	}

	cart.TotalPrice = cart.TotalPrice + cartProduct.Price

	cart, err = uc.cartRepo.UpdateCart(ctx, cart, userID)

	if err != nil {
		return cart, errors.New("FAILED TO UPDATE CART")
	}

	return cart, nil
}

func (uc *Cart) GetCart(ctx context.Context, userID uint) (entity.Cart, error) {
	cart, err := uc.cartRepo.GetCart(ctx, userID)

	if err != nil {
		return cart, errors.New("FAILED TO GET CART")
	}

	return cart, nil
}

func (uc *Cart) DeleteCartContent(ctx context.Context, userID uint, ecoID uint) error {
	cart, err := uc.cartRepo.GetCart(ctx, userID)

	if err != nil {
		return errors.New("FAILED TO GET CART")
	}

	for _, product := range cart.CartProduct {
		if product.EcoID == ecoID && product.CartID == cart.ID {
			cart.TotalPrice = cart.TotalPrice - product.Price
			break
		}
	}

	cart, err = uc.cartRepo.UpdateCart(ctx, cart, userID)

	if err != nil {
		return errors.New("FAILED TO UPDATE CART")
	}

	err = uc.cartRepo.DeleteCartContent(ctx, cart.ID, ecoID)

	if err != nil {
		return errors.New("FAILED TO DELETE CART CONTENT")
	}

	return nil
}
