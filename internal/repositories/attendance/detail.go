package attendance

import (
	"context"
	"payslips/internal/presentations"
	"time"
)

func (r *repo) Detail(ctx context.Context, id string) (*presentations.Attendance, error) {
	var (
		result = presentations.Attendance{}
	)

	query := `SELECT * FROM attendance where id=:id`

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

func (r *repo) GetCheckinByDate(ctx context.Context, user_id string, date time.Time) (*presentations.Attendance, error) {
	var (
		result = presentations.Attendance{}
	)

	query := `SELECT * FROM attendance where user_id=:user_id and date(check_in)=date(:check_in)`

	args := map[string]interface{}{
		"check_in": date,
		"user_id":  user_id,
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

func (r *repo) FindByPayrollID(ctx context.Context, userID, payrollID string) ([]presentations.Attendance, error) {
	var (
		result = []presentations.Attendance{}
	)

	query := `SELECT * FROM attendance where payroll_id=:payroll_id and user_id=:user_id`

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
