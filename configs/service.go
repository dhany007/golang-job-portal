package configs

import (
	"github.com/dhany007/golang-job-portal/services/delivery/rest"
	"github.com/dhany007/golang-job-portal/services/usecase"
)

func (c *Config) InitService() (err error) {
	testUsecase := usecase.NewTestUsecase()

	router := rest.NewHandler(testUsecase)

	c.Router = router
	return nil
}
