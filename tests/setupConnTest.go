package tests

import (
	"fmt"
	"net/http"

	"github.com/dhany007/golang-job-portal/services/delivery/rest"
	"github.com/dhany007/golang-job-portal/services/repository/database"
	"github.com/dhany007/golang-job-portal/services/repository/postgres"
	"github.com/dhany007/golang-job-portal/services/usecase"
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

func SetupHandlerTest(db *database.DB) http.Handler {
	userRepository := postgres.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)

	companyRepository := postgres.NewCompanyRepository(db)
	companyUsecase := usecase.NewCompanyUsecase(companyRepository)

	candidateRepo := postgres.NewCandidateRepository(db)
	candidateUsecase := usecase.NewCandidateUsecase(candidateRepo)

	router := rest.NewHandler(userUsecase, companyUsecase, candidateUsecase)

	return router
}
