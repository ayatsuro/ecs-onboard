package handler

import (
	"ecs-onboard/model"
	"ecs-onboard/service"
	"github.com/gin-gonic/gin"
)

// OnboardNs
// @Tags Namespace
// @Summary onboard a namespace
// @Description Creates in Dell ECS: a namespace, a IAM user RW and a IAM user RO and store their secret access keys in Vault
// @Accept json
// @Produce json
// @param ns body model.Namespace true "the namespace to onboard"
// @Router /namespace/onboard [post]
func OnboardNs(ctx *gin.Context) {
	var ns model.Namespace
	if err := ctx.BindJSON(&ns); err != nil {
		ctx.AbortWithError(400, err)
		return
	}
	path := "/namespace/onboard"
	status, err := service.ReqVault("POST", path, ns, nil)
	if status != 200 {
		ctx.AbortWithStatusJSON(status, err.Error())
	}
}

func MigrateNs(ctx *gin.Context) {
	var ns model.Namespace
	if err := ctx.BindJSON(&ns); err != nil {
		ctx.AbortWithError(400, err)
	}
}

func Test(ctx *gin.Context) {
	stat, err := service.ReqVault("GET", "/sys/policies/acl?list=true", nil, nil)
	if stat != 200 {
		ctx.AbortWithError(stat, err)
	}
}
