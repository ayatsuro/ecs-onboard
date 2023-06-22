package model

import (
	"encoding/json"
	"regexp"
)

type OnboardNamespace struct {
	Namespace string `json:"namespace" binding:"required"`
	Username  string `json:"username"  binding:"required"`
	SafeId    string `json:"safe_id"`
}

type MigrateNamespace struct {
	Namespace string `json:"namespace" binding:"required"`
	SafeId    string `json:"safe_id"`
}

type IamUser struct {
	Username  string `json:"username" binding:"required"`
	Namespace string `json:"namespace" binding:"required"`
	SafeId    string `json:"safe_id"`
}

func (u IamUser) IsBrid() bool {
	return IsBrid(u.Username)
}

func IsBrid(name string) bool {
	var validBrid = regexp.MustCompile(`^[a-zA-Z]\d{8}$`)
	return validBrid.MatchString(name)
}

func (u IamUser) ToJwtAuthRole() JwtAuthRole {
	bc, _ := json.Marshal(&map[string]string{"BRID": u.Username})
	return JwtAuthRole{
		BoundAudiences: []string{"OBJECT_STORE", "CSM_DEV", "CSM_INT"},
		BoundClaims:    bc,
		RoleType:       "jwt",
		TokenPolicies:  []string{"object-store/" + u.RoleName()},
		TokenTtl:       0,
		TokenMaxTtl:    0,
		UserClaim:      "sub",
	}
}

func (u IamUser) RoleName() string {
	return u.SafeId + "_" + u.Username
}

type JwtAuthRole struct {
	BoundAudiences  []string        `json:"bound_audiences"`
	BoundClaims     json.RawMessage `json:"bound_claims"`
	BoundClaimsType string          `json:"bound_claims_type,omitempty"`
	RoleType        string          `json:"role_type"`
	TokenPolicies   []string        `json:"token_policies"`
	TokenTtl        int32           `json:"token_ttl"`
	TokenMaxTtl     int32           `json:"token_max_ttl"`
	UserClaim       string          `json:"user_claim"`
}

type Roles struct {
	Names []string `json:"role_names"`
}

type Role struct {
	Name string `json:"role_name"`
}
