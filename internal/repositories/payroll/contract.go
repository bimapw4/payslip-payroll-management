package payroll

import (
	"context"
	"payslips/internal/presentations"
	"payslips/pkg/meta"
)

type Payroll interface {
	Create(ctx context.Context, input presentations.Payroll) error
	Update(ctx context.Context, payload presentations.Payroll) error
	Detail(ctx context.Context, id string) (*presentations.Payroll, error)
	List(ctx context.Context, m *meta.Params, userID string) ([]presentations.Payroll, error)
	UpdatePayroll(ctx context.Context, payload presentations.Payroll) error
}
