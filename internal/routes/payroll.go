package routes

import (
	"payslips/internal/handlers"
	"payslips/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func PayrollRouter(app *fiber.App, handler handlers.Handlers, m *middleware.Authentication) {
	app.Post("/payroll", m.Authentication, handler.Payroll.Create)
	app.Put("/payroll/running/:id", m.Authentication, handler.Payroll.Running)
	app.Get("/payroll/generate/payslips/:id", m.Authentication, handler.Payroll.GeneratePayslip)
	app.Get("/payroll/summary/payslip/:id", m.Authentication, handler.Payroll.SummaryPayslip)
}
