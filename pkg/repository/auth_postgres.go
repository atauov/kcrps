package repository

import (
	"fmt"
	"github.com/atauov/kcrps/models/request"
	"github.com/jmoiron/sqlx"
	"time"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user request.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (created_at, company_name, username, password_hash, role) values ($1, $2, $3, $4, $5) RETURNING id", usersTable)
	row := r.db.QueryRow(query, time.Now(), user.CompanyName, user.Username, user.Password, user.Role)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (request.User, error) {
	var user request.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, username, password)

	return user, err
}

func (r *AuthPostgres) GetUserIdByApiKey(api string) (int, error) {
	var userId int
	query := fmt.Sprintf("SELECT id FROM %s WHERE api_key=$1", usersTable)
	err := r.db.Get(&userId, query, api)

	return userId, err
}
