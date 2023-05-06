package usecase

import (
	"context"
	"gin/src/entity"
	"gin/src/model"
	"gin/src/repository"
)

type EcoTourismInterface interface {
	Create(ctx context.Context, newEcotourism model.PostEcotourisms) (entity.Ecotourism, error)
	GetAll(ctx context.Context) ([]entity.Ecotourism, error)
	GetByID(ctx context.Context, ecoID uint) (entity.Ecotourism, error)
	GetByCategory(ctx context.Context, category string) ([]entity.Ecotourism, error)
	GetByRegion(ctx context.Context, region string) ([]entity.Ecotourism, error)
	GetByPrice(ctx context.Context, startPrice float64, endPrice float64) ([]entity.Ecotourism, error)
}

type EcoTourism struct {
	ecotourismRepo repository.EcoTourismInterface
}

func InitEcoTourism(ecotourismRepo repository.EcoTourismInterface) EcoTourismInterface {
	return &EcoTourism{
		ecotourismRepo: ecotourismRepo,
	}
}

func (uc *EcoTourism) Create(ctx context.Context, newEcotourism model.PostEcotourisms) (entity.Ecotourism, error) {
	var ecotourism entity.Ecotourism

	ecotourism = entity.Ecotourism{
		Region:          newEcotourism.Region,
		Category:        newEcotourism.Category,
		Thumbnail:       newEcotourism.Thumbnail,
		Name:            newEcotourism.Name,
		Rating:          newEcotourism.Rating,
		TotalRatings:    newEcotourism.TotalRatings,
		Price:           newEcotourism.Price,
		Description:     newEcotourism.Description,
		OperationalTime: newEcotourism.OperationalTime,
		Maps:            newEcotourism.Maps,
	}

	ecotourism, err := uc.ecotourismRepo.Create(ctx, ecotourism)

	if err != nil {
		return ecotourism, err
	}

	return ecotourism, nil
}

func (uc *EcoTourism) GetAll(ctx context.Context) ([]entity.Ecotourism, error) {
	ecotourisms, err := uc.ecotourismRepo.GetAll(ctx)

	if err != nil {
		return ecotourisms, err
	}

	return ecotourisms, nil
}

func (uc *EcoTourism) GetByID(ctx context.Context, ecoID uint) (entity.Ecotourism, error) {
	ecotourism, err := uc.ecotourismRepo.GetByID(ctx, ecoID)
	if err != nil {
		return ecotourism, err
	}
	return ecotourism, nil
}

func (uc *EcoTourism) GetByCategory(ctx context.Context, category string) ([]entity.Ecotourism, error) {
	ecotourisms, err := uc.ecotourismRepo.GetByCategory(ctx, category)
	if err != nil {
		return ecotourisms, err
	}
	return ecotourisms, nil
}

func (uc *EcoTourism) GetByRegion(ctx context.Context, region string) ([]entity.Ecotourism, error) {
	ecotourisms, err := uc.ecotourismRepo.GetByRegion(ctx, region)
	if err != nil {
		return ecotourisms, err
	}
	return ecotourisms, nil
}

func (uc *EcoTourism) GetByPrice(ctx context.Context, startPrice float64, endPrice float64) ([]entity.Ecotourism, error) {
	ecotourisms, err := uc.ecotourismRepo.GetByPrice(ctx, startPrice, endPrice)
	if err != nil {
		return ecotourisms, err
	}
	return ecotourisms, nil
}
