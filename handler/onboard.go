package handler

import (
	"ecs-onboard/model"
	"ecs-onboard/service"
	"github.com/gin-gonic/gin"
)

const objectStore = "/object-store"

// OnboardNamespace
// @Tags Namespace
// @Summary Creates an IAM user, an access key, stores it in a Vault role, creates a Vault policy
// @Description If the safe_id is omitted in the payload, it will be derived from the namespace (first part of a split on '-')
// @Accept json
// @Produce json
// @param ns body model.OnboardNamespace true "the namespace to onboard"
// @Router /namespace/onboard [post]
func OnboardNamespace(ctx *gin.Context) {
	var ns model.OnboardNamespace
	if err := ctx.BindJSON(&ns); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	if ok := service.CheckSafeId(&ns); !ok {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "unknown safe " + ns.SafeId})
		return
	}
	// check there is a safe v4
	// talk to vault regionally

	// onboard namespace
	var role model.Role
	path := objectStore + "/namespace/onboard"
	status, err := service.ReqVault("POST", path, ns, &role)
	if status != 200 {
		ctx.AbortWithStatusJSON(status, gin.H{"error": err.Error()})
		return
	}
	// create policy for role
	status, err = service.CreatePolicy(role.Name)
	if status != 200 {
		ctx.AbortWithStatusJSON(status, gin.H{"error": err.Error()})
		return
	}

}

// MigrateNamespace
// @Tags Namespace
// @Summary Fetches all the users of an ECS namespace. For native users, creates a IAM user, an access key, stores it in a Vault role, creates a Vault policy. For IAM users, creates an access key, stores it in a Vault role, creates a Vault policy.
// @Description If the safe_id is omitted in the payload, it will be derived from the namespace name (first part of a split on '-')
// @Accept json
// @Produce json
// @param ns body model.MigrateNamespace true "the namespace to migrate"
// @Router /namespace/migrate [post]
func MigrateNamespace(ctx *gin.Context) {
	var ns model.MigrateNamespace
	if err := ctx.BindJSON(&ns); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	if ok := service.CheckSafeId(&ns); !ok {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "unknown safe " + ns.SafeId})
		return
	}
	// check there is a safe v4
	// talk to vault regionally
	path := objectStore + "/namespace/migrate"
	var roles model.Roles
	status, err := service.ReqVault("POST", path, ns, &roles)
	if status != 200 {
		ctx.AbortWithStatusJSON(status, gin.H{"error": err.Error()})
		return
	}
	for _, roleName := range roles.Names {
		status, err = service.CreatePolicy(roleName)
		if status != 200 {
			ctx.AbortWithStatusJSON(status, gin.H{"error": err.Error()})
			return
		}
	}

}

// OnboardIamUser
// @Tags User
// @Summary Creates a IAM user in an ECS namespace, creates an access key, stores it in a Vault role, creates a Vault policy. And if the username is a BRID, creates a JWT authrole.
// @Description If the safe_id is omitted in the payload, it will be derived from the namespace name (first part of a split on '-')
// @Accept json
// @Produce json
// @param user body model.IamUser true "the user to onboard"
// @Router /user [post]
func OnboardIamUser(ctx *gin.Context) {
	var user model.IamUser
	if err := ctx.BindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	if ok := service.CheckSafeId(&user); !ok {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "unknown safe " + user.SafeId})
		return
	}
	path := objectStore + "/role"
	var role model.Role
	status, err := service.ReqVault("POST", path, user, &role)
	if status != 200 {
		ctx.AbortWithStatusJSON(status, gin.H{"error": err.Error()})
		return
	}
	status, err = service.CreatePolicy(role.Name)
	if status != 200 {
		ctx.AbortWithStatusJSON(status, gin.H{"error": err.Error()})
		return
	}
	if user.IsBrid() {
		//status, err = service.CreateJwtAuthRole(role.Name)
		//if status != 200 {
		//	ctx.AbortWithStatusJSON(status, gin.H{"error": err.Error()})
		//	return
		//}
	}
}

// DeleteIamUser
// @Tags User
// @Summary Deletes the IAM user, the Vault role and policy, and the JWT authrole if any
// @Produce json
// @param username path string true "the IAM user to delete"
// @Router /user/{username} [delete]
func DeleteIamUser(ctx *gin.Context) {
	user := ctx.Param("username")
	path := objectStore + "/role/" + user
	status, err := service.ReqVault("DELETE", path, nil, nil)
	if status != 200 {
		ctx.AbortWithStatusJSON(status, gin.H{"error": err.Error()})
		return
	}
	// delete policy
	// delete jwt if brid
}

func Test(ctx *gin.Context) {
	stat, err := service.ReqVault("GET", "/sys/policies/acl?list=true", nil, nil)
	if stat != 200 {
		ctx.AbortWithError(stat, err)
	}
}
