package tests

import (
	"net/http"
	"strings"
	"testing"
)

var (
	urlRegister = "http://localhost:11010/users/register"
	urlLogin    = "http://localhost:11010/users/login"
)

func TestRegisterUser(t *testing.T) {
	db, err := InitPostgresTest()
	AssertNoError(t, err)

	server := NewServer(db)

	db.Exec("TRUNCATE users;")
	db.Exec("TRUNCATE companies;")
	db.Exec("TRUNCATE candidates;")

	testCases := []struct {
		desc string
		body string
		code int
	}{
		{
			desc: "failed register, email not valid",
			body: `
			{
				"email": "manunggal",
				"password": "admin123",
				"role": 1
			}
		`,
			code: 400,
		},
		{
			desc: "failed register, email required",
			body: `
			{
				"password": "admin123",
				"role": 1
			}
		`,
			code: 400,
		},
		{
			desc: "failed register, password required",
			body: `
			{
				"email": "manunggal@gmail.com",
				"role": 1
			}
		`,
			code: 400,
		},
		{
			desc: "failed register, length password range(8|32)",
			body: `
			{
				"email": "manunggal@gmail.com",
				"password": "admin",
				"role": 1
			}
		`,
			code: 400,
		},
		{
			desc: "failed register, length password range(8|32)",
			body: `
			{
				"email": "manunggal@gmail.com",
				"password": "adminadminadminadminadminadminadmin",
				"role": 1
			}
		`,
			code: 400,
		},
		{
			desc: "success register as company",
			body: `
			{
				"email": "company1@gmail.com",
				"password": "admin123",
				"role": 1
			}
		`,
			code: 200,
		},
		{
			desc: "success register as company",
			body: `
			{
				"email": "company2@gmail.com",
				"password": "admin123",
				"role": 1
			}
		`,
			code: 200,
		},
		{
			desc: "failed register, email exists",
			body: `
			{
				"email": "company2@gmail.com",
				"password": "admin123",
				"role": 1
			}
		`,
			code: 400,
		},
		{
			desc: "success register as candidate",
			body: `
			{
				"email": "candidate1@gmail.com",
				"password": "admin123",
				"role": 2
			}
		`,
			code: 200,
		},
		{
			desc: "success register as candidate",
			body: `
			{
				"email": "candidate2@gmail.com",
				"password": "admin123",
				"role": 2
			}
		`,
			code: 200,
		},
		{
			desc: "failed register, email exists",
			body: `
			{
				"email": "candidate2@gmail.com",
				"password": "admin123",
				"role": 2
			}
		`,
			code: 400,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			response := server.Request(http.MethodPost, urlRegister, strings.NewReader(tC.body))
			AssertEqualCode(t, response.StatusCode, tC.code)
		})
	}
}

func TestLoginUser(t *testing.T) {
	db, err := InitPostgresTest()
	AssertNoError(t, err)

	server := NewServer(db)

	testCases := []struct {
		desc string
		body string
		code int
	}{
		{
			desc: "failed register, email not valid",
			body: `
			{
				"email": "company",
				"password": "admin123"
			}
		`,
			code: 400,
		},
		{
			desc: "failed register, email required",
			body: `
			{
				"password": "admin123"
			}
		`,
			code: 400,
		},
		{
			desc: "failed register, password required",
			body: `
			{
				"email": "company@gmail.com"
			}
		`,
			code: 400,
		},
		{
			desc: "failed register, users not found",
			body: `
			{
				"email": "company@gmail.com",
				"password": "admin"
			}
		`,
			code: 400,
		},
		{
			desc: "success login user",
			body: `
			{
				"email": "company1@gmail.com",
				"password": "admin123"
			}
		`,
			code: 200,
		},
		{
			desc: "success login user",
			body: `
			{
				"email": "candidate1@gmail.com",
				"password": "admin123"
			}
		`,
			code: 200,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			response := server.Request(http.MethodPost, urlLogin, strings.NewReader(tC.body))
			AssertEqualCode(t, response.StatusCode, tC.code)
		})
	}
}
