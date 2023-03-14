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

type CandidateExperienceArgument struct {
	CandidateID string `json:"candidate_id" valid:"required"`
	CompanyName string `json:"company_name" valid:"required"`
	Title       string `json:"title"  valid:"required"`
	Description string `json:"description"`
	DateStart   string `json:"date_start" valid:"required"`
	DateEnd     string `json:"date_end"`
}

type CandidateExperience struct {
	ID          int       `json:"id,omitempty" db:"id"`
	CandidateID string    `json:"candidate_id" db:"candidate_id"`
	CompanyName string    `json:"company_name" db:"company_name"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	DateStart   time.Time `json:"date_start" db:"date_start"`
	DateEnd     time.Time `json:"date_end" db:"date_end"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	ModifiedAt  time.Time `json:"modified_at" db:"modified_at"`
}

type UpdateCandidateExperienceArgument struct {
	CompanyName string `json:"company_name"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DateStart   string `json:"date_start"`
	DateEnd     string `json:"date_end"`
}
