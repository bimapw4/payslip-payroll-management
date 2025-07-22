package overtime

import (
	"context"
	"payslips/internal/presentations"
	"time"
)

func (r *repo) Update(ctx context.Context, payload presentations.Overtime) error {

	query := `
		update overtime set 
			start_time=:start_time, 
			end_time=:end_time, 
			duration=:duration,
			updated_by=:updated_by, 
			updated_at=:updated_at 
		where id = :id
   	`

	_, err := r.db.NamedExecContext(ctx, query, map[string]interface{}{
		"id":         payload.ID,
		"start_time": payload.StartTime,
		"end_time":   payload.EndTime,
		"duration":   payload.Duration,
		"updated_at": time.Now(),
		"updated_by": payload.UpdatedBy,
	})

	return err
}

func (r *repo) UpdatePayrollID(ctx context.Context, payrollID string, updatedBy string, start, end time.Time) error {

	query := `
		update overtime set 
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
