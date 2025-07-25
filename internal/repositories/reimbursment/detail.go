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

func (r *repo) FindByPayrollID(ctx context.Context, userID, payrollID string) ([]presentations.Reimbursement, error) {
	var (
		result = []presentations.Reimbursement{}
	)

	query := `SELECT * FROM reimbursement where payroll_id=:payroll_id and user_id=:user_id`

	args := map[string]interface{}{
		"payroll_id": payrollID,
		"user_id":    userID,
	}

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return nil, r.translateError(err)
	}

	err = stmt.SelectContext(ctx, &result, args)
	if err != nil {
		return nil, r.translateError(err)
	}

	return result, nil
}
