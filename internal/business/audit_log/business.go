package auditlog

import (
	"context"
	"fmt"
	"payslips/internal/presentations"
	"payslips/internal/repositories"
	"payslips/pkg/meta"
)

type Contract interface {
	Create(ctx context.Context, input presentations.AuditLog) error
	List(ctx context.Context, m *meta.Params) ([]presentations.AuditLog, error)
}

type business struct {
	repo *repositories.Repository
}

func NewBusiness(repo *repositories.Repository) Contract {
	return &business{
		repo: repo,
	}
}

func (b *business) Create(ctx context.Context, input presentations.AuditLog) error {
	err := b.repo.AuditLog.Create(ctx, input)
	if err != nil {
		return err
	}

	fmt.Println("saada == ", err)
	return nil
}

func (b *business) List(ctx context.Context, m *meta.Params) ([]presentations.AuditLog, error) {
	res, err := b.repo.AuditLog.List(ctx, m)
	if err != nil {
		return nil, err
	}
	return res, nil
}
