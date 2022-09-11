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
	GetListCompanies(ctx context.Context) (result []models.Companies, err error)
	GetDetailCompany(ctx context.Context, companyID string) (result models.Company, err error)
}
