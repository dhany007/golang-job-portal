package rest

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/dhany007/golang-job-portal/models"
	"github.com/dhany007/golang-job-portal/models/response"
	"github.com/dhany007/golang-job-portal/services/utils"
	"github.com/julienschmidt/httprouter"
)

func (h handler) GetListDresscode(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		ctx    context.Context
		result []models.CompanySubCode
		err    error
	)

	ctx = r.Context()

	// usecase getlistdresscode
	result, err = h.companyUsecase.GetListDresscode(ctx)
	if err != nil {
		errCode, _ := strconv.Atoi(err.Error())
		log.Println("[company] [delivery] [GetListDresscode] while companyUsecase.GetListDresscode")
		response.Result(w, errCode)
		return
	}

	// return data
	response.ResultWithData(w, response.SuccesOk, result)
}

func (h handler) GetListBenefitcode(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		ctx    context.Context
		result []models.CompanySubCode
		err    error
	)

	// usecase getbenefitscode
	ctx = r.Context()
	result, err = h.companyUsecase.GetListBenefitcode(ctx)
	if err != nil {
		errCode, _ := strconv.Atoi(err.Error())
		log.Println("[company] [delivery] [GetListBenefitcode] while companyUsecase.GetListBenefitcode")
		response.Result(w, errCode)
		return
	}

	// return data
	response.ResultWithData(w, response.SuccesOk, result)
}

func (h handler) GetListSizecode(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		ctx    context.Context
		result []models.CompanySubCode
		err    error
	)

	// usecase getbenefitscode
	ctx = r.Context()
	result, err = h.companyUsecase.GetListSizecode(ctx)
	if err != nil {
		errCode, _ := strconv.Atoi(err.Error())
		log.Println("[company] [delivery] [GetListSizecode] while companyUsecase.GetListSizecode")
		response.Result(w, errCode)
		return
	}

	// return data
	response.ResultWithData(w, response.SuccesOk, result)
}

func (h handler) UpdateCompany(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		result models.Company
		args   models.CompanyArgument
		err    error
	)
	// TODO: check param id
	companyId := ps.ByName("companyId")
	if err != nil {
		log.Printf("[company] [delivery] [UpdateCompany] while get paramater id, err:%+v\n", err)
		response.ResultError(w, response.ErrorInvalidParameter, err)
		return
	}

	// TODO: check if current user as company is same with company id
	userID := utils.GetAuthorization(r.Context()).ID
	if userID != companyId {
		log.Println("[company] [delivery] [UpdateCompany] while ErrorUnauthorized")
		response.Result(w, response.ErrorUnauthorized)
		return
	}

	// TODO: binding body json
	err = json.NewDecoder(r.Body).Decode(&args)
	if err != nil {
		log.Printf("[company] [delivery] [UpdateCompany] while body binding, err:%+v\n", err)
		response.ResultError(w, response.ErrorBadRequest, err)
		return
	}

	args.ID = companyId

	// TODO: usecase update company
	result, err = h.companyUsecase.UpdateCompany(r.Context(), args)
	if err != nil {
		errCode, _ := strconv.Atoi(err.Error())
		log.Println("[company] [delivery] [UpdateCompany] while companyUsecase.UpdateCompany")
		response.Result(w, errCode)
		return
	}

	// return data
	response.ResultWithData(w, response.SuccesOk, result)
}
