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
	Comment(ctx context.Context, comment entity.Comment) (entity.Comment, error)
	AddCart(ctx context.Context, userID uint) (entity.Cart, error)
	UpdateCart(ctx context.Context, cart entity.Cart, userID uint) (entity.Cart, error)
	DeleteCartContent(ctx context.Context, cartID uint, ecoID uint) error
	GetCart(ctx context.Context, userID uint) (entity.Cart, error)
	Delete(ctx context.Context, user entity.User) error
	CreatePurchase(ctx context.Context, purchase entity.Purchase) (entity.Purchase, error)
	PurchasesHistory(ctx context.Context, userID uint) ([]entity.Purchase, error)
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

func (r *User) Delete(ctx context.Context, user entity.User) error {
	if err := r.sql.Debug().WithContext(ctx).Delete(&user).Error; err != nil {
		return err
	}
	return nil
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

func (r *User) AddCart(ctx context.Context, userID uint) (entity.Cart, error) {
	var cart entity.Cart
	if err := r.sql.Debug().WithContext(ctx).Preload("CartProduct").Where(entity.Cart{UserID: userID}).Attrs(entity.Cart{UserID: userID}).FirstOrCreate(&cart).Error; err != nil {
		return cart, err
	}
	return cart, nil
}

func (r *User) UpdateCart(ctx context.Context, cart entity.Cart, userID uint) (entity.Cart, error) {
	if err := r.sql.Debug().WithContext(ctx).Where("user_id = ?", userID).Save(&cart).Error; err != nil {
		return cart, err
	}
	return cart, nil
}

func (r *User) GetCart(ctx context.Context, userID uint) (entity.Cart, error) {
	var cart entity.Cart
	if err := r.sql.Debug().WithContext(ctx).Preload("CartProduct").Where("user_id = ?", userID).First(&cart).Error; err != nil {
		return cart, err
	}
	return cart, nil
}

func (r *User) DeleteCartContent(ctx context.Context, cartID uint, ecoID uint) error {
	var product entity.CartProduct
	if err := r.sql.Debug().WithContext(ctx).Where("eco_id = ? AND cart_id = ?", ecoID, cartID).Delete(&product).Error; err != nil {
		return err
	}
	return nil
}

func (r *User) CreatePurchase(ctx context.Context, purchase entity.Purchase) (entity.Purchase, error) {
	if err := r.sql.Debug().WithContext(ctx).Create(&purchase).Error; err != nil {
		return purchase, err
	}

	return purchase, nil
}

func (r *User) PurchasesHistory(ctx context.Context, userID uint) ([]entity.Purchase, error) {
	var purchases []entity.Purchase

	if err := r.sql.Debug().WithContext(ctx).Where("user_id = ?", userID).Find(&purchases).Error; err != nil {
		return purchases, err
	}

	return purchases, nil
}
