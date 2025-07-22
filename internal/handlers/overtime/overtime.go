package overtime

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
	Update(c *fiber.Ctx) error
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

func (h *handler) Create(c *fiber.Ctx) error {

	var (
		Entity = "Overtime"
	)

	var payload entity.Overtime
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

	re, err := h.business.Overtime.Overtime(c.UserContext(), payload)
	if err != nil {
		return response.NewResponse(Entity).
			Errors("Failed create overtime", err.Error()).
			JSON(c, fiber.StatusBadRequest)
	}

	return response.NewResponse(Entity).
		Success("Create Overtime Successfully", re).
		JSON(c, fiber.StatusOK)
}

func (h *handler) Update(c *fiber.Ctx) error {

	var (
		Entity = "UpdateOvertime"
	)

	var payload entity.Overtime
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

	id := c.Params("id")

	if err := validation.Validate(id, is.UUID); err != nil {
		return response.NewResponse(Entity).
			Errors("Failed update overtime", err.Error()).
			JSON(c, fiber.StatusBadRequest)
	}

	re, err := h.business.Overtime.Update(c.UserContext(), payload, id)
	if err != nil {
		return response.NewResponse(Entity).
			Errors("Failed update overtime", err.Error()).
			JSON(c, fiber.StatusBadRequest)
	}

	return response.NewResponse(Entity).
		Success("Update Overtime Successfully", re).
		JSON(c, fiber.StatusOK)
}

func (h *handler) List(c *fiber.Ctx) error {

	var (
		Entity = "ListOvertime"
	)

	q := c.Queries()

	m := meta.NewParams(q)

	result, err := h.business.Overtime.List(c.UserContext(), &m)
	if err != nil {
		return response.NewResponse(Entity).
			Errors("Failed to list overtime", err.Error()).
			JSON(c, fiber.StatusBadRequest)
	}

	return response.NewResponse(Entity).
		SuccessWithMeta("List Overtime successfully", result, m).
		JSON(c, fiber.StatusOK)
}
