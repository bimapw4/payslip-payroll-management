package payroll

import (
	"payslips/internal/business"
	"payslips/internal/entity"
	"payslips/internal/response"
	"payslips/pkg/meta"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	Create(c *fiber.Ctx) error
	Running(c *fiber.Ctx) error
	GeneratePayslip(c *fiber.Ctx) error
	SummaryPayslip(c *fiber.Ctx) error
}

type handler struct {
	business business.Business
}

func NewHandler(business business.Business) Handler {
	return &handler{
		business: business,
	}
}

func (h *handler) Create(c *fiber.Ctx) error {
	var (
		Entity = "Payroll"
	)

	var payload entity.Payroll
	if err := c.BodyParser(&payload); err != nil {
		return response.NewResponse(Entity).
			Errors("Failed to parse request body", err).
			JSON(c, fiber.StatusBadRequest)
	}

	if err := payload.Validation(); err != nil {
		return response.NewResponse(Entity).
			Errors("Failed to parse request body", err).
			JSON(c, fiber.StatusBadRequest)
	}

	err := h.business.Payroll.CreatePayroll(c.UserContext(), payload)
	if err != nil {
		return response.NewResponse(Entity).
			Errors("Failed create payroll", err).
			JSON(c, fiber.StatusBadRequest)
	}

	return response.NewResponse(Entity).
		Success("Success Create Payroll", nil).
		JSON(c, fiber.StatusOK)
}

func (h *handler) Running(c *fiber.Ctx) error {
	var (
		Entity = "RunningPayroll"
	)

	if err := validation.Validate(c.Params("id"), is.UUID); err != nil {
		return response.NewResponse(Entity).
			Errors("Failed running payroll", err.Error()).
			JSON(c, fiber.StatusBadRequest)
	}

	err := h.business.Payroll.RunningPayroll(c.UserContext(), c.Params("id"))
	if err != nil {
		return response.NewResponse(Entity).
			Errors("Failed running payroll", err).
			JSON(c, fiber.StatusBadRequest)
	}

	return response.NewResponse(Entity).
		Success("Success Running Payroll", nil).
		JSON(c, fiber.StatusOK)
}

func (h *handler) GeneratePayslip(c *fiber.Ctx) error {
	var (
		Entity = "GeneratePayslip"
	)

	if err := validation.Validate(c.Params("id"), is.UUID); err != nil {
		return response.NewResponse(Entity).
			Errors("Failed payroll generate payslip", err.Error()).
			JSON(c, fiber.StatusBadRequest)
	}

	result, err := h.business.Payroll.GeneratePayslip(c.UserContext(), c.Params("id"))
	if err != nil {
		return response.NewResponse(Entity).
			Errors("Failed payroll generate payslip", err).
			JSON(c, fiber.StatusBadRequest)
	}

	return response.NewResponse(Entity).
		Success("Success Payroll Generate Payslip", result).
		JSON(c, fiber.StatusOK)
}

func (h *handler) SummaryPayslip(c *fiber.Ctx) error {
	var (
		Entity = "SummaryPayslip"
	)

	if err := validation.Validate(c.Params("id"), is.UUID); err != nil {
		return response.NewResponse(Entity).
			Errors("Failed payroll summary payslip", err.Error()).
			JSON(c, fiber.StatusBadRequest)
	}

	query := c.Queries()

	m := meta.NewParams(query)

	result, err := h.business.Payroll.ListSummary(c.UserContext(), &m, c.Params("id"))
	if err != nil {
		return response.NewResponse(Entity).
			Errors("Failed payroll summary payslip", err).
			JSON(c, fiber.StatusBadRequest)
	}

	return response.NewResponse(Entity).
		SuccessWithMeta("Success Payroll Summary Payslip", result, m).
		JSON(c, fiber.StatusOK)
}
