package handler

import (
	"ecs-onboard/model"
	"github.com/gin-gonic/gin"
)

func OnboardNs(ctx *gin.Context) {
	var ns model.Namespace
	if err := ctx.BindJSON(&ns); err != nil {
		ctx.AbortWithError(400, err)
	}

}

func MigrateNs(ctx *gin.Context) {
	var ns model.Namespace
	if err := ctx.BindJSON(&ns); err != nil {
		ctx.AbortWithError(400, err)
	}

}
