package attendance

import (
	"context"
	"payslips/internal/presentations"
	"payslips/pkg/meta"
	"time"
)

type Attendance interface {
	Create(ctx context.Context, input presentations.Attendance) error
	Detail(ctx context.Context, id string) (*presentations.Attendance, error)
	List(ctx context.Context, m *meta.Params, userID string) ([]presentations.Attendance, error)
	GetCheckinByDate(ctx context.Context, user_id string, date time.Time) (*presentations.Attendance, error)
	Update(ctx context.Context, payload presentations.Attendance) error
	UpdatePayrollID(ctx context.Context, payrollID string, updatedBy string, start, end time.Time) error
	FindByPayrollID(ctx context.Context, userID, payrollID string) ([]presentations.Attendance, error)
}
