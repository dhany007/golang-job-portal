package models

import "time"

type Candidate struct {
	ID               string    `json:"id" db:"id"`
	Email            string    `json:"email" db:"email"`
	FirstName        string    `json:"first_name" db:"first_name"`
	LastName         string    `json:"last_name" db:"last_name"`
	PhoneNumber      string    `json:"phone_number" db:"phone_number"`
	TelpNumber       string    `json:"telp_number" db:"telp_number"`
	Address          string    `json:"address" db:"address"`
	ProfilPictureUrl string    `json:"profil_picture_url" db:"profil_picture_url"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
	ModifiedAt       time.Time `json:"modified_at" db:"modified_at"`
}

type CandidateArgument struct {
	ID               string `json:"id,omitempty"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	PhoneNumber      string `json:"phone_number" valid:"required"`
	TelpNumber       string `json:"telp_number"`
	Address          string `json:"address" valid:"required"`
	ProfilPictureUrl string `json:"profil_picture_url" db:"profil_picture_url"`
}
