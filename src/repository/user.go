package repository

import (
	"context"
	"gin/database/mysql"
	"gin/src/entity"
	"mime/multipart"

	supabasestorageuploader "github.com/adityarizkyramadhan/supabase-storage-uploader"
)

type UserInterface interface {
	Create(ctx context.Context, user entity.User) (entity.User, error)
	GetByEmail(ctx context.Context, email string) (entity.User, error)
	Update(ctx context.Context, user entity.User) (entity.User, error)
	GetByID(ctx context.Context, userID uint) (entity.User, error)
	UploadPhotoProfile(PhotoUser *multipart.FileHeader) (string, error)
	UploadPhotoComment(PhotoComment []*multipart.FileHeader) ([]string, error)
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

func (r *User) Create(ctx context.Context, user entity.User) (entity.User, error) {
	if err := r.sql.Debug().WithContext(ctx).Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *User) GetByEmail(ctx context.Context, email string) (entity.User, error) {
	var user entity.User
	if err := r.sql.Debug().WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *User) Update(ctx context.Context, user entity.User) (entity.User, error) {
	if err := r.sql.Debug().WithContext(ctx).Where("id = ?", user.ID).Save(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *User) GetByID(ctx context.Context, userID uint) (entity.User, error) {
	var user entity.User
	if err := r.sql.Debug().WithContext(ctx).Where("id = ?", userID).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *User) UploadPhotoProfile(PhotoUser *multipart.FileHeader) (string, error) {
	link, err := r.supabase.Upload(PhotoUser)
	if err != nil {
		return "", err
	}

	return link, nil
}

func (r *User) UploadPhotoComment(PhotoComment []*multipart.FileHeader) ([]string, error) {
	linkPhotos := []string{}
	for _, file := range PhotoComment {
		link, err := r.supabase.Upload(file)
		if err != nil {
			return linkPhotos, err
		}
		linkPhotos = append(linkPhotos, link)
	}

	return linkPhotos, nil
}

func (r *User) Comment(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	if err := r.sql.Debug().WithContext(ctx).Create(&comment).Error; err != nil {
		return comment, err
	}

	return comment, nil
}
