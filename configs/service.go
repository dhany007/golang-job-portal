package configs

import (
	"github.com/dhany007/golang-job-portal/services/delivery/rest"
	"github.com/dhany007/golang-job-portal/services/repository/postgres"
	"github.com/dhany007/golang-job-portal/services/repository/redis"
	"github.com/dhany007/golang-job-portal/services/usecase"
)

func (c *Config) InitService() (err error) {
	cacheRepo := redis.NewCacheRepository(c.RedisClient)

	userRepository := postgres.NewUserRepository(c.DB)
	userUsecase := usecase.NewUserUsecase(userRepository)

	companyRepository := postgres.NewCompanyRepository(c.DB, cacheRepo)
	companyUsecase := usecase.NewCompanyUsecase(companyRepository)

	candidateRepo := postgres.NewCandidateRepository(c.DB)
	candidateUsecase := usecase.NewCandidateUsecase(candidateRepo)

	router := rest.NewHandler(userUsecase, companyUsecase, candidateUsecase)

	c.Router = router
	return nil
}
