package services

import (
	"context"

	"github.com/dhany007/golang-job-portal/models"
)

type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (result models.User, err error)
	GetUserById(ctx context.Context, id string) (result models.User, err error)
	CreateUser(ctx context.Context, args models.User) (result models.User, err error)
}

type CompanyRepository interface {
	GetListDresscode(ctx context.Context) (result []models.CompanySubCode, err error)
	GetListBenefitcode(ctx context.Context) (result []models.CompanySubCode, err error)
	GetListSizecode(ctx context.Context) (result []models.CompanySubCode, err error)
	CheckCompanyById(ctx context.Context, id string) (result models.Company, err error)
	UpdateCompany(ctx context.Context, args models.Company) (err error)
	GetListCompanies(ctx context.Context, args models.ListData) (result []models.Companies, err error)
	GetCountCompanies(ctx context.Context) (result int, err error)
	GetDetailCompany(ctx context.Context, id string) (result models.Company, err error)
	CreateReviewCompany(ctx context.Context, args models.ReviewCompany) (result models.ReviewCompany, err error)
	GetReviewCompany(ctx context.Context, companyID string, args models.ListData) (result []models.ReviewCompany, err error)
	GetCountReviewCompany(ctx context.Context, companyID string) (result int, err error)
}

type CandidateRepository interface {
	UpdateCandidate(ctx context.Context, args models.Candidate) (err error)
	CheckCandidateByEmail(ctx context.Context, email string) (result []models.Candidate, err error)
	CheckCandidateById(ctx context.Context, id string) (result models.Candidate, err error)
	GetDetailCandidate(ctx context.Context, id string) (result models.Candidate, err error)
	AddExperience(ctx context.Context, args models.CandidateExperience) (result models.CandidateExperience, err error)
	GetExperienceById(ctx context.Context, experienceId int) (result models.CandidateExperience, err error)
	UpdateExperience(ctx context.Context, args models.CandidateExperience) (err error)
}
