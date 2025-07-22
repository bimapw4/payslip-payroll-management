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
