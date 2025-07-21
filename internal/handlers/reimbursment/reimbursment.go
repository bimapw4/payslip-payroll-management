package reimbursment

import (
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"payslips/internal/business"
	"payslips/internal/entity"
	"payslips/internal/response"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Preview(c *fiber.Ctx) error
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

	amountstr := c.FormValue("amount")

	description := c.FormValue("description")

	attachment, err := c.FormFile("attachment")
	if err != nil {
		return response.NewResponse(Entity).
			Errors("Failed to parse request body", nil).
			JSON(c, fiber.StatusInternalServerError)
	}

	ext := filepath.Ext(attachment.Filename)
	switch ext {
	case ".png", ".jpg", ".jpeg":
	default:
		err = errors.New("content type must be a html")
		return response.NewResponse(Entity).
			Errors("Failed to parse request body", err.Error()).
			JSON(c, fiber.StatusBadRequest)
	}

	if attachment.Size > (5 * 1024 * 1024) {
		err = errors.New("file size cannot be more than 5MB")
		return response.NewResponse(Entity).
			Errors("Failed to parse request body", err.Error()).
			JSON(c, fiber.StatusBadRequest)
	}

	path := fmt.Sprintf("./storage/%v-%v", time.Now().Unix(), attachment.Filename)

	err = c.SaveFile(attachment, path)
	if err != nil {
		log.Println(err)
		return response.NewResponse(Entity).
			Errors("Failed to parse request body", err.Error()).
			JSON(c, fiber.StatusBadRequest)
	}

	result, err := h.business.Reimbursment.Create(c.UserContext(), entity.ReimbursementCreate{
		Amount:      amountstr,
		Description: description,
		Attachment:  path,
	})
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

	id := c.Params("id")

	amountstr := c.FormValue("amount")

	description := c.FormValue("description")

	attachment, err := c.FormFile("attachment")
	if err != nil {
		return response.NewResponse("CreatePayment").
			Errors("Failed to parse request body", nil).
			JSON(c, fiber.StatusInternalServerError)
	}

	ext := filepath.Ext(attachment.Filename)
	switch ext {
	case ".png", ".jpg", ".jpeg":
	default:
		err = errors.New("content type must be a html")
		return response.NewResponse("CreatePayment").
			Errors("Failed to parse request body", err.Error()).
			JSON(c, fiber.StatusBadRequest)
	}

	if attachment.Size > (5 * 1024 * 1024) {
		err = errors.New("file size cannot be more than 5MB")
		return response.NewResponse("CreatePayment").
			Errors("Failed to parse request body", err.Error()).
			JSON(c, fiber.StatusBadRequest)
	}

	path := fmt.Sprintf("./storage/%v-%v", time.Now().Unix(), attachment.Filename)

	err = c.SaveFile(attachment, path)
	if err != nil {
		log.Println(err)
		return response.NewResponse("CreatePayment").
			Errors("Failed to parse request body", err.Error()).
			JSON(c, fiber.StatusBadRequest)
	}

	result, err := h.business.Reimbursment.Update(c.UserContext(), entity.ReimbursementUpdate{
		Amount:      amountstr,
		Description: description,
		Attachment:  path,
		Id:          id,
	})
	if err != nil {
		return response.NewResponse(Entity).
			Errors("Failed update reimbursment", err).
			JSON(c, fiber.StatusBadRequest)
	}

	return response.NewResponse(Entity).
		Success("Success Update Reimbursment", result).
		JSON(c, fiber.StatusOK)
}

func (h *handler) Preview(c *fiber.Ctx) error {
	detail, err := h.business.Reimbursment.Detail(c.Context(), c.Params("id"))
	if err != nil {
		return response.NewResponse("ReimbursmentPreview").
			Errors("Failed to fetch reservation", err).
			JSON(c, fiber.StatusInternalServerError)
	}

	return c.SendFile(detail.Attachment)
}
