package queries

const (
	QueryInsertCandidate = `
		INSERT INTO candidates(id, email, first_name)
		VALUES($1, $2, $3)
	`

	QueryGetDetailCandidate = `
		SELECT
			id,
			COALESCE(email, '') as email,
			COALESCE(first_name, '') as first_name,
			COALESCE(last_name, '') as last_name,
			COALESCE(phone_number, '') as phone_number,
			COALESCE(telp_number, '') as telp_number,
			COALESCE(address, '') as address,
			COALESCE(profil_picture_url, '') as profil_picture_url,
			created_at,
			modified_at
		FROM
			candidates
		WHERE
			id = $1
	`

	QueryUpdateCandidate = `
		UPDATE candidates
		SET
			email = $1,
			first_name = $2,
			last_name = $3,
			phone_number = $4,
			telp_number = $5,
			address = $6,
			profil_picture_url = $7
		WHERE
			id = $8
	`
)
