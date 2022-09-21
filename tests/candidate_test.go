package tests

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/dhany007/golang-job-portal/models"
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
	candidateId1 := "3cc1fefe-bde0-459e-a3af-ded514e6b102"
	candidateId2 := "f05432a3-ce4d-45f9-8603-33378332f736"

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
		code:        400,
		candidateId: candidateId1,
	})

	testCases = append(testCases, testCase{
		desc: "failed, phonenumber required",
		body: models.CandidateArgument{
			FirstName: "kalai",
			LastName:  "saragih",
		},
		code:        400,
		candidateId: candidateId1,
	})

	testCases = append(testCases, testCase{
		desc: "failed, address required",
		body: models.CandidateArgument{
			FirstName:   "kalai",
			LastName:    "saragih",
			PhoneNumber: "082109091010",
		},
		code:        400,
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
		code:        401,
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
		code:        200,
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
