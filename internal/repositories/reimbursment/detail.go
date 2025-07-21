package reimbursment

import (
	"context"
	"payslips/internal/presentations"
)

func (r *repo) Detail(ctx context.Context, id string) (*presentations.Reimbursement, error) {
	var (
		result = presentations.Reimbursement{}
	)

	query := `SELECT * FROM reimbursement where id=:id`

	args := map[string]interface{}{
		"id": id,
	}

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return nil, r.translateError(err)
	}

	err = stmt.GetContext(ctx, &result, args)
	if err != nil {
		return nil, r.translateError(err)
	}

	return &result, nil
}
