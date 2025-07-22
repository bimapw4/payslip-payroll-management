package overtime

import (
	"context"
	"payslips/internal/presentations"
	"payslips/pkg/meta"
	"time"
)

type Overtime interface {
	Create(ctx context.Context, input presentations.Overtime) error
	List(ctx context.Context, m *meta.Params, userID string) ([]presentations.Overtime, error)
	Update(ctx context.Context, payload presentations.Overtime) error
	Detail(ctx context.Context, id string) (*presentations.Overtime, error)
	UpdatePayrollID(ctx context.Context, payrollID string, updatedBy string, start, end time.Time) error
	FindByPayrollID(ctx context.Context, userID, payrollID string) ([]presentations.Overtime, error)
	GetOvertimeByDate(ctx context.Context, user_id string, date time.Time) (*presentations.Attendance, error)
}
