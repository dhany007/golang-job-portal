package tests

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/dhany007/golang-job-portal/models"
	"github.com/dhany007/golang-job-portal/tests/repo"
)

var (
	prefixCandidate = "http://localhost:11010/candidates"
)

func TestUpdateCandidate(t *testing.T) {
	db, err := InitPostgresTest()
	AssertNoError(t, err)

	server := NewServer(db)

	candidate1 := `
		{
			"email": "candidate1@gmail.com",
			"password": "admin123"
		}
	`

	tokenCandidate1, err := GetAccessToken(candidate1)
	AssertNoError(t, err)

	user := repo.NewUser(db)
	userCandidate1, err := user.GetUserByEmail("candidate1@gmail.com")
	AssertNoError(t, err)
	candidateId1 := userCandidate1.ID

	userCandidate2, err := user.GetUserByEmail("candidate2@gmail.com")
	AssertNoError(t, err)
	candidateId2 := userCandidate2.ID

	type testCase struct {
		desc        string
		body        models.CandidateArgument
		code        int
		candidateId string
	}

	var testCases []testCase

	testCases = append(testCases, testCase{
		desc:        "failed, while binding json",
		body:        models.CandidateArgument{},
		code:        http.StatusBadRequest,
		candidateId: candidateId1,
	})

	testCases = append(testCases, testCase{
		desc: "failed, phonenumber required",
		body: models.CandidateArgument{
			FirstName: "kalai",
			LastName:  "saragih",
		},
		code:        http.StatusBadRequest,
		candidateId: candidateId1,
	})

	testCases = append(testCases, testCase{
		desc: "failed, address required",
		body: models.CandidateArgument{
			FirstName:   "kalai",
			LastName:    "saragih",
			PhoneNumber: "082109091010",
		},
		code:        http.StatusBadRequest,
		candidateId: candidateId1,
	})

	testCases = append(testCases, testCase{
		desc: "failed, if user id not same with candidate id who updated, unauthorized",
		body: models.CandidateArgument{
			FirstName:        "kalai",
			LastName:         "saragih",
			PhoneNumber:      "082109091010",
			Address:          "simanabun",
			TelpNumber:       "082109091010",
			ProfilPictureUrl: "profil.jpg",
		},
		code:        http.StatusUnauthorized,
		candidateId: candidateId2,
	})

	testCases = append(testCases, testCase{
		desc: "success update company",
		body: models.CandidateArgument{
			FirstName:        "kalai",
			LastName:         "saragih",
			PhoneNumber:      "082109091010",
			Address:          "simanabun",
			TelpNumber:       "082109091010",
			ProfilPictureUrl: "profil.jpg",
		},
		code:        http.StatusOK,
		candidateId: candidateId1,
	})

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			url := prefixCandidate + "/" + tC.candidateId
			body, err := json.Marshal(&tC.body)
			AssertNoError(t, err)
			response := server.RequestWithAuthentication(http.MethodPut, url, strings.NewReader(string(body)), tokenCandidate1)
			AssertEqualCode(t, response.StatusCode, tC.code)
		})
	}
}
