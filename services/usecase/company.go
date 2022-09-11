package usecase

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/dhany007/golang-job-portal/models"
	"github.com/dhany007/golang-job-portal/models/response"
	"github.com/dhany007/golang-job-portal/services"
)

type companyUsecase struct {
	repo services.CompanyRepository
}

func NewCompanyUsecase(repo services.CompanyRepository) services.CompanyUsecase {
	return &companyUsecase{repo}
}

func (c companyUsecase) GetListDresscode(ctx context.Context) (result []models.CompanySubCode, err error) {
	// get dresscodes
	result, err = c.repo.GetListDresscode(ctx)
	if err != nil {
		log.Printf("[company] [usecase] [GetListDresscode] while repo.GetListDresscode, err:%+v\n", err)
		return
	}

	// check if data not found
	if result == nil {
		err = response.NewErrork(response.ErrorNotFound)
		log.Println("[company] [usecase] [GetListDresscode] while ErrorNotFound")
		return
	}

	return
}

func (c companyUsecase) GetListBenefitcode(ctx context.Context) (result []models.CompanySubCode, err error) {
	// get dresscodes
	result, err = c.repo.GetListBenefitcode(ctx)
	if err != nil {
		log.Printf("[company] [usecase] [GetListBenefitcode] while repo.GetListBenefitcode, err:%+v\n", err)
		return
	}

	// check if data not found
	if result == nil {
		err = response.NewErrork(response.ErrorNotFound)
		log.Println("[company] [usecase] [GetListBenefitcode] while ErrorNotFound")
		return
	}

	return
}

func (c companyUsecase) GetListSizecode(ctx context.Context) (result []models.CompanySubCode, err error) {
	// get dresscodes
	result, err = c.repo.GetListSizecode(ctx)
	if err != nil {
		log.Printf("[company] [usecase] [GetListSizecode] while repo.GetListSizecode, err:%+v\n", err)
		return
	}

	// check if data not found
	if result == nil {
		err = response.NewErrork(response.ErrorNotFound)
		log.Println("[company] [usecase] [GetListSizecode] while ErrorNotFound")
		return
	}

	return
}

func (c companyUsecase) UpdateCompany(ctx context.Context, args models.CompanyArgument) (result models.Company, err error) {
	var (
		company models.Company
	)

	// TODO: check company exists
	company, err = c.repo.CheckCompanyById(ctx, args.ID)
	if err != nil {
		log.Println("[company] [usecase] [UpdateCompany] while CheckCompanyById")
		return
	}
	if company.ID == "" {
		err = response.NewErrork(response.ErrorNotFound)
		log.Println("[company] [usecase] [UpdateCompany] while ErrorNotFound")
		return
	}

	// TODO: check email availability if change
	company, _ = c.repo.CheckCompanyByEmail(ctx, args.Email)
	if company.ID != "" {
		err = response.NewErrork(response.ErrorRegisEmail)
		log.Println("[company] [usecase] [UpdateCompany] while CheckCompanyByEmail")
		return
	}

	// reinitialize data
	company = models.Company{
		ID:               args.ID,
		Email:            args.Email,
		Name:             args.Name,
		Description:      args.Description,
		Address:          args.Address,
		Website:          args.Website,
		PhoneNumber:      args.PhoneNumber,
		TelpNumber:       args.TelpNumber,
		ProfilPictureUrl: args.ProfilPictureUrl,
		Dress:            strconv.Itoa(args.Dress),
		Size:             strconv.Itoa(args.Size),
		Benefit:          args.Benefit,
		ModifiedAt:       time.Now(),
	}

	// TODO: repository update company
	result, err = c.repo.UpdateCompany(ctx, company)
	if err != nil {
		err = response.NewErrork(response.ErrorServerError)
		log.Printf("[company] [usecase] [UpdateCompany] while ErrorServerError, email:%+v\n", args.Email)
		return result, err
	}

	return
}

func (c companyUsecase) GetListCompanies(ctx context.Context) (result []models.Companies, err error) {
	// get companies
	result, err = c.repo.GetListCompanies(ctx)
	if err != nil {
		log.Printf("[company] [usecase] [GetListCompanies] while repo.GetListCompanies, err:%+v\n", err)
		return
	}

	// check if data not found
	if result == nil {
		err = response.NewErrork(response.ErrorNotFound)
		log.Println("[company] [usecase] [GetListCompanies] while ErrorNotFound")
		return
	}

	return
}
