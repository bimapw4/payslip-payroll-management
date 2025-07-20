package repositories

import (
	"payslips/internal/repositories/users"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Users users.Users
}

func NewRepository(db *sqlx.DB) Repository {
	return Repository{
		Users: users.NewRepo(db),
	}
}
