package repo

const (
	queryGetUser = `
		SELECT id FROM users WHERE role = $1;
	`

	queryGetUserByEmail = `
		SELECT id FROM users WHERE email = $1;
	`
)
