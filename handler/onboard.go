package handler

import (
	"ecs-onboard/model"
	"ecs-onboard/service"
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

func Test(ctx *gin.Context) {
	if err := service.ReqVault("GET", "/sys/policies/acl?list=true", nil, nil); err != nil {
		ctx.AbortWithError(500, err)
	}
}
