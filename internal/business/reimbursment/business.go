package reimbursment

import (
	"context"
	"payslips/internal/common"
	"payslips/internal/entity"
	"payslips/internal/presentations"
	"payslips/internal/repositories"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type Contract interface {
	Create(ctx context.Context, payload entity.ReimbursementCreate) (*presentations.Reimbursement, error)
	Update(ctx context.Context, payload entity.ReimbursementUpdate) (*presentations.Reimbursement, error)
	Detail(ctx context.Context, Id string) (*presentations.Reimbursement, error)
}

type business struct {
	repo *repositories.Repository
}

func NewBusiness(repo *repositories.Repository) Contract {
	return &business{
		repo: repo,
	}
}

func (b *business) Create(ctx context.Context, payload entity.ReimbursementCreate) (*presentations.Reimbursement, error) {

	userCtx := common.GetUserCtx(ctx)

	amount, _ := strconv.Atoi(payload.Amount)
	data := presentations.Reimbursement{
		ID:          uuid.NewString(),
		UserID:      userCtx.UserID,
		Amount:      amount,
		Description: payload.Description,
		Attachment:  payload.Attachment,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		CreatedBy:   userCtx.Username,
	}
	err := b.repo.Reimbursement.Create(ctx, data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (b *business) Update(ctx context.Context, payload entity.ReimbursementUpdate) (*presentations.Reimbursement, error) {

	userCtx := common.GetUserCtx(ctx)

	reimburs, err := b.repo.Reimbursement.Detail(ctx, payload.Id)
	if err != nil {
		return nil, err
	}

	if reimburs.UserID != userCtx.UserID {
		return nil, common.ErrForbidden
	}

	amount, _ := strconv.Atoi(payload.Amount)
	data := presentations.Reimbursement{
		ID:          payload.Id,
		Amount:      amount,
		Description: payload.Description,
		Attachment:  payload.Attachment,
		UpdatedAt:   time.Now(),
		UpdatedBy:   userCtx.Username,
	}
	err = b.repo.Reimbursement.Update(ctx, data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (b *business) Detail(ctx context.Context, Id string) (*presentations.Reimbursement, error) {

	userCtx := common.GetUserCtx(ctx)

	reimburs, err := b.repo.Reimbursement.Detail(ctx, Id)
	if err != nil {
		return nil, err
	}

	if reimburs.UserID != userCtx.UserID {
		return nil, common.ErrForbidden
	}

	return reimburs, nil
}
