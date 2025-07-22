package overtime

import (
	"context"
	"payslips/internal/common"
	"payslips/internal/entity"
	"payslips/internal/presentations"
	"payslips/internal/repositories"
	"payslips/pkg/meta"
	"time"

	"github.com/google/uuid"
)

type Contract interface {
	Overtime(ctx context.Context, payload entity.Overtime) (*presentations.Overtime, error)
	Update(ctx context.Context, payload entity.Overtime, id string) (*presentations.Overtime, error)
	List(ctx context.Context, m *meta.Params) ([]presentations.Overtime, error)
}

type business struct {
	repo *repositories.Repository
}

func NewBusiness(repo *repositories.Repository) Contract {
	return &business{
		repo: repo,
	}
}

func (b *business) Overtime(ctx context.Context, payload entity.Overtime) (*presentations.Overtime, error) {

	if payload.GetDuration() > 3 {
		return nil, common.Error("overtime cannot more than 3 hours")
	}

	if payload.GetDuration() < 0 {
		return nil, common.Error("overtime cannot less than 0 hours")
	}

	userctx := common.GetUserCtx(ctx)

	data := presentations.Overtime{
		ID:        uuid.NewString(),
		UserID:    userctx.UserID,
		StartTime: payload.StartTime,
		EndTime:   payload.EndTime,
		Duration:  payload.GetDuration(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		CreatedBy: userctx.Username,
	}

	err := b.repo.Overtime.Create(ctx, data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (b *business) Update(ctx context.Context, payload entity.Overtime, id string) (*presentations.Overtime, error) {

	if payload.GetDuration() > 3 {
		return nil, common.Error("overtime cannot more than 3 hours")
	}

	if payload.GetDuration() < 0 {
		return nil, common.Error("overtime cannot less than 0 hours")
	}

	userctx := common.GetUserCtx(ctx)

	ext, err := b.repo.Overtime.Detail(ctx, id)
	if err != nil {
		return nil, err
	}

	if ext.UserID != userctx.UserID {
		return nil, common.ErrForbidden
	}

	data := presentations.Overtime{
		ID:        ext.ID,
		UserID:    userctx.UserID,
		StartTime: payload.StartTime,
		EndTime:   payload.EndTime,
		Duration:  payload.GetDuration(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		CreatedBy: userctx.Username,
	}

	err = b.repo.Overtime.Update(ctx, data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (b *business) List(ctx context.Context, m *meta.Params) ([]presentations.Overtime, error) {
	userctx := common.GetUserCtx(ctx)
	return b.repo.Overtime.List(ctx, m, userctx.UserID)
}
