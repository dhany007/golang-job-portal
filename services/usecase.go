package services

import (
	"context"

	"github.com/dhany007/golang-job-portal/models"
)

type TestUsecase interface {
	PingTest()
}

type UserUsecase interface {
	Register(ctx context.Context, args models.User) (result models.User, err error)
	Login(ctx context.Context, args models.UserLoginArgument) (result models.UserLoginResponse, err error)
	RefreshToken(ctx context.Context, token string) (result models.UserLoginResponse, err error)
}
