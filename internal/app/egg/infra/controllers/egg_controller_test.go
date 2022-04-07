package controllers

import (
	"example.com/go/boiler/internal/app/egg/usecases"
	"example.com/go/boiler/internal/app/egg/usecases/mocks"
	"example.com/go/boiler/pkg/tests"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
)

func init() {

}
func TestEggController_CreateEgg(t *testing.T) {
	createEgg := new(mocks.CreateEggUseCaseMock)
	controller := NewEggController(createEgg)

	t.Run("should create an egg", func(t *testing.T) {
		createEgg.On("Execute", mock.AnythingOfTypeArgument("usecases.CreateEggCommand")).
			Once().
			Return(&usecases.CreateEggResponse{Name: "little egg"}, nil)

		expectedResponse := `{"name":"little egg"}`

		response := tests.DoRequest(http.MethodPost, "", []gin.Param{}, `{}`, controller.Create)

		assert.Equal(t, http.StatusCreated, response.Code)
		assert.Equal(t, expectedResponse, response.Body.String())
	})
}
