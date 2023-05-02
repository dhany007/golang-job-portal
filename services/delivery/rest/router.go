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

	// companies router
	// public router
	router.GET("/companies/detail/:companyId", h.GetDetailCompany)
	router.GET("/companies/reviews/:companyId", h.GetReviewCompany)
	router.GET("/companies", h.GetListCompanies)
	router.GET("/companies/dress-codes", h.GetListDresscode)
	router.GET("/companies/benefit-codes", h.GetListBenefitcode)
	router.GET("/companies/size-codes", h.GetListSizecode)

	// below use token for authentication
	router.POST("/companies/reviews", middleware.Authentication(h.PostReviewCompany))
	router.PUT("/companies/:companyId", middleware.Authentication(h.UpdateCompany))

	// candidates router
	router.PUT("/candidates/:candidateId", middleware.Authentication(h.UpdateCandidate))
	router.POST("/candidates/experiences", middleware.Authentication(h.AddExperience))
	router.PUT("/candidates/:candidateId/experiences/:experienceId", middleware.Authentication(h.UpdateExperience))

	return
}
