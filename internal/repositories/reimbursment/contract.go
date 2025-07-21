package reimbursment

import (
	"context"
	"payslips/internal/presentations"
	"time"
)

type Reimbursment interface {
	Create(ctx context.Context, input presentations.Reimbursement) error
	Detail(ctx context.Context, id string) (*presentations.Reimbursement, error)
	Update(ctx context.Context, payload presentations.Reimbursement) error
	UpdatePayrollID(ctx context.Context, payrollID string, updatedBy string, start, end time.Time) error
}
