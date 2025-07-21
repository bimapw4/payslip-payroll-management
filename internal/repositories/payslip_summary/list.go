package payslipsummary

import (
	"context"
	"fmt"
	"payslips/internal/presentations"
	"payslips/pkg/meta"
	"strings"
)

func (r *repo) List(ctx context.Context, m *meta.Params, payrollID string) ([]presentations.PayslipSummary, error) {
	var (
		result = []presentations.PayslipSummary{}
	)

	q, err := meta.Parse(m)
	if err != nil {
		return nil, err
	}
	query := `SELECT * FROM payslip_summary where 1=1 ORDER BY created_at DESC OFFSET :offset LIMIT :limit`

	query = strings.Replace(
		query,
		" ORDER BY created_at DESC ",
		fmt.Sprintf(" ORDER BY %s %s ", q.OrderBy, q.OrderDirection),
		1,
	)

	if m.SearchBy != "" {
		query = strings.ReplaceAll(query, "1=1", fmt.Sprintf("%v='%v'", m.SearchBy, m.Search))
	}

	args := map[string]interface{}{
		"offset": q.Offset,
		"limit":  q.Limit,
	}

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return nil, r.translateError(err)
	}

	err = stmt.SelectContext(ctx, &result, args)
	if err != nil {
		return nil, r.translateError(err)
	}

	total, totalTakeHome, err := r.Count(ctx, payrollID)
	if err != nil {
		return nil, r.translateError(err)
	}

	m.TotalItems = total
	m.TotalTakeHomePay = totalTakeHome

	return result, nil
}

func (r *repo) Count(ctx context.Context, payrollID string) (int, int, error) {
	var count int
	var totalTakeHome int

	query := `SELECT COUNT(*), COALESCE(SUM(take_home_pay), 0) FROM payslip_summary WHERE payroll_id = :payroll_id`

	args := map[string]interface{}{
		"payroll_id": payrollID,
	}

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return 0, 0, r.translateError(err)
	}

	err = stmt.QueryRowxContext(ctx, args).Scan(&count, &totalTakeHome)
	if err != nil {
		return 0, 0, r.translateError(err)
	}

	return count, totalTakeHome, nil
}
