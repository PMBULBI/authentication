package controller

import (
	"context"
	"github.com/PMBULBI/authentication/model"
	"github.com/PMBULBI/types/schemas"
	"strings"
)

func (c controller) CheckByEmailPass(ctx context.Context, request *model.UserPass) (data schemas.Pendaftaran, err error) {
	return c.auth.GetByEmailPass(ctx, request.Email, request.Password)
}

func (c controller) CheckByPhoneNumPass(ctx context.Context, request *model.UserPass) (data schemas.Pendaftaran, err error) {
	return c.auth.GetByPhoneNumPass(ctx, strings.ReplaceAll(request.PhoneNum, "+", ""), request.Password)
}
