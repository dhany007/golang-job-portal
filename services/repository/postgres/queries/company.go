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

	QueryListBenefitByCompanyID = `
		SELECT
			cb.id , cbc.value
		FROM
			company_benefits cb
		JOIN
			company_benefits_codes cbc
		ON
			cb.benefit_id = cbc.id
		WHERE
			cb.company_id = $1
	`

	QueryDeleteBenefitsByCompanyID = `
		DELETE FROM company_benefits
		WHERE company_id = $1
	`

	QueryListCompanySizecodes = `
		SELECT
			COALESCE(id, 0) as id,
			COALESCE(value, '') as value
		FROM
			company_size_codes
	`

	QueryCheckCompanyById = `
		SELECT id, email
		FROM
			users
		WHERE
			id = $1
	`

	QueryCheckAvailableEmail = `
		SELECT id, email
		FROM
			users
		WHERE
			email = $1
	`

	QueryUpdateCompany = `
		UPDATE companies
		SET
			email = $1,
			name = $2,
			description = $3,
			address = $4,
			website = $5,
			phone_number = $6,
			telp_number = $7,
			profil_picture_url = $8,
			dresscode_code = $9,
			size_code = $10,
			modified_at = $11
		WHERE
			id = $12
	`

	QueryInsertBenefit = `
		INSERT INTO company_benefits (company_id, benefit_id)
		VALUES ($1, $2)
	`

	QueryDetailCompanyByID = `
		SELECT
			COALESCE(c.email, '') AS email,
			COALESCE(c.name, '') AS name,
			COALESCE(c.description, '') AS description,
			COALESCE(c.address, '') AS address,
			COALESCE(c.website, '') AS website,
			COALESCE(c.phone_number, '') AS phone_number,
			COALESCE(c.telp_number, '') AS telp_number,
			COALESCE(c.profil_picture_url, '') AS profil_picture_url,
			COALESCE(cdc.value, '') AS dress,
			COALESCE(czc.value, '') AS size,
			c.created_at,
			c.modified_at
		FROM
			companies c
		LEFT JOIN
			company_dresscode_codes cdc
		ON
			c.dresscode_code = cdc.id
		LEFT JOIN
			company_size_codes czc
		ON
			c.size_code = czc.id
		WHERE
			c.id = $1
	`

	QueryListCompanies = `
		SELECT
			c.id,
			COALESCE(c.name, '') AS name,
			COALESCE(AVG(cr.rating), 0) AS rating,
			COALESCE(COUNT(cr.rating), 0) AS count_review
		FROM
			companies c
		LEFT JOIN
			company_reviews cr
		ON
			c.id = cr.company_id
		GROUP BY
			c.id
		ORDER BY
			rating DESC
		LIMIT $1
		OFFSET $2
	`

	QueryRatingCompany = `
		SELECT
			COALESCE(AVG(cr.rating), 0) AS rating,
			COALESCE(COUNT(cr.rating), 0) AS count_review
		FROM
			companies c
		LEFT JOIN
			company_reviews cr
		ON
			c.id = cr.company_id
		WHERE
			c.id = $1
		GROUP BY
			c.id
	`

	QueryInsertReviewCompany = `
		INSERT INTO company_reviews(company_id, candidate_id, rating, review)
		VALUES($1, $2, $3, $4)
	`

	QueryGetReviewCompanyID = `
		SELECT
			COALESCE(id, 0) AS id,
			company_id,
			candidate_id,
			COALESCE(rating, 0) AS rating,
			COALESCE(review, '') AS review
		FROM
			company_reviews
		WHERE
			company_id = $1
			AND candidate_id = $2
	`

	QueryGetReviewCompany = `
		SELECT
			COALESCE(id, 0) AS id,
			company_id,
			candidate_id,
			COALESCE(rating, 0) AS rating,
			COALESCE(review, '') AS review
		FROM
			company_reviews
		WHERE
			company_id = $1
		ORDER BY
			created_at DESC
		LIMIT $2
		OFFSET $3
	`

	QueryGetCountReviewCompany = `
		SELECT
			COUNT(*)
		FROM
			company_reviews
		WHERE
			company_id = $1
	`

	QueryGetCountCompanies = `
		SELECT
			COUNT(*)
		FROM
			companies
	`
)
