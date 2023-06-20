package model

type Namespace struct {
	Namespace string `json:"namespace" binding:"required"`
	Username  string `json:"username"  binding:"required"`
}

type IamUser struct {
	Username  string `json:"username" binding:"required"`
	Namespace string `json:"namespace" binding:"required"`
}

type OnboardedNamespace struct {
	Namespace   string `json:"namespace"`
	Username    string `json:"username"`
	AccessKeyId string `json:"access_key_id"`
}

func (n OnboardedNamespace) ToRoleName() string {
	return n.Namespace + "_" + n.Username + "_" + n.AccessKeyId
}
