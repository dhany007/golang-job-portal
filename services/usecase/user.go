package usecase

import (
	"context"
	"log"

	"github.com/dhany007/golang-job-portal/models"
	"github.com/dhany007/golang-job-portal/models/response"
	"github.com/dhany007/golang-job-portal/services"
	"github.com/google/uuid"
)

type userUsecase struct {
	repo services.UserRepository
}

func NewUserUsecase(repo services.UserRepository) services.UserUsecase {
	return userUsecase{repo}
}

func (u userUsecase) Register(ctx context.Context, args models.User) (result models.User, err error) {
	var (
		user models.User
	)

	// chek if user exists
	user, err = u.repo.GetUserByEmail(ctx, args.Email)
	if err != nil {
		log.Printf("[user] [usecase] [Register] while GetUserByEmail, err:%+v\n", err)
		return result, err
	}

	if user.ID != "" {
		err = response.NewErrork(response.ErrorRegisEmail)
		log.Printf("[user] [usecase] [Register] while GetUserByEmail, email:%+v\n", args.Email)
		return result, err
	}

	// generate uuid users
	uuidUser := uuid.New().String()
	args.ID = uuidUser

	// save data user, check role: 1=company, 2=candidate
	result, err = u.repo.CreateUser(ctx, args)
	if err != nil {
		err = response.NewErrork(response.ErrorServerError)
		log.Printf("[user] [usecase] [Register] while repo.CreateUser, email:%+v\n", args.Email)
		return result, err
	}

	return result, nil
}
