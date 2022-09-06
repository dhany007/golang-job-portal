package rest

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/dhany007/golang-job-portal/models"
	"github.com/dhany007/golang-job-portal/models/response"
	"github.com/julienschmidt/httprouter"
)

func (h handler) GetListDresscode(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		ctx    context.Context
		result []models.CompanySubCode
		err    error
	)

	// usecase getlistdresscode
	ctx = r.Context()
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
