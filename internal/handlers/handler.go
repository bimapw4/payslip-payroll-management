package handlers

import (
	"payslips/internal/business"

	"payslips/internal/handlers/attendance"
	"payslips/internal/handlers/auth"
)

type Handlers struct {
	Auth       auth.Handler
	Attendance attendance.Handler
}

func NewHandler(business business.Business) Handlers {
	return Handlers{
		Auth:       auth.NewHandler(business),
		Attendance: attendance.NewHandler(business),
	}
}
