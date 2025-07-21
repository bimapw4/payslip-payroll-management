package payroll

import (
	"context"
	"payslips/internal/presentations"
)

func (r *repo) Detail(ctx context.Context, id string) (*presentations.Payroll, error) {
	var (
		result = presentations.Payroll{}
	)

	query := `SELECT * FROM payroll where id=:id`

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
