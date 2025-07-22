package auditlog

import (
	"context"
	"payslips/internal/presentations"
	"payslips/pkg/meta"
)

type AuditLog interface {
	Create(ctx context.Context, input presentations.AuditLog) error
	List(ctx context.Context, m *meta.Params) ([]presentations.AuditLog, error)
}
