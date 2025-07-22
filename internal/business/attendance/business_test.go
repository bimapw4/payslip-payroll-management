package attendance_test

import (
	"context"
	"payslips/internal/business/attendance"
	"payslips/internal/common"
	"payslips/internal/consts"
	"payslips/internal/entity"
	"payslips/internal/presentations"
	"payslips/internal/repositories"
	attendancerep "payslips/internal/repositories/attendance"
	"payslips/internal/repositories/overtime"
	"payslips/internal/repositories/payroll"
	payslipsummary "payslips/internal/repositories/payslip_summary"
	"payslips/internal/repositories/reimbursment"
	"payslips/internal/repositories/users"
	"payslips/pkg/meta"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestBusiness_Attendance_CheckIn_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsers := users.NewMockUsers(ctrl)
	mockAttendance := attendancerep.NewMockAttendance(ctrl)
	mockOvertime := overtime.NewMockOvertime(ctrl)
	mockReimbursement := reimbursment.NewMockReimbursment(ctrl)
	mockPayslipSummary := payslipsummary.NewMockPayslipSummary(ctrl)
	mockPayroll := payroll.NewMockPayroll(ctrl)

	repo := &repositories.Repository{
		Users:          mockUsers,
		Attendance:     mockAttendance,
		Payroll:        mockPayroll,
		Overtime:       mockOvertime,
		Reimbursement:  mockReimbursement,
		PayslipSummary: mockPayslipSummary,
	}

	b := attendance.NewBusiness(repo)

	ctx := common.SetUserCtx(context.Background(), &entity.Claim{
		UserID:   "user-1",
		Username: "user1",
	})

	input := entity.AttendanceInput{
		Type:     consts.AttendanceCheckin,
		Datetime: time.Now(),
	}

	mockAttendance.EXPECT().
		GetCheckinByDate(ctx, "user-1", gomock.Any()).
		Return(nil, nil)

	mockAttendance.EXPECT().
		Create(ctx, gomock.AssignableToTypeOf(presentations.Attendance{})).
		Return(nil)

	err := b.Attendance(ctx, input)
	assert.NoError(t, err)
}

func TestBusiness_Attendance_CheckIn_Duplicate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsers := users.NewMockUsers(ctrl)
	mockAttendance := attendancerep.NewMockAttendance(ctrl)
	mockOvertime := overtime.NewMockOvertime(ctrl)
	mockReimbursement := reimbursment.NewMockReimbursment(ctrl)
	mockPayslipSummary := payslipsummary.NewMockPayslipSummary(ctrl)
	mockPayroll := payroll.NewMockPayroll(ctrl)

	repo := &repositories.Repository{
		Users:          mockUsers,
		Attendance:     mockAttendance,
		Payroll:        mockPayroll,
		Overtime:       mockOvertime,
		Reimbursement:  mockReimbursement,
		PayslipSummary: mockPayslipSummary,
	}

	b := attendance.NewBusiness(repo)

	ctx := common.SetUserCtx(context.Background(), &entity.Claim{
		UserID:   "user-1",
		Username: "user1",
	})

	now := time.Now()

	mockAttendance.EXPECT().
		GetCheckinByDate(ctx, "user-1", gomock.Any()).
		Return(&presentations.Attendance{
			ID:      "att-1",
			UserID:  "user-1",
			CheckIn: now,
		}, nil)

	input := entity.AttendanceInput{
		Type:     consts.AttendanceCheckin,
		Datetime: now,
	}

	err := b.Attendance(ctx, input)
	assert.ErrorIs(t, err, presentations.ErrAttendanceAlreadyExist)
}

func TestBusiness_Attendance_CheckOut_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsers := users.NewMockUsers(ctrl)
	mockAttendance := attendancerep.NewMockAttendance(ctrl)
	mockOvertime := overtime.NewMockOvertime(ctrl)
	mockReimbursement := reimbursment.NewMockReimbursment(ctrl)
	mockPayslipSummary := payslipsummary.NewMockPayslipSummary(ctrl)
	mockPayroll := payroll.NewMockPayroll(ctrl)

	repo := &repositories.Repository{
		Users:          mockUsers,
		Attendance:     mockAttendance,
		Payroll:        mockPayroll,
		Overtime:       mockOvertime,
		Reimbursement:  mockReimbursement,
		PayslipSummary: mockPayslipSummary,
	}

	b := attendance.NewBusiness(repo)

	ctx := common.SetUserCtx(context.Background(), &entity.Claim{
		UserID:   "user-1",
		Username: "user1",
	})

	now := time.Now()

	mockAttendance.EXPECT().
		GetCheckinByDate(ctx, "user-1", gomock.Any()).
		Return(&presentations.Attendance{
			ID:     "att-1",
			UserID: "user-1",
		}, nil)

	mockAttendance.EXPECT().
		Update(ctx, gomock.AssignableToTypeOf(presentations.Attendance{})).
		Return(nil)

	input := entity.AttendanceInput{
		Type:     consts.AttendanceCheckout,
		Datetime: now,
	}

	err := b.Attendance(ctx, input)
	assert.NoError(t, err)
}

func TestBusiness_Attendance_CheckOut_NotExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsers := users.NewMockUsers(ctrl)
	mockAttendance := attendancerep.NewMockAttendance(ctrl)
	mockOvertime := overtime.NewMockOvertime(ctrl)
	mockReimbursement := reimbursment.NewMockReimbursment(ctrl)
	mockPayslipSummary := payslipsummary.NewMockPayslipSummary(ctrl)
	mockPayroll := payroll.NewMockPayroll(ctrl)

	repo := &repositories.Repository{
		Users:          mockUsers,
		Attendance:     mockAttendance,
		Payroll:        mockPayroll,
		Overtime:       mockOvertime,
		Reimbursement:  mockReimbursement,
		PayslipSummary: mockPayslipSummary,
	}

	b := attendance.NewBusiness(repo)

	ctx := common.SetUserCtx(context.Background(), &entity.Claim{
		UserID:   "user-1",
		Username: "user1",
	})

	now := time.Now()

	mockAttendance.EXPECT().
		GetCheckinByDate(ctx, "user-1", gomock.Any()).
		Return(nil, nil)

	input := entity.AttendanceInput{
		Type:     consts.AttendanceCheckout,
		Datetime: now,
	}

	err := b.Attendance(ctx, input)
	assert.ErrorIs(t, err, presentations.ErrAttendanceNotExist)
}

func TestBusiness_Attendance_CheckOut_AlreadyExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsers := users.NewMockUsers(ctrl)
	mockAttendance := attendancerep.NewMockAttendance(ctrl)
	mockOvertime := overtime.NewMockOvertime(ctrl)
	mockReimbursement := reimbursment.NewMockReimbursment(ctrl)
	mockPayslipSummary := payslipsummary.NewMockPayslipSummary(ctrl)
	mockPayroll := payroll.NewMockPayroll(ctrl)

	repo := &repositories.Repository{
		Users:          mockUsers,
		Attendance:     mockAttendance,
		Payroll:        mockPayroll,
		Overtime:       mockOvertime,
		Reimbursement:  mockReimbursement,
		PayslipSummary: mockPayslipSummary,
	}

	b := attendance.NewBusiness(repo)

	ctx := common.SetUserCtx(context.Background(), &entity.Claim{
		UserID:   "user-1",
		Username: "user1",
	})

	now := time.Now()
	checkoutTime := time.Now()

	mockAttendance.EXPECT().
		GetCheckinByDate(ctx, "user-1", gomock.Any()).
		Return(&presentations.Attendance{
			ID:       "att-1",
			UserID:   "user-1",
			CheckOut: &checkoutTime,
		}, nil)

	input := entity.AttendanceInput{
		Type:     consts.AttendanceCheckout,
		Datetime: now,
	}

	err := b.Attendance(ctx, input)
	assert.ErrorIs(t, err, presentations.ErrAttendanceAlreadyExist)
}

func TestBusiness_Attendance_UnknownType(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsers := users.NewMockUsers(ctrl)
	mockAttendance := attendancerep.NewMockAttendance(ctrl)
	mockOvertime := overtime.NewMockOvertime(ctrl)
	mockReimbursement := reimbursment.NewMockReimbursment(ctrl)
	mockPayslipSummary := payslipsummary.NewMockPayslipSummary(ctrl)
	mockPayroll := payroll.NewMockPayroll(ctrl)

	repo := &repositories.Repository{
		Users:          mockUsers,
		Attendance:     mockAttendance,
		Payroll:        mockPayroll,
		Overtime:       mockOvertime,
		Reimbursement:  mockReimbursement,
		PayslipSummary: mockPayslipSummary,
	}

	b := attendance.NewBusiness(repo)

	ctx := common.SetUserCtx(context.Background(), &entity.Claim{
		UserID:   "user-1",
		Username: "user1",
	})

	input := entity.AttendanceInput{
		Type:     "unknown-type",
		Datetime: time.Now(),
	}

	err := b.Attendance(ctx, input)
	assert.ErrorIs(t, err, common.ErrBadRequest)
}

func TestBusiness_Attendance_List_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsers := users.NewMockUsers(ctrl)
	mockAttendance := attendancerep.NewMockAttendance(ctrl)
	mockOvertime := overtime.NewMockOvertime(ctrl)
	mockReimbursement := reimbursment.NewMockReimbursment(ctrl)
	mockPayslipSummary := payslipsummary.NewMockPayslipSummary(ctrl)
	mockPayroll := payroll.NewMockPayroll(ctrl)

	repo := &repositories.Repository{
		Users:          mockUsers,
		Attendance:     mockAttendance,
		Payroll:        mockPayroll,
		Overtime:       mockOvertime,
		Reimbursement:  mockReimbursement,
		PayslipSummary: mockPayslipSummary,
	}

	b := attendance.NewBusiness(repo)

	ctx := common.SetUserCtx(context.Background(), &entity.Claim{
		UserID:   "user-1",
		Username: "user1",
	})

	metaParams := &meta.Params{
		// Page:    1,
		// Limit:   10,
		// OrderBy: "created_at",
		// Sort:    "desc",
	}

	now := time.Now()
	end := now.Add(1 * time.Hour)
	payrollID := "payroll-123"
	expectedResult := []presentations.Attendance{
		{
			ID:        "att-123",
			UserID:    "user-1",
			CheckIn:   now.Add(-9 * time.Hour),
			CheckOut:  &end,
			PayrollID: &payrollID,
			CreatedAt: now.Add(-10 * time.Hour),
			UpdatedAt: now,
			CreatedBy: "user1",
			UpdatedBy: "user1",
		},
	}

	mockAttendance.EXPECT().
		List(ctx, metaParams, "user-1").
		Return(expectedResult, nil)

	result, err := b.List(ctx, metaParams)

	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
}
