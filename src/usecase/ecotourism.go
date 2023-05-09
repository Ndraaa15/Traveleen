package usecase

import (
	"context"
	"errors"
	"gin/src/entity"
	"gin/src/model"
	"gin/src/repository"
)

type EcoTourismInterface interface {
	Create(ctx context.Context, newEcotourism model.PostEcotourisms) (entity.Ecotourism, error)
	GetAll(ctx context.Context) ([]entity.Ecotourism, error)
	GetByID(ctx context.Context, ecoID uint) (entity.Ecotourism, error)
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
		return ecotourism, errors.New("FAILED TO CREATE ECOTOURISM")
	}

	return ecotourism, nil
}

func (uc *EcoTourism) GetAll(ctx context.Context) ([]entity.Ecotourism, error) {
	ecotourisms, err := uc.ecotourismRepo.GetAll(ctx)

	if err != nil {
		return ecotourisms, errors.New("FAILED TO GET ALL ECOTOURISM")
	}

	return ecotourisms, nil
}

func (uc *EcoTourism) GetByID(ctx context.Context, ecoID uint) (entity.Ecotourism, error) {
	ecotourism, err := uc.ecotourismRepo.GetByID(ctx, ecoID)

	if err != nil {
		return ecotourism, errors.New("FAILED TO GET ECOTOURISM")
	}

	return ecotourism, nil
}
