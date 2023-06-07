package types

type Account struct {
	Id       int    `json:"id" db:"id"`
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}
