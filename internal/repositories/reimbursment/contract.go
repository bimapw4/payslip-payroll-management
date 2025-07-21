package reimbursment

import (
	"context"
	"payslips/internal/presentations"
)

type Reimbursment interface {
	Create(ctx context.Context, input presentations.Reimbursement) error
	Detail(ctx context.Context, id string) (*presentations.Reimbursement, error)
	Update(ctx context.Context, payload presentations.Reimbursement) error
}
