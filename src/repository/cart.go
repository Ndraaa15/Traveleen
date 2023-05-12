package repository

import (
	"context"
	"gin/database/mysql"
	"gin/src/entity"

	supabasestorageuploader "github.com/adityarizkyramadhan/supabase-storage-uploader"
)

type CartInterface interface {
	AddCart(ctx context.Context, userID uint) (entity.Cart, error)
	UpdateCart(ctx context.Context, cart entity.Cart, userID uint) (entity.Cart, error)
	DeleteCartContent(ctx context.Context, cartID uint, ecoID uint) error
	GetCart(ctx context.Context, userID uint) (entity.Cart, error)
}

type Cart struct {
	sql      mysql.DB
	supabase supabasestorageuploader.SupabaseClientService
}

func InitCart(sql mysql.DB, supabase supabasestorageuploader.SupabaseClientService) CartInterface {
	return &Cart{
		sql:      sql,
		supabase: supabase,
	}
}

func (r *Cart) AddCart(ctx context.Context, userID uint) (entity.Cart, error) {
	var cart entity.Cart
	if err := r.sql.Debug().WithContext(ctx).Preload("CartProduct").Where(entity.Cart{UserID: userID}).Attrs(entity.Cart{UserID: userID}).FirstOrCreate(&cart).Error; err != nil {
		return cart, err
	}
	return cart, nil
}

func (r *Cart) UpdateCart(ctx context.Context, cart entity.Cart, userID uint) (entity.Cart, error) {
	if err := r.sql.Debug().WithContext(ctx).Where("user_id = ?", userID).Save(&cart).Error; err != nil {
		return cart, err
	}
	return cart, nil
}

func (r *Cart) GetCart(ctx context.Context, userID uint) (entity.Cart, error) {
	var cart entity.Cart
	if err := r.sql.Debug().WithContext(ctx).Preload("CartProduct").Where("user_id = ?", userID).First(&cart).Error; err != nil {
		return cart, err
	}
	return cart, nil
}

func (r *Cart) DeleteCartContent(ctx context.Context, cartID uint, ecoID uint) error {
	var product entity.CartProduct
	if err := r.sql.Debug().WithContext(ctx).Where("eco_id = ? AND cart_id = ?", ecoID, cartID).Delete(&product).Error; err != nil {
		return err
	}
	return nil
}
