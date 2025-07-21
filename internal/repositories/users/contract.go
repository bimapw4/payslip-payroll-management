package users

import (
	"context"
	"payslips/internal/presentations"
)

type Users interface {
	Create(ctx context.Context, input presentations.Users) error
	Detail(ctx context.Context, id string) (*presentations.Users, error)
	UpdatePassword(ctx context.Context, userID, password, updatedBy string) error
	Update(ctx context.Context, payload presentations.Users) error
	DeleteUser(ctx context.Context, userID, updatedBy string) error
	GetUserByUsername(ctx context.Context, username string) (*presentations.Users, error)
	GetAllUsers(ctx context.Context) ([]presentations.Users, error)
}
