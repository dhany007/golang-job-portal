package usecase

import (
	"context"
	"testing"

	"github.com/dhany007/golang-job-portal/models"
	"github.com/dhany007/golang-job-portal/services/repository/postgres/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUsecase_Company_GetListDresscode(t *testing.T) {
	type testCase struct {
		desc               string
		wantError          bool
		onGetListDresscode func(mock *mocks.MockCompanyRepository)
	}

	var testCases []testCase

	testCases = append(testCases, testCase{
		desc:      "failed, dresscode not found",
		wantError: true,
		onGetListDresscode: func(mock *mocks.MockCompanyRepository) {
			mock.EXPECT().GetListDresscode(gomock.Any()).Return(
				[]models.CompanySubCode{},
				nil,
			)
		},
	})

	testCases = append(testCases, testCase{
		desc:      "success get dresscode",
		wantError: false,
		onGetListDresscode: func(mock *mocks.MockCompanyRepository) {
			mock.EXPECT().GetListDresscode(gomock.Any()).Return(
				[]models.CompanySubCode{
					{
						ID:    1,
						Value: "Casual",
					},
				},
				nil,
			)
		},
	})

	for _, tC := range testCases {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		repo := mocks.NewMockCompanyRepository(mockCtrl)
		usecase := companyUsecase{repo}

		if tC.onGetListDresscode != nil {
			tC.onGetListDresscode(repo)
		}

		result, serr := usecase.GetListDresscode(context.Background())

		t.Run(tC.desc, func(t *testing.T) {
			if tC.wantError {
				assert.Equal(t, len(result), 0)
			} else {
				assert.Nil(t, serr)
			}
		})
	}
}

func TestUsecase_Company_GetListSizecode(t *testing.T) {
	type testCase struct {
		desc              string
		wantError         bool
		onGetListSizecode func(mock *mocks.MockCompanyRepository)
	}

	var testCases []testCase

	testCases = append(testCases, testCase{
		desc:      "failed, sizecodes not found",
		wantError: true,
		onGetListSizecode: func(mock *mocks.MockCompanyRepository) {
			mock.EXPECT().GetListSizecode(gomock.Any()).Return(
				[]models.CompanySubCode{},
				nil,
			)
		},
	})

	testCases = append(testCases, testCase{
		desc:      "success get sizecodes",
		wantError: false,
		onGetListSizecode: func(mock *mocks.MockCompanyRepository) {
			mock.EXPECT().GetListSizecode(gomock.Any()).Return(
				[]models.CompanySubCode{
					{
						ID:    1,
						Value: "Self-employed",
					},
				},
				nil,
			)
		},
	})

	for _, tC := range testCases {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		repo := mocks.NewMockCompanyRepository(mockCtrl)
		usecase := companyUsecase{repo}

		if tC.onGetListSizecode != nil {
			tC.onGetListSizecode(repo)
		}

		result, serr := usecase.GetListSizecode(context.Background())

		t.Run(tC.desc, func(t *testing.T) {
			if tC.wantError {
				assert.Equal(t, len(result), 0)
			} else {
				assert.Nil(t, serr)
			}
		})
	}
}

func TestUsecase_Company_GetListBenefitcode(t *testing.T) {
	type testCase struct {
		desc                 string
		wantError            bool
		onGetListBenefitcode func(mock *mocks.MockCompanyRepository)
	}

	var testCases []testCase

	testCases = append(testCases, testCase{
		desc:      "failed, benefit not found",
		wantError: true,
		onGetListBenefitcode: func(mock *mocks.MockCompanyRepository) {
			mock.EXPECT().GetListBenefitcode(gomock.Any()).Return(
				[]models.CompanySubCode{},
				nil,
			)
		},
	})

	testCases = append(testCases, testCase{
		desc:      "success get benefit",
		wantError: false,
		onGetListBenefitcode: func(mock *mocks.MockCompanyRepository) {
			mock.EXPECT().GetListBenefitcode(gomock.Any()).Return(
				[]models.CompanySubCode{
					{
						ID:    1,
						Value: "Benefits that are required by law",
					},
				},
				nil,
			)
		},
	})

	for _, tC := range testCases {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		repo := mocks.NewMockCompanyRepository(mockCtrl)
		usecase := companyUsecase{repo}

		if tC.onGetListBenefitcode != nil {
			tC.onGetListBenefitcode(repo)
		}

		result, serr := usecase.GetListBenefitcode(context.Background())

		t.Run(tC.desc, func(t *testing.T) {
			if tC.wantError {
				assert.Equal(t, len(result), 0)
			} else {
				assert.Nil(t, serr)
			}
		})
	}
}
