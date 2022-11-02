package rest

import (
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

func (h handler) UpdateCandidate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		result models.Candidate
		args   models.CandidateArgument
		err    error
	)

	// check param id
	candidateId := ps.ByName("candidateId")
	if candidateId == "" {
		log.Println("[candidate] [delivery] [UpdateCandidate] while get params")
		response.Result(w, response.ErrorInvalidParameter)
		return
	}

	// binding body json
	err = json.NewDecoder(r.Body).Decode(&args)
	if err != nil {
		log.Printf("[candidate] [delivery] [UpdateCandidate] while binding body, err:%+v\n", err)
		response.Result(w, response.ErrorBadRequest)
		return
	}

	// validate args
	_, err = govalidator.ValidateStruct(args)
	if err != nil {
		log.Printf("[candidate] [delivery] [UpdateCandidate] while ValidateStruct, err:%+v\n", err)
		response.ResultError(w, response.ErrorValidation, err)
		return
	}

	// check candidateID is same with id user login
	userId := utils.GetAuthorization(r.Context()).ID
	if userId != candidateId {
		log.Println("[candidate] [delivery] [UpdateCandidate] while GetAuthorization")
		response.Result(w, response.ErrorUnauthorized)
		return
	}

	args.ID = candidateId

	// usecase update candidate
	result, err = h.candidateUsecase.UpdateCandidate(r.Context(), args)
	if err != nil {
		errCode, _ := strconv.Atoi(err.Error())
		log.Println("[candidate] [delivery] [UpdateCandidate] while candidateUsecase.UpdateCandidate")
		response.Result(w, errCode)
		return
	}

	response.ResultWithData(w, response.SuccesOk, result)
}

func (h handler) AddExperience(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		result models.CandidateExperience
		args   models.CandidateExperienceArgument
		err    error
	)

	// binding body json
	err = json.NewDecoder(r.Body).Decode(&args)
	if err != nil {
		log.Printf("[candidate] [delivery] [AddExperience] while binding body, err:%+v\n", err)
		response.Result(w, response.ErrorBadRequest)
		return
	}

	// validate args
	_, err = govalidator.ValidateStruct(args)
	if err != nil {
		log.Printf("[candidate] [delivery] [AddExperience] while ValidateStruct, err:%+v\n", err)
		response.ResultError(w, response.ErrorValidation, err)
		return
	}

	// check candidateID is same with id user login
	userId := utils.GetAuthorization(r.Context()).ID
	if userId != args.CandidateID {
		log.Println("[candidate] [delivery] [AddExperience] while GetAuthorization")
		response.Result(w, response.ErrorUnauthorized)
		return
	}

	// usecase add experience
	result, err = h.candidateUsecase.AddExperience(r.Context(), args)
	if err != nil {
		errCode, _ := strconv.Atoi(err.Error())
		log.Println("[candidate] [delivery] [AddExperience] while candidateUsecase.AddExperience")
		response.Result(w, errCode)
		return
	}

	response.ResultWithData(w, response.SuccesOk, result)
}
