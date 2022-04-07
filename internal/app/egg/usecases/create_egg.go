package usecases

import (
	"example.com/go/boiler/internal/app/egg/domain"
)

type (
	CreateEggCommand struct {
		Name string `json:"name"`
	}

	CreateEggResponse struct {
		Name string `json:"name"`
	}

	CreateEggUseCase interface {
		Execute(command CreateEggCommand) (*CreateEggResponse, error)
	}

	createEggUseCase struct {
		eggRepository domain.EggRepository
	}
)

func NewCreateEggUseCase(repository domain.EggRepository) CreateEggUseCase {
	return &createEggUseCase{eggRepository: repository}
}

func (c createEggUseCase) Execute(command CreateEggCommand) (*CreateEggResponse, error) {
	return &CreateEggResponse{}, nil
}
