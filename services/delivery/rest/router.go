package rest

import (
	"github.com/dhany007/golang-job-portal/services"
	"github.com/dhany007/golang-job-portal/services/middleware"
	"github.com/julienschmidt/httprouter"
)

type handler struct {
	userUsecase    services.UserUsecase
	companyUsecase services.CompanyUsecase
}

func NewHandler(
	userUsecase services.UserUsecase,
	companyUsecase services.CompanyUsecase,
) (router *httprouter.Router) {
	router = httprouter.New()

	h := handler{
		userUsecase,
		companyUsecase,
	}

	// users router
	router.POST("/users/register", h.Register)
	router.POST("/users/login", h.Login)
	router.POST("/users/refresh-token", h.RefreshToken)
	router.POST("/users/logout", h.Logout)

	// below use token for authentication
	// companies router
	router.GET("/companies/dress-codes", middleware.Authentication(h.GetListDresscode))
	router.GET("/companies/benefit-codes", middleware.Authentication(h.GetListBenefitcode))
	router.GET("/companies/size-codes", middleware.Authentication(h.GetListSizecode))
	router.PUT("/companies/:companyId", middleware.Authentication(h.UpdateCompany))

	return
}
