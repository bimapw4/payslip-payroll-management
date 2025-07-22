package auditlog

import (
	"context"
	"fmt"
	"payslips/internal/presentations"
	"payslips/pkg/meta"
	"strings"
)

func (r *repo) List(ctx context.Context, m *meta.Params) ([]presentations.AuditLog, error) {
	var logs []presentations.AuditLog

	q, err := meta.Parse(m)
	if err != nil {
		return nil, err
	}

	query := `SELECT * FROM audit_log ORDER BY created_at DESC OFFSET :offset LIMIT :limit`

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

	return logs, nil
}
