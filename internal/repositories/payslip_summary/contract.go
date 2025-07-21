package payslipsummary

import (
	"context"
	"payslips/internal/presentations"
	"payslips/pkg/meta"
)

type PayslipSummary interface {
	Create(ctx context.Context, input presentations.PayslipSummary) error
	List(ctx context.Context, m *meta.Params, payrollID string) ([]presentations.PayslipSummary, error)
}
