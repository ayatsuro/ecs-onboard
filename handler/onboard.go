package handler

import (
	"ecs-onboard/model"
	"ecs-onboard/service"
	"github.com/gin-gonic/gin"
	"regexp"
	"strings"
)

const objectStore = "/object-store"

// OnboardNamespace
// @Tags Namespace
// @Summary onboards a namespace
// @Description In Dell ECS, creates a namespace, a IAM user and an AccessKey. In Vault, stores the Secret Access Key
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
	if ns.SafeId == "" {
		safeId, _, _ := strings.Cut(ns.Namespace, "-")
		ns.SafeId = safeId
	}
	ns.SafeId = strings.ToLower(ns.SafeId)
	// check there is a safe v4
	// talk to vault regionally

	// onboard namespace
	var onboarded model.RoleName
	path := objectStore + "/namespace/onboard"
	status, err := service.ReqVault("POST", path, ns, &onboarded)
	if status != 200 {
		ctx.AbortWithStatusJSON(status, gin.H{"error": err.Error()})
		return
	}
	// create policy for role
	status, err = service.CreatePolicy(onboarded.RoleName)
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
// @param ns body model.MigrateNamespace true "the namespace to migrate"
// @Router /namespace/migrate [post]
func MigrateNamespace(ctx *gin.Context) {
	var ns model.MigrateNamespace
	if err := ctx.BindJSON(&ns); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	if ns.SafeId == "" {
		safeId, _, _ := strings.Cut(ns.Namespace, "-")
		ns.SafeId = safeId
	}
	ns.SafeId = strings.ToLower(ns.SafeId)
	// check there is a safe v4
	// talk to vault regionally
	path := objectStore + "/namespace/migrate"
	var roles model.RoleNames
	status, err := service.ReqVault("POST", path, ns, &roles)
	if status != 200 {
		ctx.AbortWithStatusJSON(status, gin.H{"error": err.Error()})
		return
	}
	for _, roleName := range roles.RoleNames {
		status, err = service.CreatePolicy(roleName)
		if status != 200 {
			ctx.AbortWithStatusJSON(status, gin.H{"error": err.Error()})
			return
		}
	}

}

// OnboardBrid
// @Tags Brid
// @Summary onboard a brid to a namespace as IAM user
// @Description In Dell ECS, creates a IAM user and an AccessKey. In Vault, stores the secret access keys and creates a JWT auth role bound to the Brid
// @Accept json
// @Produce json
// @param brid body model.IamUser true "the user to onboard"
// @Router /bird/onboard [post]
func OnboardBrid(ctx *gin.Context) {
	var user model.IamUser
	if err := ctx.BindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	var validBrid = regexp.MustCompile(`^[a-zA-Z]\d{8}$`)
	if !validBrid.MatchString(user.Username) {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "invalid brid (1 letter followed by 8 digits)"})
		return
	}

	// check brid existence in SF dump users
	path := objectStore + "/iam-user"
	status, err := service.ReqVault("POST", path, user, nil)
	if status != 200 {
		ctx.AbortWithStatusJSON(status, gin.H{"error": err.Error()})
		return
	}
	//create policy
	// create jwt
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
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	path := objectStore + "/iam-user"
	status, err := service.ReqVault("POST", path, user, nil)
	if status != 200 {
		ctx.AbortWithStatusJSON(status, gin.H{"error": err.Error()})
		return
	}
	// create policy
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
