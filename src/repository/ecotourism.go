package repository

import (
	"context"
	"gin/database/mysql"
	"gin/src/entity"

	supabasestorageuploader "github.com/adityarizkyramadhan/supabase-storage-uploader"
)

type EcoTourismInterface interface {
	GetAll(ctx context.Context) ([]entity.Ecotourism, error)
	Create(ctx context.Context, newEcotourism entity.Ecotourism) (entity.Ecotourism, error)
	GetByID(ctx context.Context, ecoID uint) (entity.Ecotourism, error)
	Update(ctx context.Context, ecotourism entity.Ecotourism) (entity.Ecotourism, error)
}

type EcoTourism struct {
	sql      mysql.DB
	supabase supabasestorageuploader.SupabaseClientService
}

func InitEcoTourism(sql mysql.DB, supabase supabasestorageuploader.SupabaseClientService) EcoTourismInterface {
	return &EcoTourism{
		sql:      sql,
		supabase: supabase,
	}
}

func (r *EcoTourism) Create(ctx context.Context, newEcotourism entity.Ecotourism) (entity.Ecotourism, error) {
	if err := r.sql.Debug().WithContext(ctx).Create(&newEcotourism).Error; err != nil {
		return newEcotourism, err
	}

	return newEcotourism, nil
}

func (r *EcoTourism) GetAll(ctx context.Context) ([]entity.Ecotourism, error) {
	var allTourism []entity.Ecotourism

	if err := r.sql.Debug().WithContext(ctx).Find(&allTourism).Error; err != nil {
		return allTourism, err
	}

	return allTourism, nil
}

func (r *EcoTourism) GetByID(ctx context.Context, ecoID uint) (entity.Ecotourism, error) {
	var ecotourism entity.Ecotourism
	if err := r.sql.Debug().WithContext(ctx).Preload("Comment.User").Where("id = ?", ecoID).First(&ecotourism).Error; err != nil {
		return ecotourism, err
	}

	return ecotourism, nil
}

func (r *EcoTourism) Update(ctx context.Context, ecotourism entity.Ecotourism) (entity.Ecotourism, error) {
	if err := r.sql.Debug().WithContext(ctx).Preload("Comment").Model(&ecotourism).Updates(ecotourism).Error; err != nil {
		return ecotourism, err
	}
	return ecotourism, nil
}
