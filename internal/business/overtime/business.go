package overtime

import (
	"context"
	"payslips/internal/common"
	"payslips/internal/entity"
	"payslips/internal/presentations"
	"payslips/internal/repositories"
	"time"

	"github.com/google/uuid"
)

type Contract interface {
	Overtime(ctx context.Context, payload entity.Overtime) error
}

type business struct {
	repo *repositories.Repository
}

func NewBusiness(repo *repositories.Repository) Contract {
	return &business{
		repo: repo,
	}
}

func (b *business) Overtime(ctx context.Context, payload entity.Overtime) error {

	if payload.GetDuration() > 3 {
		return common.Error("overtime cannot more than 3 hours")
	}

	userctx := common.GetUserCtx(ctx)

	err := b.repo.Overtime.Create(ctx, presentations.Overtime{
		ID:        uuid.NewString(),
		UserID:    userctx.UserID,
		StartTime: payload.StartTime,
		EndTime:   payload.EndTime,
		Duration:  payload.GetDuration(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		CreatedBy: userctx.Username,
	})
	if err != nil {
		return err
	}

	return nil
}
