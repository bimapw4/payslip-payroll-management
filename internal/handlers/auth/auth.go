package auth

import (
	"payslips/internal/business"
	"payslips/internal/common"
	"payslips/internal/entity"
	"payslips/internal/response"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	Login(c *fiber.Ctx) error
}

type handler struct {
	business business.Business
}

func NewHandler(business business.Business) Handler {
	return &handler{
		business: business,
	}
}

func (h *handler) Login(c *fiber.Ctx) error {

	var payload entity.Authorization
	if err := c.BodyParser(&payload); err != nil {
		return response.NewResponse("Login").
			Errors("Failed to parse request body", err).
			JSON(c, fiber.StatusBadRequest)
	}

	resp, err := h.business.Auth.Authorization(c.Context(), payload)
	if err != nil {
		return response.NewResponse("Login").
			Errors("Failed to login", common.ErrUnauthorized).
			JSON(c, fiber.StatusUnauthorized)
	}

	return response.NewResponse("Login").
		Success("Login successfully", resp).
		JSON(c, fiber.StatusOK)
}
