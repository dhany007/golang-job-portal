package configs

import (
	"github.com/dhany007/golang-job-portal/services/delivery/rest"
	"github.com/dhany007/golang-job-portal/services/repository/postgres"
	"github.com/dhany007/golang-job-portal/services/usecase"
)

func (c *Config) InitService() (err error) {
	userRepository := postgres.NewUserRepository(c.DB)
	userUsecase := usecase.NewUserUsecase(userRepository)

	companyRepository := postgres.NewCompanyRepository(c.DB)
	companyUsecase := usecase.NewCompanyUsecase(companyRepository)

	router := rest.NewHandler(userUsecase, companyUsecase)

	c.Router = router
	return nil
}
