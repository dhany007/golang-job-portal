package usecase

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/dhany007/golang-job-portal/models"
	"github.com/dhany007/golang-job-portal/models/response"
	"github.com/dhany007/golang-job-portal/services"
	"github.com/dhany007/golang-job-portal/services/utils"
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

func (u userUsecase) Login(ctx context.Context, args models.UserLoginArgument) (result models.UserLoginResponse, err error) {
	var (
		user models.User
	)

	// chek if user exists
	user, err = u.repo.GetUserByEmail(ctx, args.Email)
	if err != nil {
		log.Printf("[user] [usecase] [Login] while GetUserByEmail, err:%+v\n", err)
		return result, err
	}

	if user.ID == "" {
		err = response.NewErrork(response.ErrorUserNotFound)
		log.Printf("[user] [usecase] [Login] while ErrorUserNotFound, email:%+v\n", args.Email)
		return result, err
	}

	// compare password
	ok := utils.ComparePassword([]byte(user.Password), []byte(args.Password))
	if !ok {
		err = response.NewErrork(response.ErrorPwdNotMatch)
		log.Printf("[user] [usecase] [Login] while ErrorPwdNotMatch, email:%+v\n", args.Email)
		return result, err
	}

	refreshTokenExpiresIn, _ := strconv.Atoi(utils.GetEnv("RefreshTokenExpiresIn", "43200"))
	accessTokenExpiresIn, _ := strconv.Atoi(utils.GetEnv("AccessTokenExpiresIn", "20"))

	// generate refresh token
	refreshToken, err := utils.GenerateToken(time.Duration(refreshTokenExpiresIn), user.ID, user.Email, user.Role)
	if err != nil {
		err = response.NewErrork(response.ErrorServerError)
		log.Printf("[user] [usecase] [Login] while ErrorServerError, email:%+v\n", args.Email)
		return result, err
	}

	// generate access token
	accessToken, err := utils.GenerateToken(time.Duration(accessTokenExpiresIn), user.ID, user.Email, user.Role)
	if err != nil {
		err = response.NewErrork(response.ErrorServerError)
		log.Printf("[user] [usecase] [Login] while ErrorServerError, email:%+v\n", args.Email)
		return result, err
	}

	// return result
	result = models.UserLoginResponse{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	}

	return result, nil
}

func (u userUsecase) RefreshToken(ctx context.Context, token string) (result models.UserLoginResponse, err error) {
	var (
		user models.User
	)

	// validate token
	data, err := utils.ValidateToken(token)
	if err != nil {
		errCode, _ := strconv.Atoi(err.Error())
		return result, response.NewErrork(errCode)
	}

	// check if user exists
	userId := data["id"].(string)
	user, err = u.repo.GetUserById(ctx, userId)
	if err != nil {
		log.Printf("[user] [usecase] [RefreshToken] while GetUserById, id:%+v\n", userId)
		return result, err
	}

	if user.ID == "" {
		err = response.NewErrork(response.ErrorUserNotFound)
		log.Printf("[user] [usecase] [RefreshToken] while ErrorUserNotFound, id:%+v\n", userId)
		return result, err
	}

	// generate token
	accessTokenExpiresIn, _ := strconv.Atoi(utils.GetEnv("AccessTokenExpiresIn", "20"))

	// generate access token
	accessToken, err := utils.GenerateToken(time.Duration(accessTokenExpiresIn), user.ID, user.Email, user.Role)
	if err != nil {
		err = response.NewErrork(response.ErrorServerError)
		log.Println("[user] [usecase] [RefreshToken] while ErrorServerError")
		return result, err
	}

	// return result
	result = models.UserLoginResponse{
		RefreshToken: "",
		AccessToken:  accessToken,
	}

	return result, nil
}
