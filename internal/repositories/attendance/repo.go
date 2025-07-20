package attendance

import (
	"database/sql"
	"payslips/internal/presentations"
	"payslips/pkg/databasex"

	"github.com/jmoiron/sqlx"
)

type repo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) Attendance {
	return &repo{
		db: db,
	}
}

func (r *repo) translateError(err error) error {
	switch err {
	case sql.ErrNoRows:
		return presentations.ErrAttendanceNotExist
	case databasex.ErrUniqueViolation:
		return presentations.ErrAttendanceAlreadyExist
	default:
		return err
	}
}
