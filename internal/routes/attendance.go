package routes

import (
	"payslips/internal/handlers"
	"payslips/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func AttendanceRouter(app *fiber.App, handler handlers.Handlers, m *middleware.Authentication) {
	app.Post("/attendance", m.Authentication, m.AuditLog, handler.Attendance.Attendance)
	app.Get("/attendance", m.Authentication, m.AuditLog, handler.Attendance.ListAttendance)
}
