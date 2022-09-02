package rest

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/dhany007/golang-job-portal/models"
	"github.com/dhany007/golang-job-portal/models/response"
	"github.com/dhany007/golang-job-portal/services/utils"

	"github.com/julienschmidt/httprouter"
)

func (h handler) Register(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		args models.UserArguments
		err  error
		user models.User
		ctx  context.Context
	)

	// body json
	err = json.NewDecoder(r.Body).Decode(&args)
	if err != nil {
		log.Printf("[user] [delivery] [Register] while body binding, err:%+v\n", err)
		response.ResultError(w, response.ErrorBadRequest, err)
		return
	}

	// validate args
	_, err = govalidator.ValidateStruct(args)
	if err != nil {
		log.Printf("[user] [delivery] [Register] while ValidateStruct, err:%+v\n", err)
		response.ResultError(w, response.ErrorValidation, err)
		return
	}

	// generate hash_password
	hashPass, err := utils.HashPassword(args.Password)
	if err != nil {
		log.Printf("[user] [delivery] [Register] while HashPassword, err:%+v\n", err)
		response.ResultError(w, response.ErrorServerError, err)
		return
	}

	// reinitialize data user
	user.Email = args.Email
	user.Password = hashPass
	user.Role = args.Role

	// usecase
	ctx = r.Context()
	user, err = h.userUsecase.Register(ctx, user)
	if err != nil {
		errCode, _ := strconv.Atoi(err.Error())
		log.Println("[user] [delivery] [Register] while userUsecase.Register")
		response.Result(w, errCode)
		return
	}

	// return response ok
	response.ResultWithData(w, response.SuccesOk, user)
}
