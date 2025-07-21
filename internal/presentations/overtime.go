package presentations

import (
	"payslips/internal/common"
	"time"
)

const (
	ErrOvertimeNotExist     = common.Error("err overtime not exist")
	ErrOvertimeAlreadyExist = common.Error("err overtime already exist")
)

type Overtime struct {
	ID        string    `db:"id" json:"id"`
	UserID    string    `db:"user_id" json:"user_id"`
	StartTime time.Time `db:"start_time" json:"start_time"`
	EndTime   time.Time `db:"end_time" json:"end_time"`
	Duration  float64   `db:"duration" json:"duration"`
	PayrollID *string   `db:"payroll_id" json:"payroll_id,omitempty"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	CreatedBy string    `db:"created_by" json:"created_by"`
	UpdatedBy string    `db:"updated_by" json:"updated_by"`
}

func SumOvertime(overtimes []Overtime) float64 {
	var total float64
	for _, ot := range overtimes {
		total += ot.Duration
	}
	return total
}
