package repository

import (
	"context"
	"gin/database/mysql"
	"gin/src/entity"
	"mime/multipart"

	supabasestorageuploader "github.com/adityarizkyramadhan/supabase-storage-uploader"
)

type UserInterface interface {
	CreateUser(ctx context.Context, user entity.User) (entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
	UpdateUser(ctx context.Context, user entity.User) (entity.User, error)
	GetUserByID(ctx context.Context, userID uint) (entity.User, error)
	UploadPhotoProfile(ctx context.Context, PhotoUser *multipart.FileHeader) (string, error)
}

type User struct {
	sql      mysql.DB
	supabase supabasestorageuploader.SupabaseClientService
}

func InitUser(sql mysql.DB, supabase supabasestorageuploader.SupabaseClientService) UserInterface {
	return &User{
		sql:      sql,
		supabase: supabase,
	}
}

func (r *User) CreateUser(ctx context.Context, user entity.User) (entity.User, error) {
	if err := r.sql.Debug().WithContext(ctx).Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *User) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	var user entity.User
	if err := r.sql.Debug().WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *User) UpdateUser(ctx context.Context, user entity.User) (entity.User, error) {
	if err := r.sql.Debug().WithContext(ctx).Model(&user).Updates(user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *User) GetUserByID(ctx context.Context, userID uint) (entity.User, error) {
	var user entity.User
	if err := r.sql.Debug().WithContext(ctx).Where("ID = ?", userID).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *User) UploadPhotoProfile(ctx context.Context, PhotoUser *multipart.FileHeader) (string, error) {
	link, err := r.supabase.Upload(PhotoUser)
	if err != nil {
		return "", err
	}

	return link, nil
}
