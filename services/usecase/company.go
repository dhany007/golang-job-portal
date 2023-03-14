package usecase

import (
	"context"
	"log"
	"math"
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

	// check company exists
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

	// reinitialize data
	company = models.Company{
		ID:               args.ID,
		Email:            company.Email,
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

	// repository update company
	err = c.repo.UpdateCompany(ctx, company)
	if err != nil {
		err = response.NewErrork(response.ErrorServerError)
		log.Printf("[company] [usecase] [UpdateCompany] while ErrorServerError, email:%+v\n", company.Email)
		return result, err
	}

	// get detail company
	result, err = c.repo.GetDetailCompany(ctx, args.ID)
	if err != nil {
		err = response.NewErrork(response.ErrorServerError)
		log.Printf("[company] [usecase] [UpdateCompany] while ErrorServerError, email:%+v\n", company.Email)
		return result, err
	}

	return
}

func (c companyUsecase) GetListCompanies(ctx context.Context, args models.ListData) (result models.ListCompanies, err error) {
	var (
		companies []models.Companies
	)

	args.Offset = (args.Page - 1) * args.Limit

	// get list companies
	companies, err = c.repo.GetListCompanies(ctx, args)
	if err != nil {
		log.Printf("[company] [usecase] [GetListCompanies] while repo.GetListCompanies, err:%+v\n", err)
		return
	}

	// check if data not found
	if companies == nil {
		err = response.NewErrork(response.ErrorNotFound)
		log.Println("[company] [usecase] [GetListCompanies] while ErrorNotFound")
		return
	}

	// get total data
	totalData, err := c.repo.GetCountCompanies(ctx)
	if err != nil {
		log.Printf("[company] [usecase] [GetListCompanies] while repo.GetCountCompanies, err:%+v\n", err)
		return
	}

	totalPage := math.Ceil(float64(totalData) / float64(args.Limit))

	result.Page = args.Page
	result.TotalData = totalData
	result.ItemPerPage = args.Limit
	result.TotalPage = int(totalPage)
	result.Companies = companies

	return
}

func (c companyUsecase) GetDetailCompany(ctx context.Context, companyID string) (result models.Company, err error) {
	// repository get detail company
	result, err = c.repo.GetDetailCompany(ctx, companyID)
	if err != nil {
		err = response.NewErrork(response.ErrorServerError)
		log.Printf("[company] [usecase] [GetDetailCompany] while ErrorServerError, companyID:%+v\n", companyID)
		return result, err
	}

	// check if data not found
	if result.Email == "" {
		err = response.NewErrork(response.ErrorNotFound)
		log.Println("[company] [usecase] [GetDetailCompany] while ErrorNotFound")
		return
	}

	return
}

func (c companyUsecase) CreateReviewCompany(ctx context.Context, args models.ReviewCompanyArgument) (result models.ReviewCompany, err error) {
	var (
		company models.Company
		review  models.ReviewCompany
	)

	// check availability company
	company, err = c.repo.CheckCompanyById(ctx, args.CompanyID)
	if err != nil {
		log.Println("[company] [usecase] [CreateReviewCompany] while CheckCompanyById")
		return
	}

	if company.ID == "" {
		err = response.NewErrork(response.ErrorNotFound)
		log.Println("[company] [usecase] [CreateReviewCompany] while ErrorNotFound")
		return
	}

	// reinitialize data
	review = models.ReviewCompany{
		CompanyID:   args.CompanyID,
		CandidateID: args.CandidateID,
		Rating:      args.Rating,
		Review:      args.Review,
	}

	// repository post review company
	result, err = c.repo.CreateReviewCompany(ctx, review)
	if err != nil {
		log.Printf("[company] [usecase] [CreateReviewCompany] while ErrorServerError, args:%+v\n", args)
		return result, err
	}

	return
}

func (c companyUsecase) GetReviewCompany(ctx context.Context, companyID string, args models.ListData) (result models.ListReviewCompany, err error) {
	var (
		reviews []models.ReviewCompany
	)

	args.Offset = (args.Page - 1) * args.Limit

	// get list review
	reviews, err = c.repo.GetReviewCompany(ctx, companyID, args)
	if err != nil {
		err = response.NewErrork(response.ErrorServerError)
		log.Printf("[company] [usecase] [GetReviewCompany] while ErrorServerError, companyID:%+v\n", companyID)
		return result, err
	}

	// check if data not found
	if reviews == nil {
		err = response.NewErrork(response.ErrorNotFound)
		log.Println("[company] [usecase] [GetReviewCompany] while ErrorNotFound")
		return
	}

	// check total data
	totalData, err := c.repo.GetCountReviewCompany(ctx, companyID)
	if err != nil {
		log.Printf("[company] [usecase] [GetReviewCompany] while GetCountReviewCompany, companyID:%+v\n", companyID)
		return result, err
	}

	totalPage := math.Ceil(float64(totalData) / float64(args.Limit))

	result.Page = args.Page
	result.TotalData = totalData
	result.ItemPerPage = args.Limit
	result.TotalPage = int(totalPage)
	result.Reviews = reviews

	return
}
