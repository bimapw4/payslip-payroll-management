package users

import (
	"context"
	"payslips/internal/presentations"
)

func (r *repo) Create(ctx context.Context, input presentations.Users) error {

	query := `
    INSERT INTO users (
        id, name, username, password, salary, 
        is_admin, is_active, created_at, updated_at, created_by, updated_by
    ) VALUES (
        :id, :name, :username, :password, :salary, 
        :is_admin, :is_active, :created_at, :updated_at, :created_by, :updated_by
    )`

	_, err := r.db.NamedExecContext(ctx, query, map[string]interface{}{
		"id":         input.ID,
		"name":       input.Name,
		"username":   input.Username,
		"password":   input.Password,
		"salary":     input.Salary,
		"is_admin":   input.IsAdmin,
		"is_active":  input.IsActive,
		"created_at": input.CreatedAt,
		"updated_at": input.UpdatedAt,
		"created_by": input.CreatedBy,
		"updated_by": input.UpdatedBy,
	})

	return err
}
