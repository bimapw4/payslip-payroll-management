package middleware

import (
	"encoding/json"
	"payslips/internal/common"
	"payslips/internal/presentations"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (m *Authentication) AuditLog(c *fiber.Ctx) error {
	start := time.Now()

	var bodyCopy []byte
	if c.Request().Body() != nil {
		bodyCopy = make([]byte, len(c.Request().Body()))
		copy(bodyCopy, c.Request().Body())
	}

	err := c.Next()

	respCopy := c.Response().Body()

	log := presentations.AuditLog{
		ID:        uuid.NewString(),
		UserID:    common.GetUserCtx(c.UserContext()).UserID,
		RequestID: c.Locals("requestid").(string),
		IPAddress: c.IP(),
		Path:      c.OriginalURL(),
		Payload:   json.RawMessage(bodyCopy),
		Response:  json.RawMessage(respCopy),
		CreatedAt: start,
		UpdatedAt: start,
		CreatedBy: common.GetUserCtx(c.UserContext()).Username,
	}

	if len(bodyCopy) == 0 {
		log.Payload = json.RawMessage(`{}`)
	}
	if len(respCopy) == 0 {
		log.Response = json.RawMessage(`{}`)
	}

	_ = m.business.AuditLog.Create(c.Context(), log)

	return err
}
