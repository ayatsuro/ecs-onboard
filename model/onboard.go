package model

type Namespace struct {
	Namespace string `json:"namespace" binding:"required"`
	Username  string `json:"username"  binding:"required"`
}

type Brid struct {
	Brid      string `json:"brid" binding:"required"`
	Namespace string `json:"namespace" binding:"required"`
}

type IamUser struct {
	Username  string `json:"username" binding:"required"`
	Namespace string `json:"namespace" binding:"required"`
}
