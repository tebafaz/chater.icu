package models

type TokenReq struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
} //@name TokenRequest

type TokenRes struct {
	Token string `json:"token"`
} //@name TokenResponse
