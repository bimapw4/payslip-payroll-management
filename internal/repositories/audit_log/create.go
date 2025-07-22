package auditlog

import (
	"context"
	"payslips/internal/presentations"
)

func (r *repo) Create(ctx context.Context, input presentations.AuditLog) error {
	query := `
	INSERT INTO audit_log (
		id, user_id, request_id, ip_address, path_invoice, payload, response, created_at, updated_at, created_by
	) VALUES (
		:id, :user_id, :request_id, :ip_address, :path_invoice, :payload, :response, :created_at, :updated_at, :created_by
	)
	`

	_, err := r.db.NamedExecContext(ctx, query, map[string]interface{}{
		"id":           input.ID,
		"user_id":      input.UserID,
		"request_id":   input.RequestID,
		"ip_address":   input.IPAddress,
		"path_invoice": input.Path,
		"payload":      input.Payload,
		"response":     input.Response,
		"created_at":   input.CreatedAt,
		"updated_at":   input.UpdatedAt,
		"created_by":   input.CreatedBy,
	})
	return err
}
