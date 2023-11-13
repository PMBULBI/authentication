package repository

import (
	"github.com/PMBULBI/authentication"
)

type repository struct {
	pendaftaran authentication.UserPassRepository
}

func NewRepository(pendaftaran authentication.UserPassRepository) authentication.Authentication {
	return &repository{
		pendaftaran: pendaftaran,
	}
}
