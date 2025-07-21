package payroll

import (
	"context"
	"payslips/internal/presentations"
)

type Payroll interface {
	Create(ctx context.Context, input presentations.Payroll) error
	Update(ctx context.Context, payload presentations.Payroll) error
	Detail(ctx context.Context, id string) (*presentations.Payroll, error)
}
