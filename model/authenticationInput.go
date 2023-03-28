package model

type AuthenticationInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AccessDetails struct {
	AccessUuid string
	UserId     uint64
}
