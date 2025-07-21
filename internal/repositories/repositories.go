package repositories

import (
	"payslips/internal/repositories/attendance"
	"payslips/internal/repositories/overtime"
	"payslips/internal/repositories/payroll"
	"payslips/internal/repositories/users"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Users      users.Users
	Attendance attendance.Attendance
	Payroll    payroll.Payroll
	Overtime   overtime.Overtime
}

func NewRepository(db *sqlx.DB) Repository {
	return Repository{
		Users:      users.NewRepo(db),
		Attendance: attendance.NewRepo(db),
		Payroll:    payroll.NewRepo(db),
		Overtime:   overtime.NewRepo(db),
	}
}
