package postgres

import (
	"context"
	"log"

	_ "embed"

	"github.com/dhany007/golang-job-portal/models"
	"github.com/dhany007/golang-job-portal/services"
	"github.com/dhany007/golang-job-portal/services/repository/database"
	"github.com/jmoiron/sqlx"
)

type candidateRepository struct {
	DB *database.DB
}

func NewCandidateRepository(db *database.DB) services.CandidateRepository {
	return &candidateRepository{DB: db}
}

// CheckCandidateByEmail implements services.CandidateRepository
func (c candidateRepository) CheckCandidateByEmail(ctx context.Context, email string) (result []models.Candidate, err error) {
	var (
		rows      *sqlx.Rows
		candidate models.Candidate
	)

	rows, err = c.DB.QueryxContext(ctx, QueryCheckAvailableEmail, email)
	if err != nil {
		log.Printf("[candidate] [repository] [CheckCandidateByEmail] while QueryCheckAvailableEmail, err:%+v\n", err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.StructScan(&candidate)
		if err != nil {
			log.Printf("[candidate] [repository] [CheckCandidateByEmail] while StructScan, err:%+v\n", err)
			return
		}
		result = append(result, candidate)
	}

	return
}

// CheckCandidateById implements services.CandidateRepository
func (c candidateRepository) CheckCandidateById(ctx context.Context, id string) (result models.Candidate, err error) {
	var (
		rows *sqlx.Rows
	)

	rows, err = c.DB.QueryxContext(ctx, QueryCheckUsersById, id)
	if err != nil {
		log.Printf("[candidate] [repository] [CheckCandidateById] while QueryCheckUsersById, err:%+v\n", err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.StructScan(&result)
		if err != nil {
			log.Printf("[candidate] [repository] [CheckCandidateById] while StructScan, err:%+v\n", err)
			return
		}
	}

	return
}

// UpdateCandidate implements services.CandidateRepository
func (c candidateRepository) UpdateCandidate(ctx context.Context, args models.Candidate) (err error) {
	_, err = c.DB.ExecContext(
		ctx,
		QueryUpdateCandidate,
		args.FirstName,
		args.LastName,
		args.PhoneNumber,
		args.TelpNumber,
		args.Address,
		args.ProfilPictureUrl,
		args.ID,
	)

	if err != nil {
		log.Printf("[candidate] [repository] [UpdateCandidate] while QueryUpdateCandidate, err:%+v\n", err)
		return
	}

	return
}

// GetDetailCandidate implements services.CandidateRepository
func (c candidateRepository) GetDetailCandidate(ctx context.Context, id string) (result models.Candidate, err error) {
	var (
		rows *sqlx.Rows
	)

	rows, err = c.DB.QueryxContext(ctx, QueryGetDetailCandidate, id)
	if err != nil {
		log.Printf("[candidate] [repository] [GetDetailCandidate] while QueryGetDetailCandidate, err:%+v\n", err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.StructScan(&result)
		if err != nil {
			log.Printf("[candidate] [repository] [GetDetailCandidate] while StructScan, err:%+v\n", err)
			return
		}
	}

	return
}

// AddExperience implements services.CandidateRepository
func (c candidateRepository) AddExperience(ctx context.Context, args models.CandidateExperience) (result models.CandidateExperience, err error) {
	_, err = c.DB.ExecContext(
		ctx,
		QueryAddExperience,
		args.CandidateID,
		args.CompanyName,
		args.Title,
		args.Description,
		args.DateStart,
		args.DateEnd,
		args.CreatedAt,
		args.ModifiedAt,
	)

	if err != nil {
		log.Printf("[candidate] [repository] [AddExperience] while QueryAddExperience, err:%+v\n", err)
		return
	}

	return args, nil
}
