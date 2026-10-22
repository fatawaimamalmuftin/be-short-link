package model

type ReqRegis struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type ReqLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type ResLogin struct {
	Token string `json:"token"`
}

type UserFromDB struct {
	ID           int    `db:"id_users"`
	Email        string `db:"email"`
	PasswordHash string `db:"password_hash"`
}
