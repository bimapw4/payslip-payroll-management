package handlers

import (
	"payslips/internal/business"

	"payslips/internal/handlers/attendance"
	"payslips/internal/handlers/auth"
	"payslips/internal/handlers/overtime"
	"payslips/internal/handlers/payroll"
)

type Handlers struct {
	Auth       auth.Handler
	Attendance attendance.Handler
	Payroll    payroll.Handler
	Overtime   overtime.Handler
}

func NewHandler(business business.Business) Handlers {
	return Handlers{
		Auth:       auth.NewHandler(business),
		Attendance: attendance.NewHandler(business),
		Payroll:    payroll.NewHandler(business),
		Overtime:   overtime.NewHandler(business),
	}
}
