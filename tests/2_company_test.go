package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/dhany007/golang-job-portal/models"
	"github.com/dhany007/golang-job-portal/tests/repo"
)

var (
	prefixCompanies = "http://localhost:11010/companies"
	urlDress        = prefixCompanies + "/dress-codes"
	urlBenefit      = prefixCompanies + "/benefit-codes"
	urlSize         = prefixCompanies + "/size-codes"
	urlPostReview   = prefixCompanies + "/reviews"
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

	user := repo.NewUser(db)
	userCompany1, err := user.GetUserByEmail("company1@gmail.com")
	AssertNoError(t, err)
	companyId1 := userCompany1.ID

	userCompany2, err := user.GetUserByEmail("company2@gmail.com")
	AssertNoError(t, err)
	companyId2 := userCompany2.ID

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
		code:      http.StatusBadRequest,
		companyId: companyId1,
	})

	testCases = append(testCases, testCase{
		desc: "failed, description required",
		body: models.CompanyArgument{
			Name: "pt. manunggal jaya",
		},
		code:      http.StatusBadRequest,
		companyId: companyId1,
	})

	testCases = append(testCases, testCase{
		desc: "failed, address required",
		body: models.CompanyArgument{
			Name:        "pt. manunggal jaya",
			Description: "maju selalu",
		},
		code:      http.StatusBadRequest,
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
		code:      http.StatusBadRequest,
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
		code:      http.StatusBadRequest,
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
		code:      http.StatusBadRequest,
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
		code:      http.StatusBadRequest,
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
		code:      http.StatusUnauthorized,
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
		code:      http.StatusOK,
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
	AssertEqualCode(t, response.StatusCode, http.StatusOK)
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
	AssertEqualCode(t, response.StatusCode, http.StatusOK)
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
	AssertEqualCode(t, response.StatusCode, http.StatusOK)
}

func TestPostReviewCompany(t *testing.T) {
	db, err := InitPostgresTest()
	AssertNoError(t, err)

	server := NewServer(db)

	candidate1 := `
		{
			"email": "candidate1@gmail.com",
			"password": "admin123"
		}
	`

	candidate2 := `
		{
			"email": "candidate2@gmail.com",
			"password": "admin123"
		}
	`

	tokenCandidate1, err := GetAccessToken(candidate1)
	AssertNoError(t, err)

	tokenCandidate2, err := GetAccessToken(candidate2)
	AssertNoError(t, err)

	user := repo.NewUser(db)
	userCompany, err := user.GetCompany()
	AssertNoError(t, err)

	companyId1 := userCompany[0].ID
	companyId2 := userCompany[1].ID

	userCandidate, err := user.GetCandidate()
	AssertNoError(t, err)

	candidateId1 := userCandidate[0].ID
	candidateId2 := userCandidate[1].ID

	db.Exec("TRUNCATE company_reviews;")

	type testCase struct {
		desc  string
		body  models.ReviewCompanyArgument
		code  int
		token string
	}

	var testCases []testCase

	testCases = append(testCases, testCase{
		desc:  "failed, token invalid",
		body:  models.ReviewCompanyArgument{},
		code:  http.StatusBadRequest,
		token: "",
	})

	testCases = append(testCases, testCase{
		desc:  "failed, while binding body",
		body:  models.ReviewCompanyArgument{},
		code:  http.StatusBadRequest,
		token: tokenCandidate1,
	})

	testCases = append(testCases, testCase{
		desc: "failed, companyid required",
		body: models.ReviewCompanyArgument{
			CandidateID: candidateId1,
			Rating:      5,
			Review:      "Good",
		},
		code:  http.StatusBadRequest,
		token: tokenCandidate1,
	})

	testCases = append(testCases, testCase{
		desc: "failed, candidateid required",
		body: models.ReviewCompanyArgument{
			CompanyID: companyId1,
			Rating:    5,
			Review:    "Good",
		},
		code:  http.StatusBadRequest,
		token: tokenCandidate1,
	})

	testCases = append(testCases, testCase{
		desc: "failed, rating required",
		body: models.ReviewCompanyArgument{
			CompanyID:   companyId1,
			CandidateID: candidateId1,
			Review:      "Good",
		},
		code:  http.StatusBadRequest,
		token: tokenCandidate1,
	})

	testCases = append(testCases, testCase{
		desc: "failed, rating must range(1|5)",
		body: models.ReviewCompanyArgument{
			CompanyID:   companyId1,
			CandidateID: candidateId1,
			Rating:      0,
			Review:      "Good",
		},
		code:  http.StatusBadRequest,
		token: tokenCandidate1,
	})

	testCases = append(testCases, testCase{
		desc: "failed, rating must range(1|5)",
		body: models.ReviewCompanyArgument{
			CompanyID:   companyId1,
			CandidateID: candidateId1,
			Rating:      6,
			Review:      "Good",
		},
		code:  http.StatusBadRequest,
		token: tokenCandidate1,
	})

	testCases = append(testCases, testCase{
		desc: "failed, review required",
		body: models.ReviewCompanyArgument{
			CompanyID:   companyId1,
			CandidateID: candidateId1,
			Rating:      5,
		},
		code:  http.StatusBadRequest,
		token: tokenCandidate1,
	})

	testCases = append(testCases, testCase{
		desc: "failed, user id login as candidate must be same with candidate id",
		body: models.ReviewCompanyArgument{
			CompanyID:   companyId1,
			CandidateID: candidateId1,
			Rating:      5,
			Review:      "Good",
		},
		code:  http.StatusUnauthorized,
		token: tokenCandidate2,
	})

	testCases = append(testCases, testCase{
		desc: "failed, user id login as candidate must be same with candidate id",
		body: models.ReviewCompanyArgument{
			CompanyID:   companyId2,
			CandidateID: candidateId2,
			Rating:      5,
			Review:      "Good",
		},
		code:  http.StatusUnauthorized,
		token: tokenCandidate1,
	})

	testCases = append(testCases, testCase{
		desc: "success candidate1 post review company1",
		body: models.ReviewCompanyArgument{
			CompanyID:   companyId1,
			CandidateID: candidateId1,
			Rating:      5,
			Review:      "Good",
		},
		code:  http.StatusOK,
		token: tokenCandidate1,
	})

	testCases = append(testCases, testCase{
		desc: "success candidate1 post review company2",
		body: models.ReviewCompanyArgument{
			CompanyID:   companyId2,
			CandidateID: candidateId1,
			Rating:      5,
			Review:      "Good",
		},
		code:  http.StatusOK,
		token: tokenCandidate1,
	})

	testCases = append(testCases, testCase{
		desc: "failed, 1 candidate only can post 1 review",
		body: models.ReviewCompanyArgument{
			CompanyID:   companyId1,
			CandidateID: candidateId1,
			Rating:      5,
			Review:      "Good",
		},
		code:  http.StatusBadRequest,
		token: tokenCandidate1,
	})

	testCases = append(testCases, testCase{
		desc: "failed, 1 candidate only can post 1 review",
		body: models.ReviewCompanyArgument{
			CompanyID:   companyId1,
			CandidateID: candidateId1,
			Rating:      5,
			Review:      "Good",
		},
		code:  http.StatusBadRequest,
		token: tokenCandidate1,
	})

	testCases = append(testCases, testCase{
		desc: "success candidate2 post review company1",
		body: models.ReviewCompanyArgument{
			CompanyID:   companyId1,
			CandidateID: candidateId2,
			Rating:      5,
			Review:      "Good",
		},
		code:  http.StatusOK,
		token: tokenCandidate2,
	})

	testCases = append(testCases, testCase{
		desc: "success candidate2 post review company2",
		body: models.ReviewCompanyArgument{
			CompanyID:   companyId2,
			CandidateID: candidateId2,
			Rating:      5,
			Review:      "Good",
		},
		code:  http.StatusOK,
		token: tokenCandidate2,
	})

	testCases = append(testCases, testCase{
		desc: "failed, 1 candidate only can post 1 review",
		body: models.ReviewCompanyArgument{
			CompanyID:   companyId1,
			CandidateID: candidateId2,
			Rating:      5,
			Review:      "Good",
		},
		code:  http.StatusBadRequest,
		token: tokenCandidate2,
	})

	testCases = append(testCases, testCase{
		desc: "failed, 1 candidate only can post 1 review",
		body: models.ReviewCompanyArgument{
			CompanyID:   companyId2,
			CandidateID: candidateId2,
			Rating:      5,
			Review:      "Good",
		},
		code:  http.StatusBadRequest,
		token: tokenCandidate2,
	})

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			body, err := json.Marshal(&tC.body)
			AssertNoError(t, err)
			response := server.RequestWithAuthentication(http.MethodPost, urlPostReview, strings.NewReader(string(body)), tC.token)
			AssertEqualCode(t, response.StatusCode, tC.code)
		})
	}
}

func TestGetReviewCompany(t *testing.T) {
	db, err := InitPostgresTest()
	AssertNoError(t, err)

	server := NewServer(db)

	company1 := `
		{
			"email": "company1@gmail.com",
			"password": "admin123"
		}
	`

	company2 := `
		{
			"email": "company2@gmail.com",
			"password": "admin123"
		}
	`

	tokenCompany1, err := GetAccessToken(company1)
	AssertNoError(t, err)

	tokenCompany2, err := GetAccessToken(company2)
	AssertNoError(t, err)

	user := repo.NewUser(db)
	userCompany1, err := user.GetUserByEmail("company1@gmail.com")
	AssertNoError(t, err)
	companyId1 := userCompany1.ID

	userCompany2, err := user.GetUserByEmail("company2@gmail.com")
	AssertNoError(t, err)
	companyId2 := userCompany2.ID

	type testCase struct {
		desc      string
		companyId string
		code      int
		token     string
	}

	var testCases []testCase

	testCases = append(testCases, testCase{
		desc:      "success get review company1",
		companyId: companyId1,
		code:      http.StatusOK,
		token:     tokenCompany1,
	})

	testCases = append(testCases, testCase{
		desc:      "success get review company2",
		companyId: companyId2,
		code:      http.StatusOK,
		token:     tokenCompany2,
	})

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			url := fmt.Sprintf("%s/%s", urlPostReview, tC.companyId)
			response := server.RequestWithAuthentication(http.MethodGet, url, nil, tC.token)
			AssertEqualCode(t, response.StatusCode, tC.code)
		})
	}
}
