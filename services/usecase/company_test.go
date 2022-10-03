package usecase

import (
	"context"
	"errors"
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

		result, err := usecase.GetListDresscode(context.Background())

		t.Run(tC.desc, func(t *testing.T) {
			if tC.wantError {
				assert.Equal(t, len(result), 0)
			} else {
				assert.Nil(t, err)
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

		result, err := usecase.GetListSizecode(context.Background())

		t.Run(tC.desc, func(t *testing.T) {
			if tC.wantError {
				assert.Equal(t, len(result), 0)
			} else {
				assert.Nil(t, err)
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

		result, err := usecase.GetListBenefitcode(context.Background())

		t.Run(tC.desc, func(t *testing.T) {
			if tC.wantError {
				assert.Equal(t, len(result), 0)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestUsecase_Company_UpdateCompany(t *testing.T) {
	type testCase struct {
		desc               string
		wantError          bool
		input              models.CompanyArgument
		onCheckCompanyById func(mock *mocks.MockCompanyRepository)
		onUpdateCompany    func(mock *mocks.MockCompanyRepository)
		onGetDetailCompany func(mock *mocks.MockCompanyRepository)
	}

	var testCases []testCase

	testCases = append(testCases, testCase{
		desc:      "failed, company not found",
		wantError: true,
		input: models.CompanyArgument{
			ID:          "xxx",
			Name:        "new company",
			Description: "new description",
		},
		onCheckCompanyById: func(mock *mocks.MockCompanyRepository) {
			mock.EXPECT().CheckCompanyById(gomock.Any(), gomock.Any()).Return(
				models.Company{},
				nil,
			)
		},
	})

	testCases = append(testCases, testCase{
		desc:      "failed, something wrong",
		wantError: true,
		input: models.CompanyArgument{
			ID:          "xxx",
			Name:        "new company",
			Description: "new description",
		},
		onCheckCompanyById: func(mock *mocks.MockCompanyRepository) {
			mock.EXPECT().CheckCompanyById(gomock.Any(), gomock.Any()).Return(
				models.Company{
					ID:          "xxx",
					Name:        "company",
					Description: "description",
				},
				nil,
			)
		},
		onUpdateCompany: func(mock *mocks.MockCompanyRepository) {
			mock.EXPECT().UpdateCompany(gomock.Any(), gomock.Any()).Return(
				errors.New("something wrong"),
			)
		},
	})

	testCases = append(testCases, testCase{
		desc:      "success update company",
		wantError: false,
		input: models.CompanyArgument{
			ID:          "xxx",
			Name:        "new company",
			Description: "new description",
		},
		onCheckCompanyById: func(mock *mocks.MockCompanyRepository) {
			mock.EXPECT().CheckCompanyById(gomock.Any(), gomock.Any()).Return(
				models.Company{
					ID:          "xxx",
					Name:        "company",
					Description: "description",
				},
				nil,
			)
		},
		onUpdateCompany: func(mock *mocks.MockCompanyRepository) {
			mock.EXPECT().UpdateCompany(gomock.Any(), gomock.Any()).Return(
				nil,
			)
		},
		onGetDetailCompany: func(mock *mocks.MockCompanyRepository) {
			mock.EXPECT().GetDetailCompany(gomock.Any(), gomock.Any()).Return(
				models.Company{
					ID:          "xxx",
					Name:        "new company",
					Description: "new description",
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

		if tC.onCheckCompanyById != nil {
			tC.onCheckCompanyById(repo)
		}
		if tC.onUpdateCompany != nil {
			tC.onUpdateCompany(repo)
		}
		if tC.onGetDetailCompany != nil {
			tC.onGetDetailCompany(repo)
		}

		result, err := usecase.UpdateCompany(context.Background(), tC.input)

		t.Run(tC.desc, func(t *testing.T) {
			if tC.wantError {
				assert.NotNil(t, err)
			} else {
				assert.Equal(t, tC.input.ID, result.ID)
				assert.Nil(t, err)
				assert.NotNil(t, result)
			}
		})
	}
}

func TestUsecase_Company_GetListCompany(t *testing.T) {
	type testCase struct {
		desc                string
		wantError           bool
		input               models.ListData
		onGetListCompany    func(mock *mocks.MockCompanyRepository)
		onGetCountCompanies func(mock *mocks.MockCompanyRepository)
	}

	var testCases []testCase

	testCases = append(testCases, testCase{
		desc:      "failed, company not found",
		wantError: true,
		input:     models.ListData{},
		onGetListCompany: func(mock *mocks.MockCompanyRepository) {
			mock.EXPECT().GetListCompanies(gomock.Any(), gomock.Any()).Return(
				nil,
				nil,
			)
		},
	})

	testCases = append(testCases, testCase{
		desc:      "success get list companies",
		wantError: false,
		input:     models.ListData{},
		onGetListCompany: func(mock *mocks.MockCompanyRepository) {
			mock.EXPECT().GetListCompanies(gomock.Any(), gomock.Any()).Return(
				[]models.Companies{
					{
						ID:          "xxx",
						Name:        "company1",
						Rating:      5,
						CountReview: 2,
					},
					{
						ID:          "yyy",
						Name:        "company2",
						Rating:      4,
						CountReview: 2,
					},
				},
				nil,
			)
		},
		onGetCountCompanies: func(mock *mocks.MockCompanyRepository) {
			mock.EXPECT().GetCountCompanies(gomock.Any()).Return(2, nil)
		},
	})

	for _, tC := range testCases {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		repo := mocks.NewMockCompanyRepository(mockCtrl)
		usecase := companyUsecase{repo}

		if tC.onGetListCompany != nil {
			tC.onGetListCompany(repo)
		}
		if tC.onGetCountCompanies != nil {
			tC.onGetCountCompanies(repo)
		}

		companies, err := usecase.GetListCompanies(context.Background(), tC.input)
		t.Run(tC.desc, func(t *testing.T) {
			if tC.wantError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.NotNil(t, companies)
			}
		})
	}
}

func TestUsecase_Company_GetDetailCompany(t *testing.T) {
	type testCase struct {
		desc               string
		wantError          bool
		companyId          string
		onGetDetailCompany func(mock *mocks.MockCompanyRepository)
	}

	var testCases []testCase

	testCases = append(testCases, testCase{
		desc:      "failed, company not found",
		wantError: true,
		companyId: "xxx",
		onGetDetailCompany: func(mock *mocks.MockCompanyRepository) {
			mock.EXPECT().GetDetailCompany(gomock.Any(), gomock.Any()).Return(
				models.Company{},
				nil,
			)
		},
	})

	testCases = append(testCases, testCase{
		desc:      "success get detail company",
		wantError: false,
		companyId: "xxx",
		onGetDetailCompany: func(mock *mocks.MockCompanyRepository) {
			mock.EXPECT().GetDetailCompany(gomock.Any(), gomock.Any()).Return(
				models.Company{
					ID:          "xxx",
					Email:       "company@gmail.com",
					Name:        "company name",
					Description: "description company",
				},
				nil,
			)
		},
	})

	for _, tC := range testCases {

		mockCtrl := gomock.NewController(t)
		mockCtrl.Finish()

		repo := mocks.NewMockCompanyRepository(mockCtrl)
		usecase := companyUsecase{repo}

		if tC.onGetDetailCompany != nil {
			tC.onGetDetailCompany(repo)
		}

		company, err := usecase.GetDetailCompany(context.Background(), tC.companyId)

		t.Run(tC.desc, func(t *testing.T) {
			if tC.wantError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.NotNil(t, company)
				assert.Equal(t, company.ID, tC.companyId)
			}
		})
	}
}
