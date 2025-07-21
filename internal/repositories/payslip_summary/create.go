package payslipsummary

import (
	"context"
	"payslips/internal/presentations"
	"time"
)

func (r *repo) Create(ctx context.Context, input presentations.PayslipSummary) error {

	query := `
    INSERT INTO payslip_summary (
			id, payroll_id, user_id, base_salary, prorated_salary, 
			overtime_pay, reimbursement_total, take_home_pay, created_at, updated_at, created_by, updated_by
	) VALUES (
			:id, :payroll_id, :user_id, :base_salary, :prorated_salary, 
			:overtime_pay, :reimbursement_total, :take_home_pay, :created_at, :updated_at, :created_by, :updated_by
	)`

	_, err := r.db.NamedExecContext(ctx, query, map[string]interface{}{
		"id":                  input.ID,
		"payroll_id":          input.PayrollID,
		"user_id":             input.UserID,
		"base_salary":         input.BaseSalary,
		"prorated_salary":     input.ProratedSalary,
		"overtime_pay":        input.OvertimePay,
		"reimbursement_total": input.ReimbursementTotal,
		"take_home_pay":       input.TakeHomePay,
		"created_at":          time.Now(),
		"updated_at":          time.Now(),
		"created_by":          input.CreatedBy,
		"updated_by":          input.UpdatedBy,
	})

	return err
}
