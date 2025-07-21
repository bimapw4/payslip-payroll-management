package reimbursment

import (
	"context"
	"payslips/internal/presentations"
	"time"
)

func (r *repo) Update(ctx context.Context, payload presentations.Reimbursement) error {

	query := `
		update reimbursement set 
			amount=:amount, 
			description=:description, 
			attachment=:attachment, 
			updated_by=:updated_by, 
			updated_at=:updated_at 
		where id=:id
   	`

	_, err := r.db.NamedExecContext(ctx, query, map[string]interface{}{
		"id":          payload.ID,
		"amount":      payload.Amount,
		"description": payload.Description,
		"attachment":  payload.Attachment,
		"updated_at":  time.Now(),
		"updated_by":  payload.UpdatedBy,
	})

	return err
}
