package repo

import (
	"database/sql"

	"github.com/dhany007/golang-job-portal/services/repository/database"
)

type UserRepositoryImpl struct {
	DB *database.DB
}

type User struct {
	ID string `db:"id"`
}

type UserRepository interface {
	GetCompany() (result []User, err error)
	GetUserByEmail(email string) (result User, err error)
	GetCandidate() (result []User, err error)
}

func NewUser(db *database.DB) UserRepository {
	return &UserRepositoryImpl{DB: db}
}

// GetUserByEmail implements UserRepository
func (u UserRepositoryImpl) GetUserByEmail(email string) (result User, err error) {
	var (
		rows *sql.Rows
		id   string
	)
	rows, err = u.DB.Query(queryGetUserByEmail, email)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return
		}
		result = User{ID: id}
	}

	return
}

// GetCompany implements UserRepository
func (u UserRepositoryImpl) GetCompany() (result []User, err error) {
	var (
		rows *sql.Rows
		id   string
	)
	rows, err = u.DB.Query(queryGetUser, "1")
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return
		}
		result = append(result, User{ID: id})
	}

	return
}

// GetCandidate implements UserRepository
func (u UserRepositoryImpl) GetCandidate() (result []User, err error) {
	var (
		rows *sql.Rows
		id   string
	)
	rows, err = u.DB.Query(queryGetUser, "2")
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return
		}
		result = append(result, User{ID: id})
	}

	return
}
