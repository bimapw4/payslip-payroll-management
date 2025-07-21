package presentations

import "time"

type Payroll struct {
	ID          string    `db:"id" json:"id"`
	PeriodStart time.Time `db:"period_start" json:"period_start"`
	PeriodEnd   time.Time `db:"period_end" json:"period_end"`
	RunPayroll  bool      `db:"run_payroll" json:"run_payroll"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
	CreatedBy   string    `db:"created_by" json:"created_by"`
	UpdatedBy   string    `db:"updated_by" json:"updated_by"`
}

type PayslipResponse struct {
	PayrollID      string              `json:"payroll_id"`
	Period         Period              `json:"period"`
	Attendance     AttendanceBreakdown `json:"attendance"`
	Overtime       OvertimeBreakdown   `json:"overtime"`
	Reimbursements []Reimbursement     `json:"reimbursements"`
	TotalTakeHome  int                 `json:"total_take_home"`
}

type Period struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

type AttendanceBreakdown struct {
	WorkingDays    int `json:"working_days"`
	PresentDays    int `json:"present_days"`
	AbsentDays     int `json:"absent_days"`
	ProratedSalary int `json:"prorated_salary"`
}

type OvertimeBreakdown struct {
	TotalHours  float64 `json:"total_hours"`
	OvertimePay int     `json:"overtime_pay"`
}
