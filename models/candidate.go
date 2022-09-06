package models

type CandidateModel struct {
	ID        string `json:"id" db:"id"`
	Email     string `json:"email" db:"email"`
	FirstName string `json:"first_name" db:"first_name"`
}
