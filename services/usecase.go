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
}
