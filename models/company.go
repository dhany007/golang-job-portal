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

type ReviewCompany struct {
	ID          int    `json:"id" db:"id"`
	CompanyID   string `json:"company_id,omitempty" db:"company_id"`
	CandidateID string `json:"candidate_id" db:"candidate_id"`
	Rating      int    `json:"rating" db:"rating"`
	Review      string `json:"review" db:"review"`
}

type ReviewCompanyArgument struct {
	CompanyID   string `json:"company_id,omitempty" valid:"required"`
	CandidateID string `json:"candidate_id" valid:"required"`
	Rating      int    `json:"rating" valid:"required,numeric,range(1|5)"`
	Review      string `json:"review" valid:"required"`
}

type ListData struct {
	Offset      int `json:"offset,omitempty"`
	Limit       int `json:"limit,omitempty"`
	Page        int `json:"page,omitempty"`
	ItemPerPage int `json:"item_per_page,omitempty"`
}

type ListReviewCompany struct {
	Page        int             `json:"page"`
	TotalPage   int             `json:"total_page"`
	ItemPerPage int             `json:"item_per_page"`
	TotalData   int             `json:"total_data"`
	Reviews     []ReviewCompany `json:"reviews"`
}

type ListCompanies struct {
	Page        int         `json:"page"`
	TotalPage   int         `json:"total_page"`
	ItemPerPage int         `json:"item_per_page"`
	TotalData   int         `json:"total_data"`
	Companies   []Companies `json:"companies"`
}
