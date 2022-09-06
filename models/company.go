package models

type CompanyModel struct {
	ID    string `json:"id" db:"id"`
	Email string `json:"email" db:"email"`
}

type CompanySubCode struct {
	ID    int    `json:"id" db:"id"`
	Value string `json:"value" db:"value"`
}
