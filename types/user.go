package types

type User struct {
	Id          int    `json:"id" db:"id"`
	Name        string `json:"name" binding:"required"`
	Surname     string `json:"surname" binding:"required"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	AccountId   int    `json:"accountId"`
}
