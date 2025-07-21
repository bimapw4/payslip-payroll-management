package payroll_test

import (
	"context"
	"testing"
	"time"

	"payslips/internal/business/payroll"
	"payslips/internal/common"
	"payslips/internal/entity"
	"payslips/internal/presentations"
	"payslips/internal/repositories"
	"payslips/internal/repositories/attendance"
	"payslips/internal/repositories/overtime"
	payrollrepo "payslips/internal/repositories/payroll"
	payslipsummary "payslips/internal/repositories/payslip_summary"
	"payslips/internal/repositories/reimbursment"
	"payslips/internal/repositories/users"
	"payslips/pkg/meta"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestRunningPayroll_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Mock per domain
	mockUsers := users.NewMockUsers(ctrl)
	mockAttendance := attendance.NewMockAttendance(ctrl)
	mockOvertime := overtime.NewMockOvertime(ctrl)
	mockReimbursement := reimbursment.NewMockReimbursment(ctrl)
	mockPayslipSummary := payslipsummary.NewMockPayslipSummary(ctrl)
	mockPayroll := payrollrepo.NewMockPayroll(ctrl)

	mockrepo := &repositories.Repository{
		Users:          mockUsers,
		Attendance:     mockAttendance,
		Payroll:        mockPayroll,
		Overtime:       mockOvertime,
		Reimbursement:  mockReimbursement,
		PayslipSummary: mockPayslipSummary,
	}

	b := payroll.NewBusiness(mockrepo)

	ctx := common.SetUserCtx(context.Background(), &entity.Claim{
		UserID:   uuid.NewString(),
		Username: "admin",
	})

	mockPayroll.EXPECT().Detail(gomock.Any(), "payroll-id").Return(&presentations.Payroll{
		ID:          "payroll-id",
		PeriodStart: time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC),
		PeriodEnd:   time.Date(2025, 7, 31, 0, 0, 0, 0, time.UTC),
		RunPayroll:  false,
	}, nil)

	mockAttendance.EXPECT().UpdatePayrollID(gomock.Any(), "payroll-id", gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
	mockReimbursement.EXPECT().UpdatePayrollID(gomock.Any(), "payroll-id", gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
	mockOvertime.EXPECT().UpdatePayrollID(gomock.Any(), "payroll-id", gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

	mockUsers.EXPECT().GetAllUsers(gomock.Any()).Return([]presentations.Users{
		{ID: "user-1", Salary: 10000000},
	}, nil)

	mockAttendance.EXPECT().FindByPayrollID(gomock.Any(), "user-1", "payroll-id").Return([]presentations.Attendance{
		{CheckIn: time.Now()},
	}, nil)

	mockReimbursement.EXPECT().FindByPayrollID(gomock.Any(), "user-1", "payroll-id").Return([]presentations.Reimbursement{
		{Amount: 500000},
	}, nil)

	mockOvertime.EXPECT().FindByPayrollID(gomock.Any(), "user-1", "payroll-id").Return([]presentations.Overtime{
		{Duration: 2},
	}, nil)

	mockPayslipSummary.EXPECT().Create(gomock.Any(), gomock.AssignableToTypeOf(presentations.PayslipSummary{})).Return(nil)

	mockPayroll.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)

	err := b.RunningPayroll(ctx, "payroll-id")
	assert.NoError(t, err)
}

func TestCreatePayroll_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Mock per domain
	mockUsers := users.NewMockUsers(ctrl)
	mockAttendance := attendance.NewMockAttendance(ctrl)
	mockOvertime := overtime.NewMockOvertime(ctrl)
	mockReimbursement := reimbursment.NewMockReimbursment(ctrl)
	mockPayslipSummary := payslipsummary.NewMockPayslipSummary(ctrl)
	mockPayroll := payrollrepo.NewMockPayroll(ctrl)

	mockrepo := &repositories.Repository{
		Users:          mockUsers,
		Attendance:     mockAttendance,
		Payroll:        mockPayroll,
		Overtime:       mockOvertime,
		Reimbursement:  mockReimbursement,
		PayslipSummary: mockPayslipSummary,
	}

	b := payroll.NewBusiness(mockrepo)

	ctx := common.SetUserCtx(context.Background(), &entity.Claim{
		UserID:   uuid.NewString(),
		Username: "admin",
	})

	mockPayroll.EXPECT().
		Create(gomock.Any(), gomock.AssignableToTypeOf(presentations.Payroll{})).
		Return(nil)

	err := b.CreatePayroll(ctx, entity.Payroll{
		PeriodStart: time.Now(),
		PeriodEnd:   time.Now().Add(time.Duration(3 * time.Minute)),
	})
	assert.NoError(t, err)
}

func TestGeneratePayslip_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsers := users.NewMockUsers(ctrl)
	mockAttendance := attendance.NewMockAttendance(ctrl)
	mockOvertime := overtime.NewMockOvertime(ctrl)
	mockReimbursement := reimbursment.NewMockReimbursment(ctrl)
	mockPayslipSummary := payslipsummary.NewMockPayslipSummary(ctrl)
	mockPayroll := payrollrepo.NewMockPayroll(ctrl)

	mockrepo := &repositories.Repository{
		Users:          mockUsers,
		Attendance:     mockAttendance,
		Payroll:        mockPayroll,
		Overtime:       mockOvertime,
		Reimbursement:  mockReimbursement,
		PayslipSummary: mockPayslipSummary,
	}

	b := payroll.NewBusiness(mockrepo)

	userID := uuid.NewString()
	username := "employee1"
	salary := 8000000

	ctx := common.SetUserCtx(context.Background(), &entity.Claim{
		UserID:   userID,
		Username: username,
	})

	payrollID := "payroll-id"
	periodStart := time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC)
	periodEnd := time.Date(2025, 7, 31, 0, 0, 0, 0, time.UTC)

	mockUsers.EXPECT().
		Detail(ctx, userID).
		Return(&presentations.Users{
			ID:     userID,
			Salary: salary,
		}, nil)

	mockPayroll.EXPECT().
		Detail(ctx, payrollID).
		Return(&presentations.Payroll{
			ID:          payrollID,
			PeriodStart: periodStart,
			PeriodEnd:   periodEnd,
		}, nil)

	// Mock data for attendance
	mockAttendance.EXPECT().
		FindByPayrollID(ctx, userID, payrollID).
		Return([]presentations.Attendance{
			{CheckIn: time.Now()},
			{CheckIn: time.Now()},
		}, nil)

	// Mock data for reimbursement
	mockReimbursement.EXPECT().
		FindByPayrollID(ctx, userID, payrollID).
		Return([]presentations.Reimbursement{
			{Amount: 200000},
			{Amount: 300000},
		}, nil)

	// Mock data for overtime
	mockOvertime.EXPECT().
		FindByPayrollID(ctx, userID, payrollID).
		Return([]presentations.Overtime{
			{Duration: 2.5},
		}, nil)

	// Run GeneratePayslip
	result, err := b.GeneratePayslip(ctx, payrollID)

	assert.NoError(t, err)
	assert.Equal(t, payrollID, result.PayrollID)
	assert.Equal(t, periodStart, result.Period.Start)
	assert.Equal(t, periodEnd, result.Period.End)

	assert.Equal(t, 2, result.Attendance.PresentDays)
	assert.True(t, result.Attendance.WorkingDays > 0)
	assert.True(t, result.TotalTakeHome > 0)
}

func TestListSummary_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsers := users.NewMockUsers(ctrl)
	mockAttendance := attendance.NewMockAttendance(ctrl)
	mockOvertime := overtime.NewMockOvertime(ctrl)
	mockReimbursement := reimbursment.NewMockReimbursment(ctrl)
	mockPayslipSummary := payslipsummary.NewMockPayslipSummary(ctrl)
	mockPayroll := payrollrepo.NewMockPayroll(ctrl)

	mockrepo := &repositories.Repository{
		Users:          mockUsers,
		Attendance:     mockAttendance,
		Payroll:        mockPayroll,
		Overtime:       mockOvertime,
		Reimbursement:  mockReimbursement,
		PayslipSummary: mockPayslipSummary,
	}

	b := payroll.NewBusiness(mockrepo)

	ctx := context.Background()
	payrollID := "payroll-id"

	expectedList := []presentations.PayslipSummary{
		{
			ID:          uuid.NewString(),
			UserID:      uuid.NewString(),
			BaseSalary:  8000000,
			TakeHomePay: 8500000,
		},
		{
			ID:          uuid.NewString(),
			UserID:      uuid.NewString(),
			BaseSalary:  7000000,
			TakeHomePay: 7200000,
		},
	}

	mockPayslipSummary.EXPECT().
		List(ctx, gomock.Any(), payrollID).
		Return(expectedList, nil)

	// Run
	result, err := b.ListSummary(ctx, &meta.Params{}, payrollID)

	assert.NoError(t, err)
	assert.Equal(t, len(expectedList), len(result))
	assert.Equal(t, expectedList[0].TakeHomePay, result[0].TakeHomePay)
}
