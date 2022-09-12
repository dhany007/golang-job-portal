package models

import "time"

type User struct {
	ID         string    `json:"id" db:"id"`
	Email      string    `json:"email" db:"email"`
	IsActive   int       `json:"is_active" db:"is_active"`
	Role       int       `json:"role" db:"role"`
	Password   string    `json:"password,omitempty" db:"hash_password"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	ModifiedAt time.Time `json:"modified_at" db:"modified_at"`
}

type UserRegisterArguments struct {
	Email    string `json:"email" valid:"required,email"`
	Password string `json:"password" valid:"required,length(6|32)"`
	Role     int    `json:"role" valid:"required,numeric,range(1|2)"`
}

type UserLoginArgument struct {
	Email    string `json:"email" valid:"required,email"`
	Password string `json:"password" valid:"required,length(6|32)"`
}

type UserLoginResponse struct {
	RefreshToken string `json:"refresh_token,omitempty"`
	AccessToken  string `json:"access_token"`
}

type Authorization struct {
	ID    string
	Email string
	Role  int
}
