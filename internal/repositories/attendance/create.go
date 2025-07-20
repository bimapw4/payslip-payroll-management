package attendance

import (
	"context"
	"payslips/internal/presentations"
)

func (r *repo) Create(ctx context.Context, input presentations.Attendance) error {

	query := `
    INSERT INTO attendance (
        id, user_id, check_in, created_at, updated_at, created_by, updated_by
    ) VALUES (
        :id, :user_id, :check_in, :created_at, :updated_at, :created_by, :updated_by
    )`

	_, err := r.db.NamedExecContext(ctx, query, map[string]interface{}{
		"id":         input.ID,
		"user_id":    input.UserID,
		"check_in":   input.CheckIn,
		"created_at": input.CreatedAt,
		"updated_at": input.UpdatedAt,
		"created_by": input.CreatedBy,
		"updated_by": input.UpdatedBy,
	})

	return err
}
