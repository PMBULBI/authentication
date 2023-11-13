package controller

import (
	"context"
	"github.com/PMBULBI/types/schemas"
)

func (c controller) CheckByEmailPass(ctx context.Context, email, password string) (data schemas.Pendaftaran, err error) {
	//TODO implement me
	panic("implement me")
}

func (c controller) CheckByPhoneNumPass(ctx context.Context, phoneNum, password string) (data schemas.Pendaftaran, err error) {
	//TODO implement me
	panic("implement me")
}
