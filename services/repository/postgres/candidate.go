package postgres

import (
	"context"
	"log"

	"github.com/dhany007/golang-job-portal/models"
	"github.com/dhany007/golang-job-portal/services"
	"github.com/dhany007/golang-job-portal/services/repository/database"
	"github.com/dhany007/golang-job-portal/services/repository/postgres/queries"
	"github.com/jmoiron/sqlx"
)

type candidateRepository struct {
	DB *database.DB
}

func NewCandidateRepository(db *database.DB) services.CandidateRepository {
	return &candidateRepository{DB: db}
}

// CheckCandidateByEmail implements services.CandidateRepository
func (c candidateRepository) CheckCandidateByEmail(ctx context.Context, email string) (result models.Candidate, err error) {
	var (
		rows *sqlx.Rows
	)

	rows, err = c.DB.QueryxContext(ctx, queries.QueryCheckAvailableEmail, email)
	if err != nil {
		log.Printf("[candidate] [repository] [CheckCandidateByEmail] while QueryCheckAvailableEmail, err:%+v\n", err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.StructScan(&result)
		if err != nil {
			log.Printf("[candidate] [repository] [CheckCandidateByEmail] while StructScan, err:%+v\n", err)
			return
		}
	}

	return
}

// CheckCandidateById implements services.CandidateRepository
func (c candidateRepository) CheckCandidateById(ctx context.Context, id string) (result models.Candidate, err error) {
	var (
		rows *sqlx.Rows
	)

	rows, err = c.DB.QueryxContext(ctx, queries.QueryCheckUsersById, id)
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
		queries.QueryUpdateCandidate,
		args.Email,
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

	rows, err = c.DB.QueryxContext(ctx, queries.QueryGetDetailCandidate, id)
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
