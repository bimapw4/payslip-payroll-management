package presentations

import (
	"payslips/internal/common"
	"time"
)

const (
	ErrAttendanceNotExist     = common.Error("err attendance not exist")
	ErrAttendanceAlreadyExist = common.Error("err attendance already exist")
)

type Attendance struct {
	ID        string     `db:"id" json:"id"`
	UserID    string     `db:"user_id" json:"user_id"`
	CheckIn   time.Time  `db:"check_in" json:"check_in"`
	CheckOut  *time.Time `db:"check_out" json:"check_out,omitempty"`
	PayrollID *string    `db:"payroll_id" json:"payroll_id,omitempty"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`
	CreatedBy string     `db:"created_by" json:"created_by"`
	UpdatedBy string     `db:"updated_by" json:"updated_by"`
}
