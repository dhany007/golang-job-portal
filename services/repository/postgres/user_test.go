package postgres

import (
	"context"
	"database/sql"
	"log"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dhany007/golang-job-portal/services/repository/database"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error %s was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestRepository_User_GetUserByEmail(t *testing.T) {
	db, mock := NewMock()

	var DB *database.DB
	dbx := sqlx.NewDb(db, "jobportal-mock")
	DB = &database.DB{DB: dbx}

	rows := sqlmock.NewRows([]string{
		"id",
		"email",
		"hash_password",
		"is_active",
		"role",
		"created_at",
		"modified_at",
	}).AddRow(
		"41ea02f8-f512-4da8-b8d2-4c8178f29b51",
		"kalai@gmail.com",
		"$2a$08$qd.BS5gNvbJrLC3LG9wLI.5vb4JxDOspKWMhoiigTcRVnKoztyB7O",
		1,
		2,
		time.Now(),
		time.Now(),
	)

	mock.ExpectQuery(`
		SELECT
			(.+)
		FROM
			users
		WHERE
			(.+)
	`).WillReturnRows(rows).RowsWillBeClosed()

	userHandler := NewUserRepository(DB)
	result, errx := userHandler.GetUserByEmail(context.Background(), "kalai@gmail.com")
	assert.Nil(t, errx)
	assert.NoError(t, errx)
	assert.NotNil(t, result)
}
