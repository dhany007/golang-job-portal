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
	// check param id
	companyId := ps.ByName("companyId")
	if companyId == "" {
		log.Println("[company] [delivery] [UpdateCompany] while get paramater id")
		response.Result(w, response.ErrorInvalidParameter)
		return
	}

	// binding body json
	err = json.NewDecoder(r.Body).Decode(&args)
	if err != nil {
		log.Printf("[company] [delivery] [UpdateCompany] while body binding, err:%+v\n", err)
		response.ResultError(w, response.ErrorBadRequest, err)
		return
	}

	// validate args
	_, err = govalidator.ValidateStruct(args)
	if err != nil {
		log.Printf("[company] [delivery] [UpdateCompany] while ValidateStruct, err:%+v\n", err)
		response.ResultError(w, response.ErrorValidation, err)
		return
	}

	// check if current user as company is same with company id
	userID := utils.GetAuthorization(r.Context()).ID
	if userID != companyId {
		log.Println("[company] [delivery] [UpdateCompany] while ErrorUnauthorized")
		response.Result(w, response.ErrorUnauthorized)
		return
	}

	args.ID = companyId

	// usecase update company
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

func (h handler) GetListCompanies(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		ctx    context.Context
		result models.ListCompanies
		args   models.ListData
		err    error
	)

	// define args
	page := utils.Utint(r.URL.Query().Get("page"), 1)                  // default 1
	itemPerPage := utils.Utint(r.URL.Query().Get("item_per_page"), 10) // default 10

	args.ItemPerPage = itemPerPage
	args.Page = page
	args.Limit = itemPerPage

	// usecase list companies
	ctx = r.Context()
	result, err = h.companyUsecase.GetListCompanies(ctx, args)
	if err != nil {
		errCode, _ := strconv.Atoi(err.Error())
		log.Println("[company] [delivery] [GetListCompanies] while companyUsecase.GetListCompanies")
		response.Result(w, errCode)
		return
	}

	// return data
	response.ResultWithData(w, response.SuccesOk, result)
}

func (h handler) GetDetailCompany(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		result models.Company
		err    error
	)
	// check param id
	companyId := ps.ByName("companyId")
	if companyId == "" {
		log.Println("[company] [delivery] [DetailCompany] while get paramater id")
		response.Result(w, response.ErrorInvalidParameter)
		return
	}

	// usecase detail company
	result, err = h.companyUsecase.GetDetailCompany(r.Context(), companyId)
	if err != nil {
		errCode, _ := strconv.Atoi(err.Error())
		log.Println("[company] [delivery] [DetailCompany] while companyUsecase.GetDetailCompany")
		response.Result(w, errCode)
		return
	}

	// return data
	response.ResultWithData(w, response.SuccesOk, result)
}

func (h handler) PostReviewCompany(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		result models.ReviewCompany
		args   models.ReviewCompanyArgument
		err    error
	)

	// binding body json
	err = json.NewDecoder(r.Body).Decode(&args)
	if err != nil {
		log.Printf("[company] [delivery] [PostReviewCompany] while body binding, err:%+v\n", err)
		response.ResultError(w, response.ErrorBadRequest, err)
		return
	}

	// validate args
	_, err = govalidator.ValidateStruct(args)
	if err != nil {
		log.Printf("[company] [delivery] [PostReviewCompany] while ValidateStruct, err:%+v\n", err)
		response.ResultError(w, response.ErrorValidation, err)
		return
	}

	// only candidate can post review and check authorized
	candidate := utils.GetAuthorization(r.Context())
	if candidate.Role != 2 {
		log.Println("[company] [delivery] [CreateReviewCompany] while ErrorOnlyCandidate")
		response.ResultError(w, response.ErrorOnlyCandidate, nil)
		return
	}

	if candidate.ID != args.CandidateID {
		log.Println("[company] [delivery] [CreateReviewCompany] while ErrorUnauthorized")
		response.ResultError(w, response.ErrorUnauthorized, nil)
		return
	}

	// usecase detail company
	result, err = h.companyUsecase.CreateReviewCompany(r.Context(), args)
	if err != nil {
		errCode, _ := strconv.Atoi(err.Error())
		log.Println("[company] [delivery] [PostReviewCompany] while companyUsecase.CreateReviewCompany")
		response.Result(w, errCode)
		return
	}

	// return data
	response.ResultWithData(w, response.SuccesOk, result)
}

func (h handler) GetReviewCompany(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		result models.ListReviewCompany
		args   models.ListData
		err    error
	)
	// check param id
	companyId := ps.ByName("companyId")
	if companyId == "" {
		log.Println("[company] [delivery] [GetReviewCompany] while get paramater id")
		response.ResultError(w, response.ErrorInvalidParameter, err)
		return
	}

	// define args
	page := utils.Utint(r.URL.Query().Get("page"), 1)                  // default 1
	itemPerPage := utils.Utint(r.URL.Query().Get("item_per_page"), 10) // default 10

	args.ItemPerPage = itemPerPage
	args.Page = page
	args.Limit = itemPerPage

	// usecase list review company
	result, err = h.companyUsecase.GetReviewCompany(r.Context(), companyId, args)
	if err != nil {
		errCode, _ := strconv.Atoi(err.Error())
		log.Println("[company] [delivery] [GetReviewCompany] while companyUsecase.GetReviewCompany")
		response.Result(w, errCode)
		return
	}

	// return data
	response.ResultWithData(w, response.SuccesOk, result)
}
