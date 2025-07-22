package overtime_test

import (
	"context"
	overtimebus "payslips/internal/business/overtime"
	"payslips/internal/common"
	"payslips/internal/entity"
	"payslips/internal/presentations"
	"payslips/internal/repositories"
	"payslips/internal/repositories/attendance"
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

func TestBusiness_Overtime_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsers := users.NewMockUsers(ctrl)
	mockAttendance := attendance.NewMockAttendance(ctrl)
	mockOvertime := overtime.NewMockOvertime(ctrl)
	mockReimbursement := reimbursment.NewMockReimbursment(ctrl)
	mockPayslipSummary := payslipsummary.NewMockPayslipSummary(ctrl)
	mockPayroll := payroll.NewMockPayroll(ctrl)

	mockrepo := &repositories.Repository{
		Users:          mockUsers,
		Attendance:     mockAttendance,
		Payroll:        mockPayroll,
		Overtime:       mockOvertime,
		Reimbursement:  mockReimbursement,
		PayslipSummary: mockPayslipSummary,
	}

	b := overtimebus.NewBusiness(mockrepo)

	ctx := common.SetUserCtx(context.Background(), &entity.Claim{
		UserID:   "user-1",
		Username: "user1",
	})

	start := time.Now()
	end := start.Add(2 * time.Hour)

	input := entity.Overtime{
		StartTime: start,
		EndTime:   end,
	}

	mockOvertime.EXPECT().
		GetOvertimeByDate(ctx, "user-1", start).
		Return(nil, nil)

	mockOvertime.EXPECT().
		Create(ctx, gomock.AssignableToTypeOf(presentations.Overtime{})).
		Return(nil)

	_, err := b.Overtime(ctx, input)
	assert.NoError(t, err)
}

func TestBusiness_Overtime_Failed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsers := users.NewMockUsers(ctrl)
	mockAttendance := attendance.NewMockAttendance(ctrl)
	mockOvertime := overtime.NewMockOvertime(ctrl)
	mockReimbursement := reimbursment.NewMockReimbursment(ctrl)
	mockPayslipSummary := payslipsummary.NewMockPayslipSummary(ctrl)
	mockPayroll := payroll.NewMockPayroll(ctrl)

	mockrepo := &repositories.Repository{
		Users:          mockUsers,
		Attendance:     mockAttendance,
		Payroll:        mockPayroll,
		Overtime:       mockOvertime,
		Reimbursement:  mockReimbursement,
		PayslipSummary: mockPayslipSummary,
	}

	b := overtimebus.NewBusiness(mockrepo)

	ctx := common.SetUserCtx(context.Background(), &entity.Claim{
		UserID:   "user-1",
		Username: "user1",
	})

	start := time.Now()
	end := start.Add(4 * time.Hour) // > 3 jam

	input := entity.Overtime{
		StartTime: start,
		EndTime:   end,
	}

	_, err := b.Overtime(ctx, input)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "overtime cannot more than 3 hours")
}

func TestBusiness_Overtime_Update_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsers := users.NewMockUsers(ctrl)
	mockAttendance := attendance.NewMockAttendance(ctrl)
	mockOvertime := overtime.NewMockOvertime(ctrl)
	mockReimbursement := reimbursment.NewMockReimbursment(ctrl)
	mockPayslipSummary := payslipsummary.NewMockPayslipSummary(ctrl)
	mockPayroll := payroll.NewMockPayroll(ctrl)

	mockrepo := &repositories.Repository{
		Users:          mockUsers,
		Attendance:     mockAttendance,
		Payroll:        mockPayroll,
		Overtime:       mockOvertime,
		Reimbursement:  mockReimbursement,
		PayslipSummary: mockPayslipSummary,
	}

	b := overtimebus.NewBusiness(mockrepo)

	ctx := common.SetUserCtx(context.Background(), &entity.Claim{
		UserID:   "user-1",
		Username: "user1",
	})

	start := time.Now()
	end := start.Add(2 * time.Hour)
	id := "ot-123"

	input := entity.Overtime{
		StartTime: start,
		EndTime:   end,
	}

	mockOvertime.EXPECT().
		Detail(ctx, id).
		Return(&presentations.Overtime{
			ID:     id,
			UserID: "user-1",
		}, nil)

	mockOvertime.EXPECT().
		Update(ctx, gomock.AssignableToTypeOf(presentations.Overtime{})).
		Return(nil)

	result, err := b.Update(ctx, input, id)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, id, result.ID)
	assert.Equal(t, "user-1", result.UserID)
}

func TestBusiness_Overtime_List_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsers := users.NewMockUsers(ctrl)
	mockAttendance := attendance.NewMockAttendance(ctrl)
	mockOvertime := overtime.NewMockOvertime(ctrl)
	mockReimbursement := reimbursment.NewMockReimbursment(ctrl)
	mockPayslipSummary := payslipsummary.NewMockPayslipSummary(ctrl)
	mockPayroll := payroll.NewMockPayroll(ctrl)

	mockrepo := &repositories.Repository{
		Users:          mockUsers,
		Attendance:     mockAttendance,
		Payroll:        mockPayroll,
		Overtime:       mockOvertime,
		Reimbursement:  mockReimbursement,
		PayslipSummary: mockPayslipSummary,
	}

	b := overtimebus.NewBusiness(mockrepo)

	ctx := common.SetUserCtx(context.Background(), &entity.Claim{
		UserID:   "user-1",
		Username: "user1",
	})

	params := &meta.Params{}

	expected := []presentations.Overtime{
		{
			ID:        "ot-1",
			UserID:    "user-1",
			StartTime: time.Now().Add(-3 * time.Hour),
			EndTime:   time.Now(),
			Duration:  3,
			CreatedAt: time.Now().Add(-3 * time.Hour),
			UpdatedAt: time.Now(),
			CreatedBy: "user1",
		},
	}

	mockOvertime.EXPECT().
		List(ctx, params, "user-1").
		Return(expected, nil)

	result, err := b.List(ctx, params)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}
