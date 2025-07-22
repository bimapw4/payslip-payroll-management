package payroll

import (
	"context"
	"payslips/internal/presentations"
	"time"
)

func (r *repo) Update(ctx context.Context, payload presentations.Payroll) error {

	query := `
		update payroll set 
			period_start=:period_start, 
			period_end=:period_end, 
			updated_by=:updated_by, 
			updated_at=:updated_at 
		where id=:id
   	`

	_, err := r.db.NamedExecContext(ctx, query, map[string]interface{}{
		"id":           payload.ID,
		"period_start": payload.PeriodStart,
		"period_end":   payload.PeriodEnd,
		"updated_at":   time.Now(),
		"updated_by":   payload.UpdatedBy,
	})

	return err
}

func (r *repo) UpdatePayroll(ctx context.Context, payload presentations.Payroll) error {

	query := `
		update payroll set 
			run_payroll=:run_payroll, 
			updated_by=:updated_by, 
			updated_at=:updated_at 
		where id=:id
   	`

	_, err := r.db.NamedExecContext(ctx, query, map[string]interface{}{
		"id":          payload.ID,
		"run_payroll": payload.RunPayroll,
		"updated_at":  time.Now(),
		"updated_by":  payload.UpdatedBy,
	})

	return err
}
