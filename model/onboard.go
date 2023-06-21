package model

import (
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
	var validBrid = regexp.MustCompile(`^[a-zA-Z]\d{8}$`)
	return validBrid.MatchString(u.Username)
}

type Roles struct {
	Names []string `json:"role_names"`
}

type Role struct {
	Name string `json:"role_name"`
}
