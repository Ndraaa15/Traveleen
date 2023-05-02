package usecase

import (
	"context"
	"gin/src/entity"
	"gin/src/model"
	"gin/src/repository"
)

type EcoTourismInterface interface {
	PostEcotourism(ctx context.Context, newEcotourism model.PostEcotourisms) (entity.Ecotourism, error)
	GetAllTourisms(ctx context.Context) ([]entity.Ecotourism, error)
	GetTourismByID(ctx context.Context, ecoID uint) (entity.Ecotourism, error)
	GetTourismByCategory(ctx context.Context, category string) ([]entity.Ecotourism, error)
	GetTourismByRegion(ctx context.Context, region string) ([]entity.Ecotourism, error)
	GetTourismByPrice(ctx context.Context, startPrice float64, endPrice float64) ([]entity.Ecotourism, error)
}

type EcoTourism struct {
	ecotourismRepo repository.EcoTourismInterface
}

func InitEcoTourism(ecotourismRepo repository.EcoTourismInterface) EcoTourismInterface {
	return &EcoTourism{
		ecotourismRepo: ecotourismRepo,
	}
}

func (uc *EcoTourism) PostEcotourism(ctx context.Context, newEcotourism model.PostEcotourisms) (entity.Ecotourism, error) {
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

	ecotourism, err := uc.ecotourismRepo.PostEcotourism(ctx, ecotourism)

	if err != nil {
		return ecotourism, err
	}

	return ecotourism, nil
}

func (uc *EcoTourism) GetAllTourisms(ctx context.Context) ([]entity.Ecotourism, error) {
	ecotourisms, err := uc.ecotourismRepo.GetAllTourisms(ctx)

	if err != nil {
		return ecotourisms, err
	}

	return ecotourisms, nil
}

func (uc *EcoTourism) GetTourismByID(ctx context.Context, ecoID uint) (entity.Ecotourism, error) {
	ecotourism, err := uc.ecotourismRepo.GetTourismByID(ctx, ecoID)
	if err != nil {
		return ecotourism, err
	}
	return ecotourism, nil
}

func (uc *EcoTourism) GetTourismByCategory(ctx context.Context, category string) ([]entity.Ecotourism, error) {
	ecotourisms, err := uc.ecotourismRepo.GetTourismByCategory(ctx, category)
	if err != nil {
		return ecotourisms, err
	}
	return ecotourisms, nil
}

func (uc *EcoTourism) GetTourismByRegion(ctx context.Context, region string) ([]entity.Ecotourism, error) {
	ecotourisms, err := uc.ecotourismRepo.GetTourismByRegion(ctx, region)
	if err != nil {
		return ecotourisms, err
	}
	return ecotourisms, nil
}

func (uc *EcoTourism) GetTourismByPrice(ctx context.Context, startPrice float64, endPrice float64) ([]entity.Ecotourism, error) {
	ecotourisms, err := uc.ecotourismRepo.GetTourismByPrice(ctx, startPrice, endPrice)
	if err != nil {
		return ecotourisms, err
	}
	return ecotourisms, nil
}
