package tests

import (
	"fmt"

	"github.com/dhany007/golang-job-portal/services/repository/database"
	"github.com/dhany007/golang-job-portal/services/utils"
)

func InitPostgresTest() (*database.DB, error) {
	var (
		host     = utils.GetEnv("DB_HOST_TEST", "localhost")
		user     = utils.GetEnv("DB_USER_TEST", "postgres")
		password = utils.GetEnv("DB_PWD_TEST", "")
		dbname   = utils.GetEnv("DB_NAME_TEST", "jobportaltest")
		port     = utils.GetEnv("DB_PORT_TEST", "5432")
	)

	dsn := fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%s sslmode=disable",
		host, user, dbname, password, port,
	)

	db, err := database.NewPostgreConnection(dsn)

	if err != nil {
		return nil, err
	}

	return db, nil
}
