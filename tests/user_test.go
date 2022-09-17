package tests

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/dhany007/golang-job-portal/models"
)

var (
	prefixUsers = "http://localhost:11010/users"
	urlRegister = prefixUsers + "/register"
	urlLogin    = prefixUsers + "/login"
	urlLogout   = prefixUsers + "/logout"
)

func TestRegisterUser(t *testing.T) {
	db, err := InitPostgresTest()
	AssertNoError(t, err)

	server := NewServer(db)

	db.Exec("TRUNCATE users;")
	db.Exec("TRUNCATE companies;")
	db.Exec("TRUNCATE candidates;")

	type testCase struct {
		desc string
		body models.UserRegisterArguments
		code int
	}

	var testCases []testCase

	testCases = append(testCases, testCase{
		desc: "failed register, email not valid",
		body: models.UserRegisterArguments{
			Email:    "manunggal",
			Password: "admin123",
			Role:     1,
		},
		code: 400,
	})

	testCases = append(testCases, testCase{
		desc: "failed register, email required",
		body: models.UserRegisterArguments{
			Password: "admin123",
			Role:     1,
		},
		code: 400,
	})

	testCases = append(testCases, testCase{
		desc: "failed register, password required",
		body: models.UserRegisterArguments{
			Email: "manunggal",
			Role:  1,
		},
		code: 400,
	})

	testCases = append(testCases, testCase{
		desc: "failed register, length password range(8|32)",
		body: models.UserRegisterArguments{
			Email:    "manunggal@gmail.com",
			Password: "admin",
			Role:     1,
		},
		code: 400,
	})

	testCases = append(testCases, testCase{
		desc: "failed register, length password range(8|32)",
		body: models.UserRegisterArguments{
			Email:    "manunggal@gmail.com",
			Password: "adminadminadminadminadminadminadmin",
			Role:     1,
		},
		code: 400,
	})

	testCases = append(testCases, testCase{
		desc: "success register as company1",
		body: models.UserRegisterArguments{
			Email:    "company1@gmail.com",
			Password: "admin123",
			Role:     1,
		},
		code: 200,
	})

	testCases = append(testCases, testCase{
		desc: "success register as company2",
		body: models.UserRegisterArguments{
			Email:    "company2@gmail.com",
			Password: "admin123",
			Role:     1,
		},
		code: 200,
	})

	testCases = append(testCases, testCase{
		desc: "failed register, email exists",
		body: models.UserRegisterArguments{
			Email:    "company2@gmail.com",
			Password: "admin123",
			Role:     1,
		},
		code: 400,
	})

	testCases = append(testCases, testCase{
		desc: "success register as candidate1",
		body: models.UserRegisterArguments{
			Email:    "candidate1@gmail.com",
			Password: "admin123",
			Role:     2,
		},
		code: 200,
	})

	testCases = append(testCases, testCase{
		desc: "success register as candidate2",
		body: models.UserRegisterArguments{
			Email:    "candidate2@gmail.com",
			Password: "admin123",
			Role:     2,
		},
		code: 200,
	})

	testCases = append(testCases, testCase{
		desc: "failed register, email exists",
		body: models.UserRegisterArguments{
			Email:    "candidate2@gmail.com",
			Password: "admin123",
			Role:     2,
		},
		code: 400,
	})

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			body, err := json.Marshal(&tC.body)
			AssertNoError(t, err)

			response := server.Request(http.MethodPost, urlRegister, strings.NewReader(string(body)))
			AssertEqualCode(t, response.StatusCode, tC.code)
		})
	}
}

func TestLoginUser(t *testing.T) {
	db, err := InitPostgresTest()
	AssertNoError(t, err)

	server := NewServer(db)

	type testCase struct {
		desc string
		body models.UserLoginArgument
		code int
	}

	var testCases []testCase

	testCases = append(testCases, testCase{
		desc: "failed register, email not valid",
		body: models.UserLoginArgument{
			Email:    "company",
			Password: "admin123",
		},
		code: 400,
	})

	testCases = append(testCases, testCase{
		desc: "failed register, email required",
		body: models.UserLoginArgument{
			Password: "admin123",
		},
		code: 400,
	})

	testCases = append(testCases, testCase{
		desc: "failed register, password required",
		body: models.UserLoginArgument{
			Email: "company@gmail.com",
		},
		code: 400,
	})

	testCases = append(testCases, testCase{
		desc: "failed register, users not found",
		body: models.UserLoginArgument{
			Email:    "company@gmail.com",
			Password: "admin123",
		},
		code: 404,
	})

	testCases = append(testCases, testCase{
		desc: "failed register, password not match",
		body: models.UserLoginArgument{
			Email:    "company1@gmail.com",
			Password: "admin1234",
		},
		code: 400,
	})

	testCases = append(testCases, testCase{
		desc: "success login user",
		body: models.UserLoginArgument{
			Email:    "company1@gmail.com",
			Password: "admin123",
		},
		code: 200,
	})

	testCases = append(testCases, testCase{
		desc: "success login user",
		body: models.UserLoginArgument{
			Email:    "company2@gmail.com",
			Password: "admin123",
		},
		code: 200,
	})

	testCases = append(testCases, testCase{
		desc: "success login user",
		body: models.UserLoginArgument{
			Email:    "candidate1@gmail.com",
			Password: "admin123",
		},
		code: 200,
	})

	testCases = append(testCases, testCase{
		desc: "success login user",
		body: models.UserLoginArgument{
			Email:    "candidate2@gmail.com",
			Password: "admin123",
		},
		code: 200,
	})

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			body, err := json.Marshal(&tC.body)
			AssertNoError(t, err)

			response := server.Request(http.MethodPost, urlLogin, strings.NewReader(string(body)))
			AssertEqualCode(t, response.StatusCode, tC.code)
		})
	}
}

func TestLogoutUser(t *testing.T) {
	db, err := InitPostgresTest()
	AssertNoError(t, err)

	server := NewServer(db)

	response := server.Request(http.MethodPost, urlLogout, nil)
	AssertEqualCode(t, response.StatusCode, 200)
}
