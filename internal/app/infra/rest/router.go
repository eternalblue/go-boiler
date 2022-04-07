package rest

import (
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

var production = false

// DefaultRouter creates and configures a gin.Engine
func DefaultRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.RecoveryWithWriter(os.Stderr))

	if production {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router.NoRoute(noRouteHandler)
	pprof.Register(router)

	return router
}

func noRouteHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("resource %s not found", ctx.Request.URL.Path)})
}
