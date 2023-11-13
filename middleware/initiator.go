package middleware

import "github.com/PMBULBI/authentication"

func NewAuthMiddleware(exist authentication.UserExist, tokenHeader string) *AuthMiddleware {
	return &AuthMiddleware{
		exist:       exist,
		TokenHeader: tokenHeader,
	}
}
