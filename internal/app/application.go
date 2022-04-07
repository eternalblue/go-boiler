package app

import (
	"example.com/go/boiler/internal/app/egg/infra/controllers"
	"example.com/go/boiler/internal/app/egg/infra/repositories/memory"
	"example.com/go/boiler/internal/app/egg/usecases"
	"example.com/go/boiler/internal/app/infra/config"
)

// Application is in charge of handling all dependency injection.
type Application struct {
	EggController *controllers.EggController
}

func NewApplication(config *config.Configuration) *Application {
	eggRepository := memory.NewEggRepository(config)

	return &Application{
		EggController: controllers.NewEggController(
			usecases.NewCreateEggUseCase(eggRepository),
		),
	}
}
