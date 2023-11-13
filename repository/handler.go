package repository

import (
	"context"
	"github.com/PMBULBI/types/schemas"
)

func (r *repository) CheckByEmailPass(ctx context.Context, email, password string) (data schemas.Pendaftaran, err error) {
	return r.pendaftaran.GetByEmailPass(ctx, email, password)
}

func (r *repository) CheckByPhoneNumPass(ctx context.Context, phoneNum, password string) (data schemas.Pendaftaran, err error) {
	return r.pendaftaran.GetByPhoneNumPass(ctx, phoneNum, password)
}
