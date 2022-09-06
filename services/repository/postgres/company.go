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
