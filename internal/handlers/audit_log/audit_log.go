package auditlog

import (
	"payslips/internal/business"
	"payslips/internal/response"
	"payslips/pkg/meta"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	List(c *fiber.Ctx) error
}

type handler struct {
	business business.Business
}

func NewHandler(business business.Business) Handler {
	return &handler{
		business: business,
	}
}

func (h *handler) List(c *fiber.Ctx) error {
	var (
		Entity = "SummaryPayslip"
	)

	query := c.Queries()

	m := meta.NewParams(query)

	result, err := h.business.AuditLog.List(c.UserContext(), &m)
	if err != nil {
		return response.NewResponse(Entity).
			Errors("Failed list audit log", err).
			JSON(c, fiber.StatusBadRequest)
	}

	return response.NewResponse(Entity).
		SuccessWithMeta("Success List Audit Log", result, m).
		JSON(c, fiber.StatusOK)
}
