package controller

import (
	"github.com/PMBULBI/authentication"
)

type upass struct {
	auth authentication.Repository
}

func NewController(auth authentication.Repository) authentication.UPassController {
	return &upass{
		auth: auth,
	}
}
