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
