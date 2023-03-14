package database

import (
	"errors"
	"fmt"
	"time"

	"github.com/dhany007/golang-job-portal/services/utils"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

var availableDriver = []string{"postgres", "mysql"}

type DB struct {
	*sqlx.DB
}

func NewConnection(driver, connection string) (*DB, error) {
	if !utils.ArrayContains(availableDriver, driver) {
		return nil, errors.New("driver not found")
	}

	db, err := sqlx.Connect(driver, connection)
	if err != nil {
		return nil, fmt.Errorf("db: failed to connect database: %+v", err)
	}

	db.SetMaxOpenConns(0)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Minute * 15)

	return &DB{db}, nil
}

func NewPostgreConnection(connectionString string) (*DB, error) {
	return NewConnection("postgres", connectionString)
}
