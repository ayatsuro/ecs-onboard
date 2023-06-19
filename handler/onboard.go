package handler

import (
	"ecs-onboard/model"
	"ecs-onboard/service"
	"github.com/gin-gonic/gin"
)

// OnboardNamespace
// @Tags Namespace
// @Summary onboards a namespace
// @Description In Dell ECS, creates a namespace, a IAM user and an AccessKey. In Vault, stores the Secret Access Key
// @Accept json
// @Produce json
// @param ns body model.Namespace true "the namespace to onboard"
// @Router /namespace/onboard [post]
func OnboardNamespace(ctx *gin.Context) {
	var ns model.Namespace
	if err := ctx.BindJSON(&ns); err != nil {
		ctx.AbortWithError(400, err)
		return
	}
	path := "/namespace/onboard"
	status, err := service.ReqVault("POST", path, ns, nil)
	if status != 200 {
		ctx.AbortWithStatusJSON(status, gin.H{"error": err.Error()})
		return
	}
}

// MigrateNamespace
// @Tags Namespace
// @Summary migrates a namespace
// @Description In Dell ECS, creates a IAM user (and AccessKey) for the Native users, creates a second AccessKey for existing IAM users. In Vault, stores the Secret Access Keys
// @Accept json
// @Produce json
// @param ns path string true "the namespace to migrate"
// @Router /namespace/migrate/{namespace} [post]
func MigrateNamespace(ctx *gin.Context) {
	ns := ctx.Param("namespace")
	path := "/namespace/migrate/" + ns
	status, err := service.ReqVault("POST", path, ns, nil)
	if status != 200 {
		ctx.AbortWithStatusJSON(status, gin.H{"error": err.Error()})
		return
	}

}

// DeleteNamespace
// @Tags Namespace
// @Summary delete a namespace
// @Description In ECS, deletes the namespace and associate users. In Vault, deletes the roles and JWT auth roles
// @Produce json
// @param ns path string true "the namespace to delete"
// @Router /namespace/onboard/{namespace} [delete]
func DeleteNamespace(ctx *gin.Context) {
	ns := ctx.Param("namespace")
	path := "/namespace/onboard/" + ns
	status, err := service.ReqVault("DELETE", path, ns, nil)
	if status != 200 {
		ctx.AbortWithStatusJSON(status, gin.H{"error": err.Error()})
		return
	}
}

// OnboardBrid
// @Tags Brid
// @Summary onboard a brid to a namespace as IAM user
// @Description In Dell ECS, creates a IAM user and an AccessKey. In Vault, stores the secret access keys and creates a JWT auth role bound to the Brid
// @Accept json
// @Produce json
// @param brid body model.Brid true "the user to onboard"
// @Router /bird/onboard [post]
func OnboardBrid(ctx *gin.Context) {
	var user model.IamUser
	if err := ctx.BindJSON(&user); err != nil {
		ctx.AbortWithError(400, err)
		return
	}
	// check brid existence in SF dump users
	path := "/iam-user"
	status, err := service.ReqVault("POST", path, user, nil)
	if status != 200 {
		ctx.AbortWithStatusJSON(status, gin.H{"error": err.Error()})
		return
	}
}

// OnboardIamUser
// @Tags IamUser
// @Summary onboard a IAM user in a namespace
// @Description In Dell ECS, creates a IAM user and an AccessKey. In Vault, stores the secret access keys and creates a JWT auth role bound to the Brid
// @Accept json
// @Produce json
// @param brid body model.IamUser true "the user to onboard"
// @Router /iamuser/onboard [post]
func OnboardIamUser(ctx *gin.Context) {
	var user model.IamUser
	if err := ctx.BindJSON(&user); err != nil {
		ctx.AbortWithError(400, err)
		return
	}
	path := "/iam-user"
	status, err := service.ReqVault("POST", path, user, nil)
	if status != 200 {
		ctx.AbortWithStatusJSON(status, gin.H{"error": err.Error()})
		return
	}
}

// DeleteIamUser
// @Tags IamUser
// @Summary delete a IAM user
// @Description In Dell ECS, deletes the IAM user. In Vault, deletes the role and the JWT auth role (if any)
// @Produce json
// @param brid path string true "the user to delete"
// @Router /iamuser/{username} [delete]
func DeleteIamUser(ctx *gin.Context) {
	user := ctx.Param("username")
	path := "/role/" + user
	status, err := service.ReqVault("DELETE", path, nil, nil)
	if status != 200 {
		ctx.AbortWithStatusJSON(status, gin.H{"error": err.Error()})
		return
	}
}

func Test(ctx *gin.Context) {
	stat, err := service.ReqVault("GET", "/sys/policies/acl?list=true", nil, nil)
	if stat != 200 {
		ctx.AbortWithError(stat, err)
	}
}
