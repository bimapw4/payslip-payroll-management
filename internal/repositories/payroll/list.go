package payroll

import (
	"context"
	"fmt"
	"payslips/internal/presentations"
	"payslips/pkg/meta"
	"strings"
)

func (r *repo) List(ctx context.Context, m *meta.Params, userID string) ([]presentations.Payroll, error) {
	var logs []presentations.Payroll

	q, err := meta.Parse(m)
	if err != nil {
		return nil, err
	}

	query := `SELECT * FROM payroll ORDER BY created_at DESC OFFSET :offset LIMIT :limit`

	query = strings.Replace(
		query,
		" ORDER BY created_at DESC ",
		fmt.Sprintf(" ORDER BY %s %s ", q.OrderBy, q.OrderDirection),
		1,
	)

	args := map[string]interface{}{
		"offset": q.Offset,
		"limit":  q.Limit,
	}

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return nil, err
	}

	err = stmt.SelectContext(ctx, &logs, args)
	if err != nil {
		return nil, err
	}

	count, err := r.Count(ctx, userID)
	if err != nil {
		return nil, err
	}

	m.TotalItems = count

	return logs, nil
}

func (r *repo) Count(ctx context.Context, userID string) (int, error) {
	var count int

	query := `SELECT COUNT(*)FROM payroll`

	args := map[string]interface{}{}

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return 0, err
	}

	err = stmt.QueryRowxContext(ctx, args).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
