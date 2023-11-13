package authentication

import (
	"context"
	"github.com/PMBULBI/types/schemas"
	"github.com/gofiber/fiber/v2"
)

type UserPassRepository interface {
	GetByEmailPass(ctx context.Context, email, password string) (data schemas.Pendaftaran, err error)
	GetByPhoneNumPass(ctx context.Context, phoneNum, password string) (data schemas.Pendaftaran, err error)
}

type UserExist interface {
	CheckUserExists(ctx context.Context, email, phoneNum string) (data schemas.Pendaftaran, err error)
}

type Repository struct {
	Authentication
}

type Controller interface {
	CheckIsExist(ctx *fiber.Ctx) (data schemas.Pendaftaran, err error)
}

type Authentication interface {
	CheckByEmailPass(ctx context.Context, email, password string) (data schemas.Pendaftaran, err error)
	CheckByPhoneNumPass(ctx context.Context, phoneNum, password string) (data schemas.Pendaftaran, err error)
}
