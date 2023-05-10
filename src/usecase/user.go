package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"gin/sdk/jwt"
	"gin/sdk/numeric"
	"gin/sdk/password"
	"gin/sdk/time"
	"gin/src/entity"
	"gin/src/enum"
	"gin/src/midtrans"
	"gin/src/model"
	"gin/src/repository"
	"mime/multipart"
	"strconv"
)

type UserInterface interface {
	Register(ctx context.Context, userInput model.UserRegister) (entity.User, error)
	Login(ctx context.Context, userInput model.UserLogin) (model.UserLoginResponse, error)
	Update(ctx context.Context, userInput model.UserUpdate, userID uint) (entity.User, error)
	Delete(ctx context.Context, userID uint) error
	UploadPhotoProfile(ctx context.Context, userID uint, photoProfile *multipart.FileHeader) (entity.User, error)
	Profile(ctx context.Context, userID uint) (entity.User, error)
	Comment(ctx context.Context, ecoID uint, userID uint, photoComment []*multipart.FileHeader, data []string) (entity.Ecotourism, error)
	AddCart(ctx context.Context, userID uint, ecoID uint, newProductCart model.CartProduct) (entity.Cart, error)
	GetCart(ctx context.Context, userID uint) (entity.Cart, error)
	DeleteCartContent(ctx context.Context, userID uint, ecoID uint) error
	Payment(ctx context.Context, userID uint, paymentType model.PaymentType) (model.PurchaseResponse, error)
	PurchasesHistory(ctx context.Context, userID uint) ([]entity.Purchase, error)
}

type User struct {
	userRepo repository.UserInterface
	ecoRepo  repository.EcoTourismInterface
}

func InitUser(userRepo repository.UserInterface, ecoRepo repository.EcoTourismInterface) UserInterface {
	return &User{
		userRepo: userRepo,
		ecoRepo:  ecoRepo,
	}
}

func (uc *User) Register(ctx context.Context, userInput model.UserRegister) (entity.User, error) {
	var user entity.User

	hashedPassword, err := password.GeneratePassword(userInput.Password)

	if err != nil {
		return user, errors.New("FAILED TO GENERATE PASSWORD")
	}

	user = entity.User{
		Username: userInput.Username,
		Email:    userInput.Email,
		Password: hashedPassword,
	}

	user, err = uc.userRepo.Create(ctx, user)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (uc *User) Login(ctx context.Context, userInput model.UserLogin) (model.UserLoginResponse, error) {
	var userResponse model.UserLoginResponse

	user, err := uc.userRepo.GetByEmail(ctx, userInput.Email)

	if err != nil {
		return userResponse, errors.New("USER NOT FOUND")
	}

	if err := password.ComparePassword(user.Password, userInput.Password); err != nil {
		return userResponse, errors.New("WRONG PASSWORD")
	}

	token, err := jwt.GenerateJWTToken(user)

	if err != nil {
		return userResponse, errors.New("FAILED TO GENERATE TOKEN")
	}

	userResponse.User = user
	userResponse.Token = token

	return userResponse, nil
}

func (uc *User) Update(ctx context.Context, userInput model.UserUpdate, userID uint) (entity.User, error) {
	var userGender enum.Gender
	var userPassword string

	user, err := uc.userRepo.GetByID(ctx, userID)

	if err != nil {
		return user, errors.New("USER NOT FOUND")
	}

	if userInput.Gender {
		userGender = enum.Pria
	} else {
		userGender = enum.Wanita
	}

	if userInput.Password != "" {
		updatedPass, err := password.GeneratePassword(userInput.Password)

		if err != nil {
			return user, errors.New("FAILED TO GENERATE PASSWORD")
		}

		userPassword = updatedPass
	} else {
		userPassword = user.Password
	}

	user = entity.User{
		ID:           userID,
		Email:        userInput.Email,
		Username:     userInput.Username,
		Contact:      userInput.Contact,
		Birthday:     userInput.Birthday,
		Gender:       userGender,
		Region:       userInput.Region,
		Password:     userPassword,
		PhotoProfile: user.PhotoProfile,
	}

	userUpdate, err := uc.userRepo.Update(ctx, user)

	if err != nil {
		return user, errors.New("FAILED UPDATE USER")
	}

	return userUpdate, nil
}

func (uc *User) UploadPhotoProfile(ctx context.Context, userID uint, photoProfile *multipart.FileHeader) (entity.User, error) {
	linkPhoto, err := uc.userRepo.UploadPhotoProfile(photoProfile)
	if err != nil {
		return entity.User{}, errors.New("FAILED TO UPLOAD PHOTO")
	}

	user, err := uc.userRepo.GetByID(ctx, userID)
	if err != nil {
		return entity.User{}, errors.New("USER NOT FOUND")
	}

	user.PhotoProfile = linkPhoto

	userUpdate, err := uc.userRepo.Update(ctx, user)

	if err != nil {
		return entity.User{}, errors.New("FAILED TO UPDATE USER")
	}

	return userUpdate, nil
}

func (uc *User) Profile(ctx context.Context, userID uint) (entity.User, error) {
	user, err := uc.userRepo.GetByID(ctx, userID)
	if err != nil {
		return user, errors.New("USER NOT FOUND")
	}

	return user, nil
}

func (uc *User) Delete(ctx context.Context, userID uint) error {
	user, err := uc.userRepo.GetByID(ctx, userID)
	if err != nil {
		return err
	}

	err = uc.userRepo.Delete(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (uc *User) Comment(ctx context.Context, ecoID uint, userID uint, photoComment []*multipart.FileHeader, data []string) (entity.Ecotourism, error) {
	var comment entity.Comment

	linkPhoto, err := uc.userRepo.UploadPhotoComment(photoComment)

	if err != nil {
		return entity.Ecotourism{}, errors.New("FAILED TO UPLOAD PHOTO")
	}

	jsonLinkPhoto, err := json.Marshal(linkPhoto)

	if err != nil {
		return entity.Ecotourism{}, errors.New("FAILED TO MARSHAL JSON")
	}

	rating, err := strconv.ParseFloat(data[0], 64)

	if err != nil {
		return entity.Ecotourism{}, errors.New("FAILED TO PARSE FLOAT")
	}

	comment = entity.Comment{
		Date:         time.GenerateDate(),
		UserID:       userID,
		EcotourismID: ecoID,
		Rating:       rating,
		Body:         data[1],
		Thumbnail:    jsonLinkPhoto,
	}

	comment, err = uc.userRepo.Comment(ctx, comment)

	if err != nil {
		return entity.Ecotourism{}, errors.New("FAILED TO CREATE COMMENT")
	}

	ecotourism, err := uc.ecoRepo.GetByID(ctx, ecoID)

	if err != nil {
		return entity.Ecotourism{}, errors.New("ECOTOURISM NOT FOUND")
	}

	totalReview := ecotourism.TotalRatings + 1
	tmp := ecotourism.Rating*float64(ecotourism.TotalRatings) + rating
	totalRating := tmp / float64(totalReview)

	ecotourism = entity.Ecotourism{
		TotalRatings: totalReview,
		Rating:       numeric.RoundingRating(totalRating),
		Comment:      append(ecotourism.Comment, comment),
	}

	ecotourism, err = uc.ecoRepo.Update(ctx, ecotourism)

	if err != nil {
		return entity.Ecotourism{}, errors.New("FAILED TO UPDATE ECOTOURISM")
	}

	return ecotourism, nil
}

func (uc *User) AddCart(ctx context.Context, userID uint, ecoID uint, newCartProduct model.CartProduct) (entity.Cart, error) {
	var cartProduct entity.CartProduct

	cart, err := uc.userRepo.AddCart(ctx, userID)

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
		CartID:      cart.ID,
		EcoID:       ecoID,
		EcoImage:    images[0].(string),
		EcoName:     ecotourism.Name,
		EcoLocation: ecotourism.Region,
		EcoCategory: ecotourism.Category,
		Quantity:    newCartProduct.Quantity,
		Price:       (ecotourism.Price * float64(newCartProduct.Quantity)),
	}

	notFound := true

	for i, product := range cart.CartProduct {
		if product.EcoID == ecoID && newCartProduct.Quantity != 0 {
			cart.TotalPrice = cart.TotalPrice - product.Price

			err := uc.userRepo.DeleteCartContent(ctx, cart.ID, ecoID)

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

	cart, err = uc.userRepo.UpdateCart(ctx, cart, userID)

	if err != nil {
		return cart, errors.New("FAILED TO UPDATE CART")
	}

	return cart, nil
}

func (uc *User) GetCart(ctx context.Context, userID uint) (entity.Cart, error) {
	cart, err := uc.userRepo.GetCart(ctx, userID)

	if err != nil {
		return cart, errors.New("FAILED TO GET CART")
	}

	return cart, nil
}

func (uc *User) DeleteCartContent(ctx context.Context, userID uint, ecoID uint) error {
	cart, err := uc.userRepo.GetCart(ctx, userID)

	if err != nil {
		return errors.New("FAILED TO GET CART")
	}

	for _, product := range cart.CartProduct {
		if product.EcoID == ecoID && product.CartID == cart.ID {
			cart.TotalPrice = cart.TotalPrice - product.Price
			break
		}
	}

	cart, err = uc.userRepo.UpdateCart(ctx, cart, userID)

	if err != nil {
		return errors.New("FAILED TO UPDATE CART")
	}

	err = uc.userRepo.DeleteCartContent(ctx, cart.ID, ecoID)

	if err != nil {
		return errors.New("FAILED TO DELETE CART CONTENT")
	}

	return nil
}

func (uc *User) Payment(ctx context.Context, userID uint, paymentType model.PaymentType) (model.PurchaseResponse, error) {
	var purchase entity.Purchase
	var purchaseResponse model.PurchaseResponse
	URL := "-"

	user, err := uc.userRepo.GetByID(ctx, userID)

	if err != nil {
		return purchaseResponse, errors.New("USER NOT FOUND")
	}

	cart, err := uc.userRepo.GetCart(ctx, userID)

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

		_, err = uc.userRepo.UpdateCart(ctx, cart, userID)

		if err != nil {
			return purchaseResponse, errors.New("FAILED TO UPDATE CART")
		}

		for _, product := range cart.CartProduct {
			purchase = entity.Purchase{
				Date:        time.GenerateDate(),
				Place:       product.EcoName + ", " + product.EcoLocation,
				Quantity:    product.Quantity,
				TotalPrice:  product.Price,
				Code:        "-",
				Status:      enum.Menunggu,
				PayCategory: enum.Online,
				UserID:      userID,
			}

			if err := uc.userRepo.DeleteCartContent(ctx, cart.ID, product.EcoID); err != nil {
				return purchaseResponse, errors.New("FAILED TO DELETE CART CONTENT")
			}

			_, err = uc.userRepo.CreatePurchase(ctx, purchase)

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

		_, err = uc.userRepo.UpdateCart(ctx, cart, userID)

		if err != nil {
			return purchaseResponse, errors.New("FAILED TO UPDATE CART")
		}

		for _, product := range cart.CartProduct {
			purchase = entity.Purchase{
				Date:        time.GenerateDate(),
				Place:       product.EcoName + ", " + product.EcoLocation,
				Quantity:    product.Quantity,
				TotalPrice:  product.Price,
				Code:        "-",
				Status:      enum.Menunggu,
				PayCategory: enum.Coin,
				UserID:      userID,
			}

			if err := uc.userRepo.DeleteCartContent(ctx, cart.ID, product.EcoID); err != nil {
				return purchaseResponse, errors.New("FAILED TO DELETE CART CONTENT")
			}

			_, err = uc.userRepo.CreatePurchase(ctx, purchase)

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
			URL:         "-",
		}
	}
	return purchaseResponse, nil
}

func (uc *User) PurchasesHistory(ctx context.Context, userID uint) ([]entity.Purchase, error) {
	purchases, err := uc.userRepo.PurchasesHistory(ctx, userID)

	if err != nil {
		return purchases, errors.New("FAILED TO GET PURCHASES HISTORY")
	}

	return purchases, nil
}
