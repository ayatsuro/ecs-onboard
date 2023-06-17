package model

type Namespace struct {
	NsName   string `json:"namespace" binding:"required"`
	Username string `json:"username"  binding:"required"`
}
