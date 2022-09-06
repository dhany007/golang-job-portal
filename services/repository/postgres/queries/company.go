package queries

const (
	QueryInsertCompany = `
		INSERT INTO companies(id, email, name)
		VALUES($1, $2, $3)
	`

	QueryListCompanyDresscodes = `
		SELECT
			COALESCE(id, 0) as id,
			COALESCE(value, '') as value
		FROM
			company_dresscode_codes
	`

	QueryListCompanyBenefitcodes = `
		SELECT
			COALESCE(id, 0) as id,
			COALESCE(value, '') as value
		FROM
			company_benefits_codes
	`

	QueryListCompanySizecodes = `
		SELECT
			COALESCE(id, 0) as id,
			COALESCE(value, '') as value
		FROM
			company_size_codes
	`
)
