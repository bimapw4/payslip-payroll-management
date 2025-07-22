package routes

import (
	"payslips/internal/handlers"
	"payslips/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func PayrollRouter(app *fiber.App, handler handlers.Handlers, m *middleware.Authentication) {
	app.Post("/payroll", m.Authentication, m.AuditLog, handler.Payroll.Create)
	app.Put("/payroll/:id", m.Authentication, m.AuditLog, handler.Payroll.Update)
	app.Get("/payroll", m.Authentication, m.AuditLog, handler.Payroll.List)
	app.Put("/payroll/running/:id", m.Authentication, m.AuditLog, handler.Payroll.Running)
	app.Get("/payroll/generate/payslips/:id", m.Authentication, m.AuditLog, handler.Payroll.GeneratePayslip)
	app.Get("/payroll/generate/payslips/:id/user/:user_id", m.Authentication, m.AuditLog, handler.Payroll.GeneratePayslipAdmin)
	app.Get("/payroll/summary/payslip/:id", m.Authentication, m.AuditLog, handler.Payroll.SummaryPayslip)
}
