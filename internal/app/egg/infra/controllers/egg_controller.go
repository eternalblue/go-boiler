package controllers

import (
	"example.com/go/boiler/internal/app/egg/usecases"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type EggController struct {
	createEgg usecases.CreateEggUseCase
}

func NewEggController(createEgg usecases.CreateEggUseCase) *EggController {
	return &EggController{createEgg: createEgg}
}

func (ctl EggController) Create(ctx *gin.Context) {
	command := usecases.CreateEggCommand{}

	if err := ctx.ShouldBindJSON(&command); err != nil {
		fmt.Print(err)
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	egg, err := ctl.createEgg.Execute(command)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
	}

	ctx.Header("Location", fmt.Sprintf("%s/%v", ctx.FullPath(), "id"))
	ctx.JSON(http.StatusCreated, egg)
}
