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
	CheckCompanyByEmail(ctx context.Context, email string) (result models.Company, err error)
	CheckCompanyById(ctx context.Context, id string) (result models.Company, err error)
	UpdateCompany(ctx context.Context, args models.Company) (result models.Company, err error)
	GetListCompanies(ctx context.Context) (result []models.Companies, err error)
}
