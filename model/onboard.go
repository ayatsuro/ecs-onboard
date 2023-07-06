package model

import (
	"encoding/json"
	"regexp"
	"strings"
)

type Role struct {
	Username  string `json:"username" binding:"required"`
	Namespace string `json:"namespace" binding:"required"`
	SafeId    string `json:"safe_id"`
}

func (u Role) IsBrid() bool {
	return IsBrid(u.Username)
}

func (u Role) GetSafeId() string {
	if u.SafeId != "" {
		return u.SafeId
	}
	// skylight naming convention
	safeId, _, _ := strings.Cut(u.Namespace, "-")
	return strings.ToLower(safeId)
}

func IsBrid(name string) bool {
	var validBrid = regexp.MustCompile(`^[a-zA-Z]\d{8}$`)
	return validBrid.MatchString(name)
}

func (u Role) ToJwtAuthRole() JwtAuthRole {
	bc, _ := json.Marshal(&map[string]string{"BRID": u.Username})
	return JwtAuthRole{
		BoundAudiences: []string{"OBJECT_STORE", "CSM_DEV", "CSM_INT"},
		BoundClaims:    bc,
		RoleType:       "jwt",
		TokenPolicies:  []string{"object-store/creds/" + u.RoleName()},
		TokenTtl:       0,
		TokenMaxTtl:    0,
		UserClaim:      "sub",
	}
}

func (u Role) RoleName() string {
	return u.GetSafeId() + "_" + u.Username
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
