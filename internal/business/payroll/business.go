package payroll

import (
	"context"
	"payslips/internal/common"
	"payslips/internal/entity"
	"payslips/internal/presentations"
	"payslips/internal/repositories"

	"github.com/google/uuid"
)

type Contract interface {
	CreatePayroll(ctx context.Context, payload entity.Payroll) error
	RunningPayroll(ctx context.Context, payrollID string) error
}

type business struct {
	repo *repositories.Repository
}

func NewBusiness(repo *repositories.Repository) Contract {
	return &business{
		repo: repo,
	}
}

func (b *business) CreatePayroll(ctx context.Context, payload entity.Payroll) error {

	userctx := common.GetUserCtx(ctx)

	err := b.repo.Payroll.Create(ctx, presentations.Payroll{
		ID:          uuid.NewString(),
		PeriodStart: payload.PeriodStart,
		PeriodEnd:   payload.PeriodEnd,
		CreatedBy:   userctx.Username,
	})
	if err != nil {
		return err
	}

	return nil
}

func (b *business) RunningPayroll(ctx context.Context, payrollID string) error {
	userctx := common.GetUserCtx(ctx)

	err := b.repo.Payroll.Update(ctx, presentations.Payroll{
		ID:         payrollID,
		RunPayroll: true,
		UpdatedBy:  userctx.Username,
	})
	if err != nil {
		return err
	}

	return nil
}
