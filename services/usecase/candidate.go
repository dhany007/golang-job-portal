package usecase

import (
	"context"
	"log"

	"github.com/dhany007/golang-job-portal/models"
	"github.com/dhany007/golang-job-portal/models/response"
	"github.com/dhany007/golang-job-portal/services"
)

type candidateUsecase struct {
	repo services.CandidateRepository
}

func NewCandidateUsecase(repo services.CandidateRepository) services.CandidateUsecase {
	return &candidateUsecase{repo}
}

func (c candidateUsecase) UpdateCandidate(ctx context.Context, args models.CandidateArgument) (result models.Candidate, err error) {
	var (
		candidate models.Candidate
	)

	// check candidate by id given
	candidate, err = c.repo.CheckCandidateById(ctx, args.ID)
	if err != nil {
		log.Printf("[company] [usecase] [UpdateCompany] while CheckCompanyById, args: %+v\n", args)
		return
	}

	// check email availability
	candidate, err = c.repo.CheckCandidateByEmail(ctx, args.Email)
	if err != nil {
		err = response.NewErrork(response.ErrorRegisEmail)
		log.Printf("[company] [usecase] [UpdateCompany] while CheckCandidateByEmail, args: %+v\n", args)
		return
	}

	// reinitialize data
	candidate.Email = args.Email
	candidate.FirstName = args.FirstName
	candidate.LastName = args.LastName
	candidate.PhoneNumber = args.PhoneNumber
	candidate.TelpNumber = args.TelpNumber
	candidate.Address = args.Address
	candidate.ProfilPictureUrl = args.ProfilPictureUrl

	// repository update candidate
	err = c.repo.UpdateCandidate(ctx, candidate)
	if err != nil {
		log.Printf("[company] [usecase] [UpdateCompany] while repo.UpdateCandidate, args: %+v\n", args)
		return
	}

	// get detail candidate (will update in another ticket)
	result, err = c.repo.GetDetailCandidate(ctx, args.ID)
	if err != nil {
		log.Printf("[company] [usecase] [GetDetailCandidate] while CheckCandidateByEmail, id: %+v\n", args.ID)
		return
	}

	return
}
