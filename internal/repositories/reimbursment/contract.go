package reimbursment

import (
	"context"
	"payslips/internal/presentations"
	"payslips/pkg/meta"
	"time"
)

type Reimbursment interface {
	Create(ctx context.Context, input presentations.Reimbursement) error
	Detail(ctx context.Context, id string) (*presentations.Reimbursement, error)
	Update(ctx context.Context, payload presentations.Reimbursement) error
	UpdatePayrollID(ctx context.Context, payrollID string, updatedBy string, start, end time.Time) error
	FindByPayrollID(ctx context.Context, userID, payrollID string) ([]presentations.Reimbursement, error)
	List(ctx context.Context, m *meta.Params, userID string) ([]presentations.Reimbursement, error)
}
