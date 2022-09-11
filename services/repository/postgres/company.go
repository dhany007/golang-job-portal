package postgres

import (
	"context"
	"log"

	"github.com/dhany007/golang-job-portal/models"
	"github.com/dhany007/golang-job-portal/services"
	"github.com/dhany007/golang-job-portal/services/repository/database"
	"github.com/dhany007/golang-job-portal/services/repository/postgres/queries"
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

	row, err = c.DB.QueryxContext(ctx, queries.QueryListCompanyDresscodes)
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

	row, err = c.DB.QueryxContext(ctx, queries.QueryListCompanyBenefitcodes)
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

	row, err = c.DB.QueryxContext(ctx, queries.QueryListCompanySizecodes)
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

func (c companyRepository) CheckCompanyByEmail(ctx context.Context, email string) (result models.Company, err error) {
	var (
		row *sqlx.Rows
	)

	row, err = c.DB.QueryxContext(ctx, queries.QueryCheckAvailableEmail, email)
	if err != nil {
		log.Printf("[company] [repository] [CheckCompanyByEmail] while queries.QueryCheckCompanyByEmail, err:%+v\n", err)
		return
	}

	defer row.Close()

	for row.Next() {
		err = row.StructScan(&result)
		if err != nil {
			log.Printf("[company] [repository] [CheckCompanyByEmail] while StructScan, err:%+v\n", err)
			return
		}
	}

	return
}

func (c companyRepository) CheckCompanyById(ctx context.Context, id string) (result models.Company, err error) {
	var (
		row *sqlx.Rows
	)

	row, err = c.DB.QueryxContext(ctx, queries.QueryCheckCompanyById, id)
	if err != nil {
		log.Printf("[company] [repository] [CheckCompanyById] while queries.QueryCheckCompanyById, err:%+v\n", err)
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
	// update benefit
	_, err = c.DB.ExecContext(
		ctx,
		queries.QueryUpdateCompany,
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

	// delete benefit existing
	_, err = c.DB.ExecContext(ctx, queries.QueryDeleteBenefitsByCompanyID, args.ID)
	if err != nil {
		log.Printf("[company] [repository] [UpdateCompany] while QueryDeleteBenefitsByCompanyID, err:%+v\n", err)
		return
	}

	// insert benefit company
	for _, v := range args.Benefit {
		_, err = c.DB.ExecContext(ctx, queries.QueryInsertBenefit, args.ID, v.ID)
		if err != nil {
			log.Printf("[company] [repository] [UpdateCompany] while QueryInsertBenefit, err:%+v\n", err)
			return
		}
	}

	return
}

func (c companyRepository) GetListCompanies(ctx context.Context) (result []models.Companies, err error) {
	var (
		row     *sqlx.Rows
		company models.Companies
	)

	row, err = c.DB.QueryxContext(ctx, queries.QueryListCompanies)
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
		queries.QueryDetailCompanyByID,
		companyId,
	)
	if err != nil {
		log.Printf("[company] [repository] [GetDetailCompany] while queries.QueryGetDetailCompanyByID, err:%+v\n", err)
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
	row, err = c.DB.QueryxContext(ctx, queries.QueryListBenefitByCompanyID, companyId)
	if err != nil {
		log.Printf("[company] [repository] [GetDetailCompany] while queries.QueryListBenefitByCompanyID, err:%+v\n", err)
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
	row, err = c.DB.QueryxContext(ctx, queries.QueryRatingCompany, companyId)
	if err != nil {
		log.Printf("[company] [repository] [GetDetailCompany] while queries.QueryRatingCompany, err:%+v\n", err)
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
