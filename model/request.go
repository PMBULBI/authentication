package model

type UserPass struct {
	Email    string `json:"email" validate:"email"`
	PhoneNum string `json:"phone_num" validate:"e164"`
	Password string `json:"password" validate:"required"`
}
