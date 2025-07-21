package presentations

import (
	"payslips/internal/common"
	"time"
)

const (
	ErrPayslipSummaryNotExist     = common.Error("err payslip summary not exist")
	ErrPayslipSummaryAlreadyExist = common.Error("err payslip summary already exist")
)

type PayslipSummary struct {
	ID                 string    `db:"id" json:"id"`
	PayrollID          string    `db:"payroll_id" json:"payroll_id"`
	UserID             string    `db:"user_id" json:"user_id"`
	BaseSalary         int       `db:"base_salary" json:"base_salary"`
	ProratedSalary     int       `db:"prorated_salary" json:"prorated_salary"`
	OvertimePay        int       `db:"overtime_pay" json:"overtime_pay"`
	ReimbursementTotal int       `db:"reimbursement_total" json:"reimbursement_total"`
	TakeHomePay        int       `db:"take_home_pay" json:"take_home_pay"`
	CreatedAt          time.Time `db:"created_at" json:"created_at"`
	UpdatedAt          time.Time `db:"updated_at" json:"updated_at"`
	CreatedBy          string    `db:"created_by" json:"created_by"`
	UpdatedBy          string    `db:"updated_by" json:"updated_by"`
}
