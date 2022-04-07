package rest

import (
	"example.com/go/boiler/internal/app"
	"github.com/gin-gonic/gin"
)

// MapRoutes on a gin.Engine from a core app.Application.
func MapRoutes(router *gin.Engine, app *app.Application) {
	v1 := router.Group("/v1")
	{
		egg := v1.Group("/eggs")
		{
			egg.POST("", app.EggController.Create)
		}
	}
}
