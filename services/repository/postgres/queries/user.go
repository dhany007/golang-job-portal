package queries

const (
	QueryGetUserByEmail = `
		SELECT
			id,
			email,
			COALESCE(hash_password, '') AS hash_password,
			COALESCE(is_active, 1) AS is_active,
			COALESCE(role, 2) AS role,
			created_at,
			modified_at
		FROM
			users
		WHERE
			email=$1
	`

	QueryGetUserById = `
		SELECT
			id,
			email,
			COALESCE(hash_password, '') AS hash_password,
			COALESCE(is_active, 1) AS is_active,
			COALESCE(role, 2) AS role,
			created_at,
			modified_at
		FROM
			users
		WHERE
			id=$1
	`

	QueryInsertUser = `
		INSERT INTO users(id, email, hash_password, role)
		VALUES($1, $2, $3, $4)
	`

	QueryUpdateEmailUser = `
		UPDATE users
		SET email = $1
		WHERE id = $2
	`

	QueryCheckUsersById = `
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
)
