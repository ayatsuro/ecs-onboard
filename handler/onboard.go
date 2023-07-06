package handler

import (
	"ecs-onboard/model"
	"ecs-onboard/service"
	"github.com/gin-gonic/gin"
	"strings"
)

const objectStore = "/object-store"

// CreateRole
// @Tags User
// @Summary Creates a IAM user in an ECS namespace, creates an access key, stores it in a Vault role, creates a Vault policy. And if the username is a BRID, creates a JWT authrole.
// @Description If the safe_id is omitted in the payload, it will be derived from the namespace name (first part of a split on '-')
// @Accept json
// @Produce json
// @param role body model.Role true "the user to onboard"
// @Router /role [post]
func CreateRole(ctx *gin.Context) {
	var role model.Role
	if err := ctx.BindJSON(&role); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	if ok := service.CheckSafeId(role.GetSafeId()); !ok {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "unknown safe " + role.GetSafeId()})
		return
	}
	path := objectStore + "/role/" + role.RoleName()
	status, err := service.ReqVault("POST", path, role, nil)
	if status != 200 {
		ctx.AbortWithStatusJSON(status, gin.H{"error": err.Error()})
		return
	}
	status, err = service.CreatePolicy(role.RoleName())
	if status != 200 {
		ctx.AbortWithStatusJSON(status, gin.H{"error": err.Error()})
		return
	}
	if role.IsBrid() {
		status, err = service.CreateJwtAuthRole(role)
		if status != 200 {
			ctx.AbortWithStatusJSON(status, gin.H{"error": err.Error()})
			return
		}
	}
}

// DeleteRole
// @Tags User
// @Summary Deletes the IAM user, the Vault role and policy, and the JWT authrole if any
// @Produce json
// @param roleName path string true "the role name, in the form <safeId>_<iamUserName> to delete"
// @Router /role/{roleName} [delete]
func DeleteRole(ctx *gin.Context) {
	roleName := ctx.Param("roleName")
	path := objectStore + "/role/" + roleName
	status, err := service.ReqVault("DELETE", path, nil, nil)
	if status != 200 {
		ctx.AbortWithStatusJSON(status, gin.H{"error": err.Error()})
		return
	}
	status, err = service.DeletePolicy(roleName)
	if status != 200 {
		ctx.AbortWithStatusJSON(status, gin.H{"error": err.Error()})
		return
	}
	_, brid, _ := strings.Cut(roleName, "_")
	if model.IsBrid(brid) {
		status, err = service.DeleteJwtAuthRole(roleName)
		if status != 200 {
			ctx.AbortWithStatusJSON(status, gin.H{"error": err.Error()})
			return
		}
	}
}

func Test(ctx *gin.Context) {
	stat, err := service.ReqVault("GET", "/sys/policies/acl?list=true", nil, nil)
	if stat != 200 {
		ctx.AbortWithError(stat, err)
	}
}
