package postgres

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/dhany007/golang-job-portal/models"
	"github.com/dhany007/golang-job-portal/services"
	"github.com/dhany007/golang-job-portal/services/repository/database"
	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	DB *database.DB
}

func NewUserRepository(db *database.DB) services.UserRepository {
	return &userRepository{db}
}

func (u userRepository) GetUserByEmail(ctx context.Context, email string) (result models.User, err error) {
	var (
		row *sqlx.Rows
	)
	row, err = u.DB.QueryxContext(ctx, QueryGetUserByEmail, email)

	if err != nil {
		log.Printf("[user] [repository] [GetUserByEmail] while QueryxContext, err:%+v\n", err)
		return result, err
	}

	defer row.Close()

	for row.Next() {
		err = row.StructScan(&result)
		if err != nil {
			log.Printf("[user] [repository] [GetUserByEmail] while StructScan, err:%+v\n", err)
			return result, err
		}
	}

	return result, err
}

func (u userRepository) CreateUser(ctx context.Context, args models.User) (result models.User, err error) {
	var (
		name string = strings.Split(args.Email, "@")[0]
	)

	// insert user
	_, err = u.DB.ExecContext(
		ctx,
		QueryInsertUser,
		args.ID,
		args.Email,
		args.Password,
		args.Role,
	)

	if err != nil {
		log.Printf("[user] [repository] [CreateUser] while QueryInsertUser, email:%+v\n", args.Email)
		return
	}

	// check if role: 1=company, 2=candidate
	if args.Role == 1 {
		// insert company
		_, err = u.DB.ExecContext(
			ctx,
			QueryInsertCompany,
			args.ID,
			args.Email,
			name,
		)

		if err != nil {
			log.Printf("[user] [repository] [CreateUser] while QueryInsertCompany, email:%+v\n", args.Email)
			return
		}
	} else if args.Role == 2 {
		// insert candidate
		_, err = u.DB.ExecContext(
			ctx,
			QueryInsertCandidate,
			args.ID,
			args.Email,
			name,
		)

		if err != nil {
			log.Printf("[user] [repository] [CreateUser] while QueryInsertCandidate, email:%+v\n", args.Email)
			return
		}
	}

	// reinitialize result
	result.ID = args.ID
	result.Email = args.Email
	result.Role = args.Role
	result.IsActive = 1
	result.CreatedAt = time.Now()
	result.ModifiedAt = time.Now()

	return result, err
}

func (u userRepository) GetUserById(ctx context.Context, id string) (result models.User, err error) {
	var (
		row *sqlx.Rows
	)
	row, err = u.DB.QueryxContext(ctx, QueryGetUserById, id)

	if err != nil {
		log.Printf("[user] [repository] [GetUserById] while QueryxContext, err:%+v\n", err)
		return result, err
	}

	defer row.Close()

	for row.Next() {
		err = row.StructScan(&result)
		if err != nil {
			log.Printf("[user] [repository] [GetUserById] while StructScan, err:%+v\n", err)
			return result, err
		}
	}

	return result, err
}
