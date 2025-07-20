package payroll

import (
	"context"
	"payslips/internal/presentations"
	"time"
)

func (r *repo) Create(ctx context.Context, input presentations.Payroll) error {

	query := `
    INSERT INTO payroll (
        id, period_start, period_end, created_at, updated_at, created_by, updated_by
    ) VALUES (
        :id, :period_start, :period_end, :created_at, :updated_at, :created_by, :updated_by
    )`

	_, err := r.db.NamedExecContext(ctx, query, map[string]interface{}{
		"id":           input.ID,
		"period_start": input.PeriodStart,
		"period_end":   input.PeriodEnd,
		"created_at":   time.Now(),
		"updated_at":   time.Now(),
		"created_by":   input.CreatedBy,
		"updated_by":   input.UpdatedBy,
	})

	return err
}
