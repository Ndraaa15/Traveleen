package repository

import (
	"context"
	"gin/database/mysql"
	"gin/src/entity"

	supabasestorageuploader "github.com/adityarizkyramadhan/supabase-storage-uploader"
)

type PurchaseInterface interface {
	CreatePurchase(ctx context.Context, purchase entity.Purchase) (entity.Purchase, error)
	PurchasesHistory(ctx context.Context, userID uint) ([]entity.Purchase, error)
}

type Purchase struct {
	sql      mysql.DB
	supabase supabasestorageuploader.SupabaseClientService
}

func InitPurchase(sql mysql.DB, supabase supabasestorageuploader.SupabaseClientService) PurchaseInterface {
	return &Purchase{
		sql:      sql,
		supabase: supabase,
	}
}

func (r *Purchase) CreatePurchase(ctx context.Context, purchase entity.Purchase) (entity.Purchase, error) {
	if err := r.sql.Debug().WithContext(ctx).Create(&purchase).Error; err != nil {
		return purchase, err
	}

	return purchase, nil
}

func (r *Purchase) PurchasesHistory(ctx context.Context, userID uint) ([]entity.Purchase, error) {
	var purchases []entity.Purchase

	if err := r.sql.Debug().WithContext(ctx).Where("user_id = ?", userID).Find(&purchases).Error; err != nil {
		return purchases, err
	}

	return purchases, nil
}
