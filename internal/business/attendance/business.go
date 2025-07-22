package attendance

import (
	"context"
	"payslips/internal/common"
	"payslips/internal/consts"
	"payslips/internal/entity"
	"payslips/internal/presentations"
	"payslips/internal/repositories"
	"payslips/pkg/meta"
	"time"

	"github.com/google/uuid"
)

type Contract interface {
	Attendance(ctx context.Context, payload entity.AttendanceInput) error
	List(ctx context.Context, m *meta.Params) ([]presentations.Attendance, error)
}

type business struct {
	repo *repositories.Repository
}

func NewBusiness(repo *repositories.Repository) Contract {
	return &business{
		repo: repo,
	}
}

func (b *business) Attendance(ctx context.Context, payload entity.AttendanceInput) error {

	switch payload.Datetime.Weekday() {
	case time.Saturday, time.Sunday:
		return common.Error("cannot submit attendance on weekend")
	}

	attendanceType := map[string]func(ctx context.Context, payload entity.AttendanceInput) error{
		consts.AttendanceCheckin: func(ctx context.Context, payload entity.AttendanceInput) error {
			return b.checkIn(ctx, payload)
		},
		consts.AttendanceCheckout: func(ctx context.Context, payload entity.AttendanceInput) error {
			return b.checkOut(ctx, payload)
		},
	}

	fn, ok := attendanceType[payload.Type]
	if !ok {
		return common.ErrBadRequest
	}

	err := fn(ctx, payload)
	if err != nil {
		return err
	}

	return nil
}

func (b *business) checkIn(ctx context.Context, payload entity.AttendanceInput) error {

	user := common.GetUserCtx(ctx)

	exist, _ := b.repo.Attendance.GetCheckinByDate(ctx, user.UserID, payload.Datetime)
	if exist != nil {
		return presentations.ErrAttendanceAlreadyExist
	}

	err := b.repo.Attendance.Create(ctx, presentations.Attendance{
		ID:        uuid.NewString(),
		UserID:    user.UserID,
		CheckIn:   payload.Datetime,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		CreatedBy: user.Username,
	})
	if err != nil {
		return err
	}

	return nil
}

func (b *business) checkOut(ctx context.Context, payload entity.AttendanceInput) error {

	user := common.GetUserCtx(ctx)

	exist, _ := b.repo.Attendance.GetCheckinByDate(ctx, user.UserID, payload.Datetime)
	if exist == nil {
		return presentations.ErrAttendanceNotExist
	}

	if exist.CheckOut != nil {
		return presentations.ErrAttendanceAlreadyExist
	}

	err := b.repo.Attendance.Update(ctx, presentations.Attendance{
		ID:        exist.ID,
		UserID:    user.UserID,
		CheckOut:  &payload.Datetime,
		UpdatedBy: user.Username,
	})
	if err != nil {
		return err
	}

	return nil
}

func (b *business) List(ctx context.Context, m *meta.Params) ([]presentations.Attendance, error) {

	userctx := common.GetUserCtx(ctx)

	return b.repo.Attendance.List(ctx, m, userctx.UserID)
}
