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
