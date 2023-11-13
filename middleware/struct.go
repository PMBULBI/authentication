package middleware

import (
	"github.com/PMBULBI/authentication"
)

type AuthMiddleware struct {
	exist       authentication.UserExist
	TokenHeader string
}
