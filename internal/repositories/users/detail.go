package users

import (
	"context"
	"payslips/internal/presentations"
)

func (r *repo) Detail(ctx context.Context, id string) (*presentations.Users, error) {
	var (
		result = presentations.Users{}
	)

	query := `SELECT * FROM users where id=:id`

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

func (r *repo) GetUserByUsername(ctx context.Context, username string) (*presentations.Users, error) {
	var (
		result = presentations.Users{}
	)

	query := `SELECT * FROM users where username=:username`

	args := map[string]interface{}{
		"username": username,
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

func (r *repo) GetAllUsers(ctx context.Context) ([]presentations.Users, error) {
	var (
		result = []presentations.Users{}
	)

	query := `SELECT id, salary FROM users`

	args := map[string]interface{}{}

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
