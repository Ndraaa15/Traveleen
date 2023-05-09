package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"gin/sdk/jwt"
	"gin/sdk/password"
	"gin/sdk/time"
	"gin/src/entity"
	"gin/src/enum"
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
		return user, err
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
		return userResponse, err
	}

	if err := password.ComparePassword(user.Password, userInput.Password); err != nil {
		return userResponse, err
	}

	token, err := jwt.GenerateJWTToken(user)

	if err != nil {
		return userResponse, err
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
		ID:       userID,
		Email:    userInput.Email,
		Username: userInput.Username,
		Contact:  userInput.Contact,
		Birthday: userInput.Birthday,
		Gender:   userGender,
		Region:   userInput.Region,
		Password: userPassword,
	}

	userUpdate, err := uc.userRepo.Update(ctx, user)

	if err != nil {
		return user, errors.New("FAILED UPDATE USER")
	}

	return userUpdate, nil
}

func (uc *User) UploadPhotoProfile(ctx context.Context, userID uint, photoProfile *multipart.FileHeader) (entity.User, error) {
	link, err := uc.userRepo.UploadPhotoProfile(photoProfile)
	if err != nil {
		return entity.User{}, err
	}

	user, err := uc.userRepo.GetByID(ctx, userID)
	if err != nil {
		return entity.User{}, err
	}

	user.PhotoProfile = link

	userUpdated, err := uc.userRepo.Update(ctx, user)
	if err != nil {
		return entity.User{}, err
	}

	return userUpdated, nil
}

func (uc *User) Profile(ctx context.Context, userID uint) (entity.User, error) {
	user, err := uc.userRepo.GetByID(ctx, userID)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (uc *User) Comment(ctx context.Context, ecoID uint, userID uint, photoComment []*multipart.FileHeader, data []string) (entity.Ecotourism, error) {
	var comment entity.Comment

	linkPhoto, err := uc.userRepo.UploadPhotoComment(photoComment)
	if err != nil {
		return entity.Ecotourism{}, err
	}
	jsonLinkPhoto, err := json.Marshal(linkPhoto)
	if err != nil {
		return entity.Ecotourism{}, err
	}

	rating, err := strconv.ParseFloat(data[0], 64)
	if err != nil {
		return entity.Ecotourism{}, err
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
		return entity.Ecotourism{}, err
	}

	ecotourism, err := uc.ecoRepo.GetByID(ctx, ecoID)

	if err != nil {
		return entity.Ecotourism{}, err
	}

	totalReview := ecotourism.TotalRatings + 1
	tmp := ecotourism.Rating*float64(ecotourism.TotalRatings) + rating
	totalRating := tmp / float64(totalReview)

	ecotourism.TotalRatings = totalReview
	ecotourism.Rating = totalRating
	ecotourism.Comment = append(ecotourism.Comment, comment)

	ecotourism, err = uc.ecoRepo.Save(ctx, ecotourism)
	if err != nil {
		return entity.Ecotourism{}, err
	}

	return ecotourism, nil
}

func (uc *User) AddCart(ctx context.Context, userID uint, ecoID uint, newCartProduct model.CartProduct) (entity.Cart, error) {
	var cartProduct entity.CartProduct

	cart, err := uc.userRepo.AddCart(ctx, userID)
	if err != nil {
		return cart, err
	}

	ecotourism, err := uc.ecoRepo.GetByID(ctx, ecoID)
	if err != nil {
		return cart, err
	}

	cartProduct = entity.CartProduct{
		CartID:      cart.ID,
		EcoID:       ecoID,
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
				return cart, err
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
		return cart, err
	}

	return cart, nil
}

func (uc *User) GetCart(ctx context.Context, userID uint) (entity.Cart, error) {
	cart, err := uc.userRepo.GetCart(ctx, userID)

	if err != nil {
		return cart, err
	}

	return cart, nil
}

func (uc *User) DeleteCartContent(ctx context.Context, userID uint, ecoID uint) error {
	cart, err := uc.userRepo.GetCart(ctx, userID)

	if err != nil {
		return err
	}

	for _, product := range cart.CartProduct {
		if product.EcoID == ecoID && product.CartID == cart.ID {
			cart.TotalPrice = cart.TotalPrice - product.Price
			break
		}
	}

	cart, err = uc.userRepo.UpdateCart(ctx, cart, userID)

	if err != nil {
		return err
	}

	err = uc.userRepo.DeleteCartContent(ctx, cart.ID, ecoID)

	if err != nil {
		return err
	}
	return nil
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
