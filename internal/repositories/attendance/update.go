package attendance

import (
	"context"
	"payslips/internal/presentations"
	"time"
)

func (r *repo) Update(ctx context.Context, payload presentations.Attendance) error {

	query := `
		update attendance set 
			check_out=:check_out, 
			updated_by=:updated_by, 
			updated_at=:updated_at 
		where id=:id
   	`

	_, err := r.db.NamedExecContext(ctx, query, map[string]interface{}{
		"id":         payload.ID,
		"check_out":  payload.CheckOut,
		"updated_at": time.Now(),
		"updated_by": payload.UpdatedBy,
	})

	return err
}

func (r *repo) UpdatePayrollID(ctx context.Context, payrollID string, updatedBy string, start, end time.Time) error {

	query := `
		update attendance set 
			payroll_id=:payroll_id, 
			updated_by=:updated_by, 
			updated_at=:updated_at 
		where created_at BETWEEN :start AND :end AND payroll_id IS NULL
   	`

	_, err := r.db.NamedExecContext(ctx, query, map[string]interface{}{
		"start":      start,
		"end":        end,
		"payroll_id": payrollID,
		"updated_at": time.Now(),
		"updated_by": updatedBy,
	})

	return err
}
