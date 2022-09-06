package usecase

import (
	"context"
	"log"

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
