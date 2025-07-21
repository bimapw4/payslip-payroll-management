package reimbursment

import (
	"context"
	"payslips/internal/presentations"
	"time"
)

func (r *repo) Create(ctx context.Context, input presentations.Reimbursement) error {

	query := `
    INSERT INTO reimbursment (
        id, user_id, amount, description, attachment, created_at, updated_at, created_by, updated_by
    ) VALUES (
        :id, :user_id, :amount, :description, :attachment, :created_at, :updated_at, :created_by, :updated_by
    )`

	_, err := r.db.NamedExecContext(ctx, query, map[string]interface{}{
		"id":          input.ID,
		"user_id":     input.UserID,
		"amount":      input.Amount,
		"description": input.Description,
		"attachment":  input.Attachment,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"created_by":  input.CreatedBy,
		"updated_by":  input.UpdatedBy,
	})

	return err
}
