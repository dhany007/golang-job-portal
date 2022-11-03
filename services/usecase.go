package services

import (
	"context"

	"github.com/dhany007/golang-job-portal/models"
)

type UserUsecase interface {
	Register(ctx context.Context, args models.User) (result models.User, err error)
	Login(ctx context.Context, args models.UserLoginArgument) (result models.UserLoginResponse, err error)
	RefreshToken(ctx context.Context, token string) (result models.UserLoginResponse, err error)
}

type CompanyUsecase interface {
	GetListDresscode(ctx context.Context) (result []models.CompanySubCode, err error)
	GetListBenefitcode(ctx context.Context) (result []models.CompanySubCode, err error)
	GetListSizecode(ctx context.Context) (result []models.CompanySubCode, err error)
	UpdateCompany(ctx context.Context, args models.CompanyArgument) (result models.Company, err error)
	GetListCompanies(ctx context.Context, args models.ListData) (result models.ListCompanies, err error)
	GetDetailCompany(ctx context.Context, companyID string) (result models.Company, err error)
	CreateReviewCompany(ctx context.Context, args models.ReviewCompanyArgument) (result models.ReviewCompany, err error)
	GetReviewCompany(ctx context.Context, companyID string, args models.ListData) (result models.ListReviewCompany, err error)
}

type CandidateUsecase interface {
	UpdateCandidate(ctx context.Context, args models.CandidateArgument) (result models.Candidate, err error)
	AddExperience(ctx context.Context, args models.CandidateExperienceArgument) (result models.CandidateExperience, err error)
	UpdateExperience(ctx context.Context, args models.CandidateExperience) (result models.CandidateExperience, err error)
}
