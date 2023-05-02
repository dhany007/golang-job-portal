package configs

import (
	"fmt"

	"github.com/dhany007/golang-job-portal/models"
	"github.com/dhany007/golang-job-portal/services/repository/database"
	"github.com/dhany007/golang-job-portal/services/utils"
)

func (c *Config) InitPostgres() error {
	var (
		host     = utils.GetEnv("DB_HOST", models.DB_HOST)
		user     = utils.GetEnv("DB_USER", models.DB_USER)
		password = utils.GetEnv("DB_PWD", models.DB_PASS)
		dbname   = utils.GetEnv("DB_NAME", models.DB_NAME)
		port     = utils.GetEnv("DB_PORT", models.DB_PORT)
	)

	dsn := fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%s sslmode=disable",
		host, user, dbname, password, port,
	)

	db, err := database.NewPostgreConnection(dsn)

	if err != nil {
		return fmt.Errorf("failed connect to database, %s", dsn)
	}

	c.DB = db

	return nil
}
