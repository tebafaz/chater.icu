package models

type TokenReq struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

type TokenRes struct {
	Token string `json:"token"`
}
