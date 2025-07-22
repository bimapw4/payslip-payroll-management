package payroll

import (
	"fmt"
	"payslips/internal/business"
	"payslips/internal/common"
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
	GeneratePayslipAdmin(c *fiber.Ctx) error
	SummaryPayslip(c *fiber.Ctx) error
	List(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
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

	userctx := common.GetUserCtx(c.UserContext())
	if !userctx.IsAdmin {
		return response.NewResponse(Entity).
			Errors("Failed payroll generate payslip", "only admin allowed").
			JSON(c, fiber.StatusForbidden)
	}

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

	res, err := h.business.Payroll.CreatePayroll(c.UserContext(), payload)
	if err != nil {
		return response.NewResponse(Entity).
			Errors("Failed create payroll", err).
			JSON(c, fiber.StatusBadRequest)
	}

	return response.NewResponse(Entity).
		Success("Success Create Payroll", res).
		JSON(c, fiber.StatusOK)
}

func (h *handler) Update(c *fiber.Ctx) error {
	var (
		Entity = "UpdatePayroll"
	)

	userctx := common.GetUserCtx(c.UserContext())
	if !userctx.IsAdmin {
		return response.NewResponse(Entity).
			Errors("Failed payroll generate payslip", "only admin allowed").
			JSON(c, fiber.StatusForbidden)
	}

	if err := validation.Validate(c.Params("id"), is.UUID); err != nil {
		return response.NewResponse(Entity).
			Errors("Failed running payroll", err.Error()).
			JSON(c, fiber.StatusBadRequest)
	}

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

	res, err := h.business.Payroll.UpdatePayroll(c.UserContext(), payload, c.Params("id"))
	if err != nil {
		return response.NewResponse(Entity).
			Errors("Failed update payroll", err).
			JSON(c, fiber.StatusBadRequest)
	}

	return response.NewResponse(Entity).
		Success("Success Update Payroll", res).
		JSON(c, fiber.StatusOK)
}

func (h *handler) List(c *fiber.Ctx) error {
	var (
		Entity = "ListPayroll"
	)

	userctx := common.GetUserCtx(c.UserContext())
	if !userctx.IsAdmin {
		return response.NewResponse(Entity).
			Errors("Failed payroll generate payslip", "only admin allowed").
			JSON(c, fiber.StatusForbidden)
	}

	query := c.Queries()

	m := meta.NewParams(query)

	result, err := h.business.Payroll.List(c.UserContext(), &m)
	if err != nil {
		return response.NewResponse(Entity).
			Errors("Failed List payroll", err).
			JSON(c, fiber.StatusBadRequest)
	}

	return response.NewResponse(Entity).
		SuccessWithMeta("Success List Payroll", result, m).
		JSON(c, fiber.StatusOK)
}

func (h *handler) Running(c *fiber.Ctx) error {
	var (
		Entity = "RunningPayroll"
	)

	userctx := common.GetUserCtx(c.UserContext())
	if !userctx.IsAdmin {
		return response.NewResponse(Entity).
			Errors("Failed payroll generate payslip", "only admin allowed").
			JSON(c, fiber.StatusForbidden)
	}

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

	userctx := common.GetUserCtx(c.UserContext())

	if err := validation.Validate(c.Params("id"), is.UUID); err != nil {
		return response.NewResponse(Entity).
			Errors("Failed payroll generate payslip", err.Error()).
			JSON(c, fiber.StatusBadRequest)
	}

	result, err := h.business.Payroll.GeneratePayslip(c.UserContext(), c.Params("id"), userctx.UserID)
	if err != nil {
		return response.NewResponse(Entity).
			Errors("Failed payroll generate payslip", err).
			JSON(c, fiber.StatusBadRequest)
	}

	return response.NewResponse(Entity).
		Success("Success Payroll Generate Payslip", result).
		JSON(c, fiber.StatusOK)
}

func (h *handler) GeneratePayslipAdmin(c *fiber.Ctx) error {
	var (
		Entity = "GeneratePayslip"
	)

	userctx := common.GetUserCtx(c.UserContext())
	if !userctx.IsAdmin {
		return response.NewResponse(Entity).
			Errors("Failed payroll generate payslip", "only admin allowed").
			JSON(c, fiber.StatusForbidden)
	}

	if err := validation.Validate(c.Params("id"), is.UUID); err != nil {
		return response.NewResponse(Entity).
			Errors("Failed payroll generate payslip", err.Error()).
			JSON(c, fiber.StatusBadRequest)
	}

	if err := validation.Validate(c.Params("user_id"), is.UUID); err != nil {
		return response.NewResponse(Entity).
			Errors("Failed payroll generate payslip", fmt.Sprintf("user id %v", err.Error())).
			JSON(c, fiber.StatusBadRequest)
	}

	result, err := h.business.Payroll.GeneratePayslip(c.UserContext(), c.Params("id"), c.Params("user_id"))
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

	userctx := common.GetUserCtx(c.UserContext())
	if !userctx.IsAdmin {
		return response.NewResponse(Entity).
			Errors("Failed payroll generate payslip", "only admin allowed").
			JSON(c, fiber.StatusForbidden)
	}

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
