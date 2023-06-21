package model

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

type RoleNames struct {
	RoleNames []string `json:"role_names"`
}

type RoleName struct {
	RoleName string `json:"role_name"`
}
