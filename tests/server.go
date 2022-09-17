package tests

import (
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/dhany007/golang-job-portal/services/delivery/rest"
	"github.com/dhany007/golang-job-portal/services/repository/database"
	"github.com/dhany007/golang-job-portal/services/repository/postgres"
	"github.com/dhany007/golang-job-portal/services/usecase"
)

type Server interface {
	Request(method string, url string, body io.Reader) *http.Response
	RequestWithAuthentication(method string, url string, body io.Reader, token string) *http.Response
	Handler() http.Handler
}

type ServerImpl struct {
	DB *database.DB
}

// RequestWithAuthentication implements Server
func (s ServerImpl) RequestWithAuthentication(method string, url string, body io.Reader, token string) *http.Response {
	request := httptest.NewRequest(method, url, body)

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+token)
	recorder := httptest.NewRecorder()
	s.Handler().ServeHTTP(recorder, request)
	response := recorder.Result()

	return response
}

func NewServer(DB *database.DB) Server {
	return &ServerImpl{DB}
}

func (s ServerImpl) Request(method string, url string, body io.Reader) *http.Response {
	request := httptest.NewRequest(method, url, body)

	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	s.Handler().ServeHTTP(recorder, request)
	response := recorder.Result()

	return response
}

func (s ServerImpl) Handler() http.Handler {
	userRepository := postgres.NewUserRepository(s.DB)
	userUsecase := usecase.NewUserUsecase(userRepository)

	companyRepository := postgres.NewCompanyRepository(s.DB)
	companyUsecase := usecase.NewCompanyUsecase(companyRepository)

	candidateRepo := postgres.NewCandidateRepository(s.DB)
	candidateUsecase := usecase.NewCandidateUsecase(candidateRepo)

	router := rest.NewHandler(userUsecase, companyUsecase, candidateUsecase)

	return router
}
