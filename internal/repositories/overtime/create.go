package overtime

import (
	"context"
	"payslips/internal/presentations"
)

func (r *repo) Create(ctx context.Context, input presentations.Overtime) error {

	query := `
    INSERT INTO overtime (
        id, user_id, start_time, end_time, duration, created_at,updated_at, created_by, updated_by
    ) VALUES (
        :id, :user_id, :start_time, :end_time, :duration, :created_at, :updated_at, :created_by, :updated_by
    )`

	_, err := r.db.NamedExecContext(ctx, query, map[string]interface{}{
		"id":         input.ID,
		"user_id":    input.UserID,
		"start_time": input.StartTime,
		"end_time":   input.EndTime,
		"duration":   input.Duration,
		"created_at": input.CreatedAt,
		"updated_at": input.UpdatedAt,
		"created_by": input.CreatedBy,
		"updated_by": input.UpdatedBy,
	})

	return err
}
