package postgres

import _ "embed"

// file format => {schema}.{table}--{command}.sql

// users query
var (

	//go:embed schemas/users/jobportal.users--check_available_email.sql
	QueryCheckAvailableEmail string

	//go:embed schemas/users/jobportal.users--check_user_by_id.sql
	QueryCheckUsersById string

	//go:embed schemas/users/jobportal.users--get_user_by_email.sql
	QueryGetUserByEmail string

	//go:embed schemas/users/jobportal.users--get_user_by_id.sql
	QueryGetUserById string

	//go:embed schemas/users/jobportal.users--insert_user.sql
	QueryInsertUser string

	//go:embed schemas/users/jobportal.users--update_email_user.sql
	QueryUpdateEmailUser string
)

// companies query
var (
	//go:embed schemas/companies/jobportal.companies--insert_company.sql
	QueryInsertCompany string

	//go:embed schemas/companies/jobportal.companies--list_company_dresscodes.sql
	QueryListCompanyDresscodes string

	//go:embed schemas/companies/jobportal.companies--list_company_benefitcodes.sql
	QueryListCompanyBenefitcodes string

	//go:embed schemas/companies/jobportal.companies--list_benefit_by_company_id.sql
	QueryListBenefitByCompanyID string

	//go:embed schemas/companies/jobportal.companies--delete_benefit_by_company_id.sql
	QueryDeleteBenefitsByCompanyID string

	//go:embed schemas/companies/jobportal.companies--list_company_size_codes.sql
	QueryListCompanySizecodes string

	//go:embed schemas/companies/jobportal.companies--update_company.sql
	QueryUpdateCompany string

	//go:embed schemas/companies/jobportal.companies--insert_benefit.sql
	QueryInsertBenefit string

	//go:embed schemas/companies/jobportal.companies--detail_company_by_id.sql
	QueryDetailCompanyByID string

	//go:embed schemas/companies/jobportal.companies--list_company.sql
	QueryListCompanies string

	//go:embed schemas/companies/jobportal.companies--rating_company.sql
	QueryRatingCompany string

	//go:embed schemas/companies/jobportal.companies--insert_review_company.sql
	QueryInsertReviewCompany string

	//go:embed schemas/companies/jobportal.companies--get_review_company_id.sql
	QueryGetReviewCompanyID string

	//go:embed schemas/companies/jobportal.companies--get_review_company.sql
	QueryGetReviewCompany string

	//go:embed schemas/companies/jobportal.companies--get_count_review_company.sql
	QueryGetCountReviewCompany string

	//go:embed schemas/companies/jobportal.companies--get_count_companies.sql
	QueryGetCountCompanies string
)

// candidates query
var (
	//go:embed schemas/candidates/jobportal.candidates--insert_candidate.sql
	QueryInsertCandidate string

	//go:embed schemas/candidates/jobportal.candidates--get_detail_candidate.sql
	QueryGetDetailCandidate string

	//go:embed schemas/candidates/jobportal.candidates--update_candidate.sql
	QueryUpdateCandidate string

	//go:embed schemas/candidates/jobportal.candidates--add_experience.sql
	QueryAddExperience string

	//go:embed schemas/candidates/jobportal.candidates--update_experience.sql
	QueryUpdateExperience string

	//go:embed schemas/candidates/jobportal.candidates--get_experience_by_id.sql
	QueryGetExperienceById string
)
