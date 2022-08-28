package configs

import (
	"fmt"
	"os"

	"github.com/dhany007/golang-job-portal/services/repository/database"
)

func (c *Config) InitPostgres() error {
	var (
		host     = os.Getenv("DB_HOST")
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PWD")
		dbname   = os.Getenv("DB_NAME")
		port     = os.Getenv("DB_PORT")
	)

	dsn := fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%s sslmode=disable",
		host, user, dbname, password, port,
	)

	db, err := database.NewPostgreConnection(dsn)

	if err != nil {
		fmt.Println("err", err)
		return fmt.Errorf("failed connect to database, %s", dsn)
	}

	c.DB = db

	return nil
}
