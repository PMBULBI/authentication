package model

type UserPass struct {
	Email    string `json:"email"`
	PhoneNum string `json:"phone_num"`
	Password string `json:"password" validate:"required"`
}
