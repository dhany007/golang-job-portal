package usecase

import (
	"context"
	"testing"

	"github.com/dhany007/golang-job-portal/models"
	"github.com/dhany007/golang-job-portal/services/repository/postgres/mocks"
	"github.com/dhany007/golang-job-portal/services/utils"
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

		result, err := usecase.Register(context.Background(), tC.input)

		t.Run(tC.desc, func(t *testing.T) {
			if tC.wantError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tC.input.Email, result.Email)
			}
		})
	}
}

func TestUsecase_User_Login(t *testing.T) {
	type testCase struct {
		desc             string
		wantError        bool
		input            models.UserLoginArgument
		onGetUserByEmail func(mock *mocks.MockUserRepository)
	}

	var testCases []testCase

	pass := "admin123"

	testCases = append(testCases, testCase{
		desc:      "failed, user not found",
		wantError: true,
		input: models.UserLoginArgument{
			Email:    "manunggal@company.com",
			Password: pass,
		},
		onGetUserByEmail: func(mock *mocks.MockUserRepository) {
			mock.EXPECT().GetUserByEmail(gomock.Any(), gomock.Any()).Return(
				models.User{},
				nil,
			)
		},
	})

	testCases = append(testCases, testCase{
		desc:      "failed, password not match",
		wantError: true,
		input: models.UserLoginArgument{
			Email:    "manunggal@company.com",
			Password: pass,
		},
		onGetUserByEmail: func(mock *mocks.MockUserRepository) {
			mock.EXPECT().GetUserByEmail(gomock.Any(), gomock.Any()).Return(
				models.User{
					ID:       "xxxx",
					Email:    "manunggal@company.com",
					Role:     1,
					Password: "xxx",
				},
				nil,
			)
		},
	})

	testCases = append(testCases, testCase{
		desc:      "success user login",
		wantError: false,
		input: models.UserLoginArgument{
			Email:    "manunggal@company.com",
			Password: "admin123",
		},
		onGetUserByEmail: func(mock *mocks.MockUserRepository) {
			mock.EXPECT().GetUserByEmail(gomock.Any(), gomock.Any()).Return(
				models.User{
					ID:       "xxxx",
					Email:    "manunggal@company.com",
					Role:     1,
					Password: hashTestPass(pass),
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

		_, err := usecase.Login(context.Background(), tC.input)

		t.Run(tC.desc, func(t *testing.T) {
			if tC.wantError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func hashTestPass(p string) string {
	result, _ := utils.HashPassword(p)
	return result
}
