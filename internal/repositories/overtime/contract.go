package overtime

import (
	"context"
	"payslips/internal/presentations"
	"time"
)

type Overtime interface {
	Create(ctx context.Context, input presentations.Overtime) error
	UpdatePayrollID(ctx context.Context, payrollID string, updatedBy string, start, end time.Time) error
	FindByPayrollID(ctx context.Context, userID, payrollID string) ([]presentations.Overtime, error)
}
