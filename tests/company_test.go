package tests

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/dhany007/golang-job-portal/models"
)

var (
	prefixCompanies = "http://localhost:11010/companies"
	urlDress        = prefixCompanies + "/dress-codes"
	urlBenefit      = prefixCompanies + "/benefit-codes"
	urlSize         = prefixCompanies + "/size-codes"
)

func TestUpdateCompanies(t *testing.T) {
	db, err := InitPostgresTest()
	AssertNoError(t, err)

	server := NewServer(db)

	company1 := `
		{
			"email": "company1@gmail.com",
			"password": "admin123"
		}
	`

	tokenCompany1, err := GetAccessToken(company1)
	AssertNoError(t, err)
	companyId1 := "588b5f49-7cd4-4efe-808a-03818c7a7e9a"
	companyId2 := "5bef8def-f996-4e1d-acac-2c0611240d1d"

	type testCase struct {
		desc      string
		body      models.CompanyArgument
		code      int
		companyId string
	}

	var testCases []testCase

	testCases = append(testCases, testCase{
		desc:      "failed, name required",
		body:      models.CompanyArgument{},
		code:      400,
		companyId: companyId1,
	})

	testCases = append(testCases, testCase{
		desc: "failed, description required",
		body: models.CompanyArgument{
			Name: "pt. manunggal jaya",
		},
		code:      400,
		companyId: companyId1,
	})

	testCases = append(testCases, testCase{
		desc: "failed, address required",
		body: models.CompanyArgument{
			Name:        "pt. manunggal jaya",
			Description: "maju selalu",
		},
		code:      400,
		companyId: companyId1,
	})

	testCases = append(testCases, testCase{
		desc: "failed, phone_number required",
		body: models.CompanyArgument{
			Name:        "pt. manunggal jaya",
			Description: "maju selalu",
			Address:     "jalan sejahtera",
			Website:     "www.sejahtera.com",
		},
		code:      400,
		companyId: companyId1,
	})

	testCases = append(testCases, testCase{
		desc: "failed, dress required",
		body: models.CompanyArgument{
			Name:             "pt. manunggal jaya",
			Description:      "maju selalu",
			Address:          "jalan sejahtera",
			Website:          "www.sejahtera.com",
			PhoneNumber:      "089809870987",
			TelpNumber:       "089809870987",
			ProfilPictureUrl: "www.sejahterra.com/image.png",
		},
		code:      400,
		companyId: companyId1,
	})

	testCases = append(testCases, testCase{
		desc: "failed, size required",
		body: models.CompanyArgument{
			Name:             "pt. manunggal jaya",
			Description:      "maju selalu",
			Address:          "jalan sejahtera",
			Website:          "www.sejahtera.com",
			PhoneNumber:      "089809870987",
			TelpNumber:       "089809870987",
			ProfilPictureUrl: "www.sejahterra.com/image.png",
			Dress:            1,
		},
		code:      400,
		companyId: companyId1,
	})

	testCases = append(testCases, testCase{
		desc: "failed, benefit required",
		body: models.CompanyArgument{
			Name:             "pt. manunggal jaya",
			Description:      "maju selalu",
			Address:          "jalan sejahtera",
			Website:          "www.sejahtera.com",
			PhoneNumber:      "089809870987",
			TelpNumber:       "089809870987",
			ProfilPictureUrl: "www.sejahterra.com/image.png",
			Dress:            1,
			Size:             2,
		},
		code:      400,
		companyId: companyId1,
	})

	testCases = append(testCases, testCase{
		desc: "failed, if user id not same with company id who updated, unauthorized",
		body: models.CompanyArgument{
			Name:             "pt. manunggal jaya",
			Description:      "maju selalu",
			Address:          "jalan sejahtera",
			Website:          "www.sejahtera.com",
			PhoneNumber:      "089809870987",
			TelpNumber:       "089809870987",
			ProfilPictureUrl: "www.sejahterra.com/image.png",
			Dress:            1,
			Size:             2,
			Benefit: []models.CompanySubCode{
				{ID: 1},
				{ID: 2},
			},
		},
		code:      401,
		companyId: companyId2,
	})

	testCases = append(testCases, testCase{
		desc: "success update company",
		body: models.CompanyArgument{
			Name:             "pt. manunggal jaya",
			Description:      "maju selalu",
			Address:          "jalan sejahtera",
			Website:          "www.sejahtera.com",
			PhoneNumber:      "089809870987",
			TelpNumber:       "089809870987",
			ProfilPictureUrl: "www.sejahterra.com/image.png",
			Dress:            1,
			Size:             2,
			Benefit: []models.CompanySubCode{
				{ID: 1},
				{ID: 2},
			},
		},
		code:      200,
		companyId: companyId1,
	})

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			url := prefixCompanies + "/" + tC.companyId
			body, err := json.Marshal(&tC.body)
			AssertNoError(t, err)
			response := server.RequestWithAuthentication(http.MethodPut, url, strings.NewReader(string(body)), tokenCompany1)
			AssertEqualCode(t, response.StatusCode, tC.code)
		})
	}
}

func TestSubDressCodesCompany(t *testing.T) {
	db, err := InitPostgresTest()
	AssertNoError(t, err)

	server := NewServer(db)

	company1 := `
		{
			"email": "company1@gmail.com",
			"password": "admin123"
		}
	`

	tokenCompany1, err := GetAccessToken(company1)
	AssertNoError(t, err)

	response := server.RequestWithAuthentication(http.MethodGet, urlDress, nil, tokenCompany1)
	AssertEqualCode(t, response.StatusCode, 200)
}

func TestSubBenefitCodesCompany(t *testing.T) {
	db, err := InitPostgresTest()
	AssertNoError(t, err)

	server := NewServer(db)

	company1 := `
		{
			"email": "company1@gmail.com",
			"password": "admin123"
		}
	`

	tokenCompany1, err := GetAccessToken(company1)
	AssertNoError(t, err)

	response := server.RequestWithAuthentication(http.MethodGet, urlBenefit, nil, tokenCompany1)
	AssertEqualCode(t, response.StatusCode, 200)
}

func TestSubSizeCodesCompany(t *testing.T) {
	db, err := InitPostgresTest()
	AssertNoError(t, err)

	server := NewServer(db)

	company1 := `
		{
			"email": "company1@gmail.com",
			"password": "admin123"
		}
	`

	tokenCompany1, err := GetAccessToken(company1)
	AssertNoError(t, err)

	response := server.RequestWithAuthentication(http.MethodGet, urlSize, nil, tokenCompany1)
	AssertEqualCode(t, response.StatusCode, 200)
}
