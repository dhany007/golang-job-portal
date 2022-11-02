package usecase

import (
	"context"
	"log"

	"github.com/dhany007/golang-job-portal/models"
	"github.com/dhany007/golang-job-portal/services"
	"github.com/dhany007/golang-job-portal/services/utils"
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
		log.Printf("[candidate] [usecase] [UpdateCandidate] while CheckCompanyById, args: %+v\n", args)
		return
	}

	// reinitialize data
	candidate.FirstName = args.FirstName
	candidate.LastName = args.LastName
	candidate.PhoneNumber = args.PhoneNumber
	candidate.TelpNumber = args.TelpNumber
	candidate.Address = args.Address
	candidate.ProfilPictureUrl = args.ProfilPictureUrl

	// repository update candidate
	err = c.repo.UpdateCandidate(ctx, candidate)
	if err != nil {
		log.Printf("[candidate] [usecase] [UpdateCandidate] while repo.UpdateCandidate, args: %+v\n", args)
		return
	}

	// get detail candidate (will update in another ticket)
	result, err = c.repo.GetDetailCandidate(ctx, args.ID)
	if err != nil {
		log.Printf("[candidate] [usecase] [GetDetailCandidate] while CheckCandidateByEmail, id: %+v\n", args.ID)
		return
	}

	return
}

// AddExperience implements services.CandidateUsecase
func (c candidateUsecase) AddExperience(ctx context.Context, args models.CandidateExperienceArgument) (result models.CandidateExperience, err error) {
	var (
		experience models.CandidateExperience
	)

	// reinitialize data
	experience.CandidateID = args.CandidateID
	experience.CompanyName = args.CompanyName
	experience.Title = args.Title
	experience.Description = args.Description
	experience.DateStart = utils.ConvertStringToUtc(args.DateStart)
	experience.DateEnd = utils.ConvertStringToUtc(args.DateEnd)

	// repository add experience
	result, err = c.repo.AddExperience(ctx, experience)
	if err != nil {
		log.Printf("[candidate] [usecase] [AddExperience] while repo.AddExperience, id: %+v\n", args.CandidateID)
		return
	}

	return
}
