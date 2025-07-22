package repositories

import (
	"payslips/internal/repositories/attendance"
	auditlog "payslips/internal/repositories/audit_log"
	"payslips/internal/repositories/overtime"
	"payslips/internal/repositories/payroll"
	payslipsummary "payslips/internal/repositories/payslip_summary"
	"payslips/internal/repositories/reimbursment"
	"payslips/internal/repositories/users"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Users          users.Users
	Attendance     attendance.Attendance
	Payroll        payroll.Payroll
	Overtime       overtime.Overtime
	Reimbursement  reimbursment.Reimbursment
	PayslipSummary payslipsummary.PayslipSummary
	AuditLog       auditlog.AuditLog
}

func NewRepository(db *sqlx.DB) Repository {
	return Repository{
		Users:          users.NewRepo(db),
		Attendance:     attendance.NewRepo(db),
		Payroll:        payroll.NewRepo(db),
		Overtime:       overtime.NewRepo(db),
		Reimbursement:  reimbursment.NewRepo(db),
		PayslipSummary: payslipsummary.NewRepo(db),
		AuditLog:       auditlog.NewRepo(db),
	}
}
