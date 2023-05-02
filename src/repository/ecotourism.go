package repository

import (
	"context"
	"gin/database/mysql"
	"gin/src/entity"

	supabasestorageuploader "github.com/adityarizkyramadhan/supabase-storage-uploader"
)

type EcoTourismInterface interface {
	GetAllTourisms(ctx context.Context) ([]entity.Ecotourism, error)
	PostEcotourism(ctx context.Context, newEcotourism entity.Ecotourism) (entity.Ecotourism, error)
	GetTourismByID(ctx context.Context, ecoID uint) (entity.Ecotourism, error)
	GetTourismByCategory(ctx context.Context, category string) ([]entity.Ecotourism, error)
	GetTourismByRegion(ctx context.Context, region string) ([]entity.Ecotourism, error)
	GetTourismByPrice(ctx context.Context, startPrice float64, endPrice float64) ([]entity.Ecotourism, error)
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

func (r *EcoTourism) GetAllTourisms(ctx context.Context) ([]entity.Ecotourism, error) {
	var allTourism []entity.Ecotourism

	if err := r.sql.Debug().WithContext(ctx).Preload("OperationalTime").Find(&allTourism).Error; err != nil {
		return allTourism, err
	}

	return allTourism, nil
}

func (r *EcoTourism) PostEcotourism(ctx context.Context, newEcotourism entity.Ecotourism) (entity.Ecotourism, error) {
	if err := r.sql.Debug().WithContext(ctx).Create(&newEcotourism).Error; err != nil {
		return newEcotourism, err
	}

	return newEcotourism, nil
}

func (r *EcoTourism) GetTourismByID(ctx context.Context, ecoID uint) (entity.Ecotourism, error) {
	var ecotourism entity.Ecotourism
	if err := r.sql.Debug().WithContext(ctx).Where("id = ?", ecoID).First(&ecotourism).Error; err != nil {
		return ecotourism, err
	}

	return ecotourism, nil
}

func (r *EcoTourism) GetTourismByCategory(ctx context.Context, category string) ([]entity.Ecotourism, error) {
	var ecotourisms []entity.Ecotourism
	if err := r.sql.Debug().WithContext(ctx).Where("category = ?", category).Find(&ecotourisms).Error; err != nil {
		return ecotourisms, err
	}

	return ecotourisms, nil
}

func (r *EcoTourism) GetTourismByRegion(ctx context.Context, region string) ([]entity.Ecotourism, error) {
	var ecotourisms []entity.Ecotourism
	if err := r.sql.Debug().WithContext(ctx).Where("region = ?", region).Find(&ecotourisms).Error; err != nil {
		return ecotourisms, err
	}

	return ecotourisms, nil
}

func (r *EcoTourism) GetTourismByPrice(ctx context.Context, startPrice float64, endPrice float64) ([]entity.Ecotourism, error) {
	var ecotourisms []entity.Ecotourism
	if err := r.sql.Debug().WithContext(ctx).Where("price BETWEEN ? AND ?", startPrice, endPrice).Find(&ecotourisms).Error; err != nil {
		return ecotourisms, err
	}

	return ecotourisms, nil
}
