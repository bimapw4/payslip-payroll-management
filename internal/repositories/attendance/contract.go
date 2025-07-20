package attendance

import (
	"context"
	"payslips/internal/presentations"
	"time"
)

type Attendance interface {
	Create(ctx context.Context, input presentations.Attendance) error
	Detail(ctx context.Context, id string) (*presentations.Attendance, error)
	GetCheckinByDate(ctx context.Context, user_id string, date time.Time) (*presentations.Attendance, error)
	Update(ctx context.Context, payload presentations.Attendance) error
}
