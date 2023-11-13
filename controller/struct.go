package controller

import (
	"github.com/PMBULBI/authentication"
)

type controller struct {
	auth authentication.Repository
}

func NewController(auth authentication.Repository) authentication.Authentication {
	return &controller{
		auth: auth,
	}
}
