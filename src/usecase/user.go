package usecase

import (
	"context"
	"errors"
	"gin/sdk/jwt"
	"gin/sdk/password"
	"gin/src/entity"
	"gin/src/enum"
	"gin/src/model"
	"gin/src/repository"
	"mime/multipart"
)

type UserInterface interface {
	Register(ctx context.Context, userInput model.UserRegister) (entity.User, error)
	Login(ctx context.Context, userInput model.UserLogin) (model.UserLoginResponse, error)
	Update(ctx context.Context, userInput model.UserUpdate, userID uint) (entity.User, error)
	UploadPhotoProfile(ctx context.Context, userID uint, photoProfile *multipart.FileHeader) (entity.User, error)
}

type User struct {
	userRepo repository.UserInterface
}

func InitUser(userRepo repository.UserInterface) UserInterface {
	return &User{
		userRepo: userRepo,
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

	user, err = uc.userRepo.CreateUser(ctx, user)

	if err != nil {
		return user, err
	}

	return user, nil

}

func (uc *User) Login(ctx context.Context, userInput model.UserLogin) (model.UserLoginResponse, error) {
	var userResponse model.UserLoginResponse

	user, err := uc.userRepo.GetUserByEmail(ctx, userInput.Email)

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
	user, err := uc.userRepo.GetUserByID(ctx, userID)

	if err != nil {
		return user, err
	}

	user.Email = userInput.Email
	user.Username = userInput.Username
	user.Contact = userInput.Contact
	user.Birthday = userInput.Birthday

	if userInput.Gender {
		user.Gender = enum.Pria
	} else {
		user.Gender = enum.Wanita
	}

	user.Region = userInput.Region

	updatedPass, err := password.GeneratePassword(userInput.Password)
	if err != nil {
		return user, err
	}
	user.Password = updatedPass

	userUpdated, err := uc.userRepo.UpdateUser(ctx, user)

	if err != nil {
		return user, errors.New("FAILED UPDATE USER")
	}

	return userUpdated, nil
}

func (uc *User) UploadPhotoProfile(ctx context.Context, userID uint, photoProfile *multipart.FileHeader) (entity.User, error) {
	link, err := uc.userRepo.UploadPhotoProfile(ctx, photoProfile)
	if err != nil {
		return entity.User{}, err
	}

	user, err := uc.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		return entity.User{}, err
	}

	user.PhotoProfile = link

	userUpdated, err := uc.userRepo.UpdateUser(ctx, user)
	if err != nil {
		return entity.User{}, err
	}

	return userUpdated, nil
}
