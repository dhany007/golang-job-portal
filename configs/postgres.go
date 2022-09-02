package configs

import (
	"fmt"

	"github.com/dhany007/golang-job-portal/services/repository/database"
	"github.com/dhany007/golang-job-portal/services/utils"
)

func (c *Config) InitPostgres() error {
	var (
		host     = utils.GetEnv("DB_HOST", DB_HOST)
		user     = utils.GetEnv("DB_USER", DB_USER)
		password = utils.GetEnv("DB_PWD", DB_PASS)
		dbname   = utils.GetEnv("DB_NAME", DB_NAME)
		port     = utils.GetEnv("DB_PORT", DB_PORT)
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
