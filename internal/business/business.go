package business

import (
	"payslips/internal/business/attendance"
	auditlog "payslips/internal/business/audit_log"
	"payslips/internal/business/auth"
	"payslips/internal/business/overtime"
	"payslips/internal/business/payroll"
	"payslips/internal/business/reimbursment"
	"payslips/internal/repositories"
)

type Business struct {
	Auth         auth.Contract
	Attendance   attendance.Contract
	Payroll      payroll.Contract
	Overtime     overtime.Contract
	Reimbursment reimbursment.Contract
	AuditLog     auditlog.Contract
}

func NewBusiness(repo *repositories.Repository) Business {
	return Business{
		Auth:         auth.NewBusiness(repo),
		Attendance:   attendance.NewBusiness(repo),
		Payroll:      payroll.NewBusiness(repo),
		Overtime:     overtime.NewBusiness(repo),
		Reimbursment: reimbursment.NewBusiness(repo),
		AuditLog:     auditlog.NewBusiness(repo),
	}
}
