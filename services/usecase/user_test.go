package usecase

import (
	"context"
	"fmt"
	"testing"

	"github.com/dhany007/golang-job-portal/models"
	"github.com/dhany007/golang-job-portal/services/repository/postgres/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUsecase_User_Register(t *testing.T) {
	type testCase struct {
		desc             string
		wantError        bool
		input            models.User
		onGetUserByEmail func(mock *mocks.MockUserRepository)
		onCreateUser     func(mock *mocks.MockUserRepository)
	}

	var testCases []testCase

	testCases = append(testCases, testCase{
		desc: "failed, email exist",
		input: models.User{
			Email:    "manunggal@company.com",
			Password: "admin123",
			Role:     1,
		},
		wantError: true,
		onGetUserByEmail: func(mock *mocks.MockUserRepository) {
			mock.EXPECT().GetUserByEmail(gomock.Any(), gomock.Any()).Return(
				models.User{
					ID:    "xxx",
					Email: "manunggal@company.com",
				},
				nil,
			)
		},
	})

	testCases = append(testCases, testCase{
		desc: "success register user",
		input: models.User{
			Email:    "manunggal@company.com",
			Password: "admin123",
			Role:     1,
		},
		wantError: false,
		onGetUserByEmail: func(mock *mocks.MockUserRepository) {
			mock.EXPECT().GetUserByEmail(gomock.Any(), gomock.Any()).Return(
				models.User{},
				nil,
			)
		},
		onCreateUser: func(mock *mocks.MockUserRepository) {
			mock.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(
				models.User{
					ID:    "xxx",
					Email: "manunggal@company.com",
				},
				nil,
			)
		},
	})

	for _, tC := range testCases {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		repo := mocks.NewMockUserRepository(mockCtrl)
		usecase := userUsecase{repo}

		if tC.onGetUserByEmail != nil {
			tC.onGetUserByEmail(repo)
		}

		if tC.onCreateUser != nil {
			tC.onCreateUser(repo)
		}

		result, serr := usecase.Register(context.Background(), tC.input)
		fmt.Println("result ", result)
		fmt.Println("serr ", serr)
		t.Run(tC.desc, func(t *testing.T) {
			if tC.wantError {
				// want error
				assert.NotNil(t, serr)
			} else {
				// want result
				assert.Nil(t, serr)
				assert.Equal(t, tC.input.Email, result.Email)
			}
		})
	}
}
