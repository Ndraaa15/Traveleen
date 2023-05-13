package usecase

import (
	"context"
	"errors"
	"gin/src/entity"
	"gin/src/repository"
)

type EcoTourismInterface interface {
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
