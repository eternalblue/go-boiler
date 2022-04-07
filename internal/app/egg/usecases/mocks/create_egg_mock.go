package mocks

import (
	"example.com/go/boiler/internal/app/egg/usecases"
	"github.com/stretchr/testify/mock"
)

type CreateEggUseCaseMock struct {
	mock.Mock
}

func (m *CreateEggUseCaseMock) Execute(command usecases.CreateEggCommand) (*usecases.CreateEggResponse, error) {
	args := m.Called(command)

	if args.Get(0) == nil {
		return nil, args.Get(1).(error)
	}

	return args.Get(0).(*usecases.CreateEggResponse), nil
}
