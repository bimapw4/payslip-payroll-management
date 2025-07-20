package repositories

import (
	"payslips/internal/repositories/attendance"
	"payslips/internal/repositories/users"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Users      users.Users
	Attendance attendance.Attendance
}

func NewRepository(db *sqlx.DB) Repository {
	return Repository{
		Users:      users.NewRepo(db),
		Attendance: attendance.NewRepo(db),
	}
}
