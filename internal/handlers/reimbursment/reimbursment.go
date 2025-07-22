package reimbursment

import (
	"log"
	"payslips/internal/business"
	"payslips/internal/entity"
	"payslips/internal/response"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	Create(c *fiber.Ctx) error
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
		Entity = "ReimbursmentCreate"
	)

	var payload entity.ReimbursementCreate
	if err := c.BodyParser(&payload); err != nil {
		return response.NewResponse(Entity).
			Errors("Failed to parse request body", err).
			JSON(c, fiber.StatusBadRequest)
	}

	result, err := h.business.Reimbursment.Create(c.UserContext(), payload)
	if err != nil {
		return response.NewResponse(Entity).
			Errors("Failed create reimbursment", err).
			JSON(c, fiber.StatusBadRequest)
	}

	return response.NewResponse(Entity).
		Success("Success Create Reimbursment", result).
		JSON(c, fiber.StatusOK)
}

func (h *handler) Update(c *fiber.Ctx) error {
	var (
		Entity = "ReimbursmentUpdate"
	)

	var payload entity.ReimbursementUpdate
	if err := c.BodyParser(&payload); err != nil {
		return response.NewResponse(Entity).
			Errors("Failed to parse request body", err).
			JSON(c, fiber.StatusBadRequest)
	}

	result, err := h.business.Reimbursment.Update(c.UserContext(), payload)
	if err != nil {
		log.Println("err", err.Error())
		return response.NewResponse(Entity).
			Errors("Failed update reimbursment", err).
			JSON(c, fiber.StatusBadRequest)
	}

	return response.NewResponse(Entity).
		Success("Success Update Reimbursment", result).
		JSON(c, fiber.StatusOK)
}
