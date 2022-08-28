package rest

import (
	"github.com/dhany007/golang-job-portal/services"
	"github.com/julienschmidt/httprouter"
)

type handler struct {
	testUsecase services.TestUsecase
}

func NewHandler(testUsecase services.TestUsecase) (router *httprouter.Router) {
	router = httprouter.New()

	h := handler{
		testUsecase: testUsecase,
	}

	router.GET("/ping", h.PingTest)

	return
}
