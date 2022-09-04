package rest

import (
	"github.com/dhany007/golang-job-portal/services"
	"github.com/julienschmidt/httprouter"
)

type handler struct {
	userUsecase services.UserUsecase
}

func NewHandler(
	userUsecase services.UserUsecase,
) (router *httprouter.Router) {
	router = httprouter.New()

	h := handler{
		userUsecase,
	}

	// user router
	router.POST("/users/register", h.Register)
	router.POST("/users/login", h.Login)
	router.POST("/users/refresh-token", h.RefreshToken)
	router.POST("/users/logout", h.Logout)

	return
}
