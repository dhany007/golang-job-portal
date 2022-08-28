package usecase

import (
	"fmt"

	"github.com/dhany007/golang-job-portal/services"
)

type testUsecase struct {
}

func NewTestUsecase() services.TestUsecase {
	return testUsecase{}
}

func (t testUsecase) PingTest() {
	fmt.Println("USECASE")
}
