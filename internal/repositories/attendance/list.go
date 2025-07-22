package attendance

import (
	"context"
	"fmt"
	"payslips/internal/presentations"
	"payslips/pkg/meta"
	"strings"
)

func (r *repo) List(ctx context.Context, m *meta.Params, userID string) ([]presentations.Attendance, error) {
	var logs []presentations.Attendance

	q, err := meta.Parse(m)
	if err != nil {
		return nil, err
	}

	query := `SELECT * FROM attendance where user_id=:user_id ORDER BY created_at DESC OFFSET :offset LIMIT :limit`

	query = strings.Replace(
		query,
		" ORDER BY created_at DESC ",
		fmt.Sprintf(" ORDER BY %s %s ", q.OrderBy, q.OrderDirection),
		1,
	)

	args := map[string]interface{}{
		"offset":  q.Offset,
		"limit":   q.Limit,
		"user_id": userID,
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

	query := `SELECT COUNT(*)FROM payslip_summary WHERE user_id=:user_id`

	args := map[string]interface{}{
		"user_id": userID,
	}

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
