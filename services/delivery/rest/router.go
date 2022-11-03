package rest

import (
	"github.com/dhany007/golang-job-portal/services"
	"github.com/dhany007/golang-job-portal/services/middleware"
	"github.com/julienschmidt/httprouter"
)

type handler struct {
	userUsecase      services.UserUsecase
	companyUsecase   services.CompanyUsecase
	candidateUsecase services.CandidateUsecase
}

func NewHandler(
	userUsecase services.UserUsecase,
	companyUsecase services.CompanyUsecase,
	candidateUsecase services.CandidateUsecase,
) (router *httprouter.Router) {
	router = httprouter.New()

	h := handler{
		userUsecase,
		companyUsecase,
		candidateUsecase,
	}

	// users router
	router.POST("/users/register", h.Register)
	router.POST("/users/login", h.Login)
	router.POST("/users/refresh-token", h.RefreshToken)
	router.POST("/users/logout", h.Logout)

	// below use token for authentication
	// companies router
	router.GET("/companies/detail/:companyId", middleware.Authentication(h.GetDetailCompany))
	router.GET("/companies/reviews/:companyId", middleware.Authentication(h.GetReviewCompany))
	router.GET("/companies/dress-codes", middleware.Authentication(h.GetListDresscode))
	router.GET("/companies/benefit-codes", middleware.Authentication(h.GetListBenefitcode))
	router.GET("/companies/size-codes", middleware.Authentication(h.GetListSizecode))
	router.GET("/companies", middleware.Authentication(h.GetListCompanies))
	router.POST("/companies/reviews", middleware.Authentication(h.PostReviewCompany))
	router.PUT("/companies/:companyId", middleware.Authentication(h.UpdateCompany))

	// candidates router
	router.PUT("/candidates/:candidateId", middleware.Authentication(h.UpdateCandidate))
	router.POST("/candidates/experiences", middleware.Authentication(h.AddExperience))
	router.PUT("/candidates/:candidateId/experiences/:experienceId", middleware.Authentication(h.UpdateExperience))

	return
}
