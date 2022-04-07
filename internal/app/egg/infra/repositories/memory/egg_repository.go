package memory

import (
	"example.com/go/boiler/internal/app/egg/domain"
	"example.com/go/boiler/internal/app/infra/config"
)

type EggRepository struct {
	size int
	eggs []domain.Egg
}

func NewEggRepository(config *config.Configuration) *EggRepository {
	return &EggRepository{size: config.EggRepository.Memory.Size}
}

func (repository *EggRepository) Create(egg domain.Egg) domain.Egg {
	repository.eggs = append(repository.eggs, egg)

	return egg
}
