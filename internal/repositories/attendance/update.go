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
