package overtime

import (
	"context"
	"payslips/internal/presentations"
)

func (r *repo) FindByPayrollID(ctx context.Context, userID, payrollID string) ([]presentations.Overtime, error) {
	var (
		result = []presentations.Overtime{}
	)

	query := `SELECT * FROM overtime where payroll_id=:payroll_id and user_id=:user_id`

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
