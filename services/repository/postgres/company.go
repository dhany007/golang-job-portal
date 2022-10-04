package postgres

import (
	"context"
	"log"
	"strings"

	"github.com/dhany007/golang-job-portal/models"
	"github.com/dhany007/golang-job-portal/models/response"
	"github.com/dhany007/golang-job-portal/services"
	"github.com/dhany007/golang-job-portal/services/repository/database"
	"github.com/jmoiron/sqlx"
)

type companyRepository struct {
	DB *database.DB
}

func NewCompanyRepository(db *database.DB) services.CompanyRepository {
	return &companyRepository{db}
}

func (c companyRepository) GetListDresscode(ctx context.Context) (result []models.CompanySubCode, err error) {
	var (
		row   *sqlx.Rows
		dress models.CompanySubCode
	)

	row, err = c.DB.QueryxContext(ctx, QueryListCompanyDresscodes)
	if err != nil {
		log.Printf("[company] [repository] [GetListDresscode] while QueryListCompanyDresscode, err:%+v\n", err)
		return
	}

	defer row.Close()

	for row.Next() {
		err = row.StructScan(&dress)
		if err != nil {
			log.Printf("[company] [repository] [GetListDresscode] while StructScan, err:%+v\n", err)
			return
		}

		result = append(result, dress)
	}

	return
}

func (c companyRepository) GetListBenefitcode(ctx context.Context) (result []models.CompanySubCode, err error) {
	var (
		row     *sqlx.Rows
		benefit models.CompanySubCode
	)

	row, err = c.DB.QueryxContext(ctx, QueryListCompanyBenefitcodes)
	if err != nil {
		log.Printf("[company] [repository] [GetListBenefitcode] while QueryListCompanyBenefitcode, err:%+v\n", err)
		return
	}

	defer row.Close()

	for row.Next() {
		err = row.StructScan(&benefit)
		if err != nil {
			log.Printf("[company] [repository] [GetListBenefitcode] while StructScan, err:%+v\n", err)
			return
		}

		result = append(result, benefit)
	}

	return
}

func (c companyRepository) GetListSizecode(ctx context.Context) (result []models.CompanySubCode, err error) {
	var (
		row  *sqlx.Rows
		size models.CompanySubCode
	)

	row, err = c.DB.QueryxContext(ctx, QueryListCompanySizecodes)
	if err != nil {
		log.Printf("[company] [repository] [GetListSizecode] while QueryListCompanySizecodes, err:%+v\n", err)
		return
	}

	defer row.Close()

	for row.Next() {
		err = row.StructScan(&size)
		if err != nil {
			log.Printf("[company] [repository] [GetListSizecode] while StructScan, err:%+v\n", err)
			return
		}

		result = append(result, size)
	}

	return
}

func (c companyRepository) CheckCompanyByEmail(ctx context.Context, email string) (result []models.Company, err error) {
	var (
		row     *sqlx.Rows
		company models.Company
	)

	row, err = c.DB.QueryxContext(ctx, QueryCheckAvailableEmail, email)
	if err != nil {
		log.Printf("[company] [repository] [CheckCompanyByEmail] while QueryCheckCompanyByEmail, err:%+v\n", err)
		return
	}

	defer row.Close()

	for row.Next() {
		err = row.StructScan(&company)
		if err != nil {
			log.Printf("[company] [repository] [CheckCompanyByEmail] while StructScan, err:%+v\n", err)
			return
		}

		result = append(result, company)
	}

	return
}

func (c companyRepository) CheckCompanyById(ctx context.Context, id string) (result models.Company, err error) {
	var (
		row *sqlx.Rows
	)

	row, err = c.DB.QueryxContext(ctx, QueryCheckUsersById, id)
	if err != nil {
		log.Printf("[company] [repository] [CheckCompanyById] while QueryCheckUsersById, err:%+v\n", err)
		return
	}

	defer row.Close()

	for row.Next() {
		err = row.StructScan(&result)
		if err != nil {
			log.Printf("[company] [repository] [CheckCompanyById] while StructScan, err:%+v\n", err)
			return
		}
	}

	return
}

func (c companyRepository) UpdateCompany(ctx context.Context, args models.Company) (err error) {
	// update company
	_, err = c.DB.ExecContext(
		ctx,
		QueryUpdateCompany,
		args.Email,
		args.Name,
		args.Description,
		args.Address,
		args.Website,
		args.PhoneNumber,
		args.TelpNumber,
		args.ProfilPictureUrl,
		args.Dress,
		args.Size,
		args.ModifiedAt,
		args.ID,
	)

	if err != nil {
		log.Printf("[company] [repository] [UpdateCompany] while QueryUpdateCompany, err:%+v\n", err)
		return
	}

	// update user
	_, err = c.DB.ExecContext(
		ctx,
		QueryUpdateEmailUser,
		args.Email,
		args.ID,
	)

	if err != nil {
		log.Printf("[company] [repository] [UpdateCompany] while QueryUpdateEmailUser, err:%+v\n", err)
		return
	}

	// delete benefit existing
	_, err = c.DB.ExecContext(ctx, QueryDeleteBenefitsByCompanyID, args.ID)
	if err != nil {
		log.Printf("[company] [repository] [UpdateCompany] while QueryDeleteBenefitsByCompanyID, err:%+v\n", err)
		return
	}

	// insert benefit company
	for _, v := range args.Benefit {
		_, err = c.DB.ExecContext(ctx, QueryInsertBenefit, args.ID, v.ID)
		if err != nil {
			log.Printf("[company] [repository] [UpdateCompany] while QueryInsertBenefit, err:%+v\n", err)
			return
		}
	}

	return
}

func (c companyRepository) GetListCompanies(ctx context.Context, args models.ListData) (result []models.Companies, err error) {
	var (
		row     *sqlx.Rows
		company models.Companies
	)

	row, err = c.DB.QueryxContext(ctx, QueryListCompanies, args.Limit, args.Offset)
	if err != nil {
		log.Printf("[company] [repository] [GetListCompanies] while QueryListCompanies, err:%+v\n", err)
		return
	}

	defer row.Close()

	for row.Next() {
		err = row.StructScan(&company)
		if err != nil {
			log.Printf("[company] [repository] [GetListCompanies] while StructScan, err:%+v\n", err)
			return
		}

		result = append(result, company)
	}

	return
}

func (c companyRepository) GetDetailCompany(ctx context.Context, companyId string) (result models.Company, err error) {
	var (
		row      *sqlx.Rows
		benefits []models.CompanySubCode
		benefit  models.CompanySubCode
		reviews  models.CompanyRating
	)

	// get detail company
	row, err = c.DB.QueryxContext(
		ctx,
		QueryDetailCompanyByID,
		companyId,
	)
	if err != nil {
		log.Printf("[company] [repository] [GetDetailCompany] while QueryGetDetailCompanyByID, err:%+v\n", err)
		return
	}

	defer row.Close()

	for row.Next() {
		err = row.StructScan(&result)
		if err != nil {
			log.Printf("[company] [repository] [GetDetailCompany] while StructScan, err:%+v\n", err)
			return
		}
	}

	// get benefit
	row, err = c.DB.QueryxContext(ctx, QueryListBenefitByCompanyID, companyId)
	if err != nil {
		log.Printf("[company] [repository] [GetDetailCompany] while QueryListBenefitByCompanyID, err:%+v\n", err)
		return
	}

	for row.Next() {
		err = row.StructScan(&benefit)
		if err != nil {
			log.Printf("[company] [repository] [GetDetailCompany] while StructScan, err:%+v\n", err)
			return
		}
		benefits = append(benefits, benefit)
	}

	result.Benefit = benefits

	// get review
	row, err = c.DB.QueryxContext(ctx, QueryRatingCompany, companyId)
	if err != nil {
		log.Printf("[company] [repository] [GetDetailCompany] while QueryRatingCompany, err:%+v\n", err)
		return
	}

	for row.Next() {
		err = row.StructScan(&reviews)
		if err != nil {
			log.Printf("[company] [repository] [GetDetailCompany] while StructScan, err:%+v\n", err)
			return
		}
	}
	result.Review = reviews

	return
}

func (c companyRepository) CreateReviewCompany(ctx context.Context, args models.ReviewCompany) (result models.ReviewCompany, err error) {
	var (
		row *sqlx.Rows
	)

	// insert review
	_, err = c.DB.ExecContext(
		ctx,
		QueryInsertReviewCompany,
		args.CompanyID,
		args.CandidateID,
		args.Rating,
		args.Review,
	)

	if err != nil && strings.Contains(err.Error(), "uq_candidate_company") {
		err = response.NewErrork(response.ErrorReviewFound)
		log.Printf("[company] [repository] [CreateReviewCompany] while uq_candidate_company, err:%+v\n", err)
		return
	}

	if err != nil {
		log.Printf("[company] [repository] [CreateReviewCompany] while QueryInsertReviewCompany, err:%+v\n", err)
		return
	}

	// get detail company review
	row, err = c.DB.QueryxContext(
		ctx,
		QueryGetReviewCompanyID,
		args.CompanyID,
		args.CandidateID,
	)
	if err != nil {
		log.Printf("[company] [repository] [CreateReviewCompany] while QueryGetReviewCompanyID, err:%+v\n", err)
		return
	}

	defer row.Close()

	for row.Next() {
		err = row.StructScan(&result)
		if err != nil {
			log.Printf("[company] [repository] [CreateReviewCompany] while StructScan, err:%+v\n", err)
			return
		}
	}

	return
}

func (c companyRepository) GetReviewCompany(ctx context.Context, companyID string, args models.ListData) (result []models.ReviewCompany, err error) {
	var (
		row    *sqlx.Rows
		review models.ReviewCompany
	)

	row, err = c.DB.QueryxContext(
		ctx,
		QueryGetReviewCompany,
		companyID,
		args.Limit,
		args.Offset,
	)

	if err != nil {
		log.Printf("[company] [repository] [GetReviewCompany] while QueryGetReviewCompany, err:%+v\n", err)
		return
	}

	defer row.Close()

	for row.Next() {
		err = row.StructScan(&review)
		if err != nil {
			log.Printf("[company] [repository] [GetReviewCompany] while StructScan, err:%+v\n", err)
			return
		}
		review.CompanyID = ""

		result = append(result, review)
	}

	return
}

func (c companyRepository) GetCountReviewCompany(ctx context.Context, companyID string) (total int, err error) {
	row, err := c.DB.QueryxContext(ctx, QueryGetCountReviewCompany, companyID)
	if err != nil {
		log.Printf("[company] [repository] [GetCountReviewCompany] while QueryGetCountReviewCompany, err:%+v\n", err)
		return
	}

	defer row.Close()

	for row.Next() {
		err = row.Scan(&total)
		if err != nil {
			log.Printf("[company] [repository] [GetCountReviewCompany] while StructScan, err:%+v\n", err)
			return
		}
	}

	return
}

func (c companyRepository) GetCountCompanies(ctx context.Context) (total int, err error) {
	row, err := c.DB.QueryxContext(ctx, QueryGetCountCompanies)
	if err != nil {
		log.Printf("[company] [repository] [GetCountCompanies] while QueryGetCountCompanies, err:%+v\n", err)
		return
	}

	defer row.Close()

	for row.Next() {
		err = row.Scan(&total)
		if err != nil {
			log.Printf("[company] [repository] [GetCountCompanies] while StructScan, err:%+v\n", err)
			return
		}
	}

	return
}
