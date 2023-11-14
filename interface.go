package authentication

import (
	"context"
	"github.com/PMBULBI/authentication/model"
	"github.com/PMBULBI/types/schemas"
)

type UserPassRepository interface {
	GetByEmailPass(ctx context.Context, email, password string) (data schemas.Pendaftaran, err error)
	GetByPhoneNumPass(ctx context.Context, phoneNum, password string) (data schemas.Pendaftaran, err error)
}

type UserExist interface {
	CheckUserExists(ctx context.Context, email, phoneNum string) (data schemas.Pendaftaran, err error)
}

type Repository struct {
	UserPassRepository
	UserExist
}

type Controller interface {
	CheckByEmailPass(ctx context.Context, request *model.UserPass) (data schemas.Pendaftaran, err error)
	CheckByPhoneNumPass(ctx context.Context, request *model.UserPass) (data schemas.Pendaftaran, err error)
}
