package services

import (
	"context"

	"github.com/dhany007/golang-job-portal/models"
)

type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (result models.User, err error)
	CreateUser(ctx context.Context, args models.User) (result models.User, err error)
}
