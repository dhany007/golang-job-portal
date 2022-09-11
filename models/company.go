package models

import "time"

type CompanyModel struct {
	ID    string `json:"id" db:"id"`
	Email string `json:"email" db:"email"`
}

type CompanySubCode struct {
	ID    int    `json:"id" db:"id"`
	Value string `json:"value" db:"value"`
}

type Company struct {
	ID               string           `json:"id,omitempty" db:"id"`
	Email            string           `json:"email" db:"email"`
	Name             string           `json:"name" db:"name"`
	Description      string           `json:"description" db:"description"`
	Address          string           `json:"address" db:"address"`
	Website          string           `json:"website" db:"website"`
	PhoneNumber      string           `json:"phone_number" db:"phone_number"`
	TelpNumber       string           `json:"telp_number" db:"telp_number"`
	ProfilPictureUrl string           `json:"profil_picture_url" db:"profil_picture_url"`
	Dress            string           `json:"dress" db:"dress"`
	Size             string           `json:"size" db:"size"`
	Benefit          []CompanySubCode `json:"benefit"`
	Review           CompanyRating    `json:"review"`
	CreatedAt        time.Time        `json:"created_at" db:"created_at"`
	ModifiedAt       time.Time        `json:"modified_at" db:"modified_at"`
}

type CompanyArgument struct {
	ID               string           `json:"id,omitempty" db:"id"`
	Email            string           `json:"email"`
	Name             string           `json:"name" valid:"required"`
	Description      string           `json:"description" valid:"required"`
	Address          string           `json:"address" valid:"required"`
	Website          string           `json:"website"`
	PhoneNumber      string           `json:"phone_number" valid:"required"`
	TelpNumber       string           `json:"telp_number"`
	ProfilPictureUrl string           `json:"profil_picture_url"`
	Dress            int              `json:"dress" valid:"required"`
	Size             int              `json:"size" valid:"required"`
	Benefit          []CompanySubCode `json:"benefit" valid:"required"`
}

type CompanyRating struct {
	Rating      float64 `json:"rating" db:"rating"`
	CountReview int     `json:"count_review" db:"count_review"`
}

type Companies struct {
	ID          string  `json:"id" db:"id"`
	Name        string  `json:"name" db:"name"`
	Rating      float64 `json:"rating" db:"rating"`
	CountReview int     `json:"count_review" db:"count_review"`
}
