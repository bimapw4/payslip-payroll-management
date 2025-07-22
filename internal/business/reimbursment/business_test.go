package reimbursment_test

import (
	"context"
	"payslips/internal/business/reimbursment"
	"payslips/internal/common"
	"payslips/internal/entity"
	"payslips/internal/presentations"
	"payslips/internal/repositories"
	"payslips/internal/repositories/attendance"
	"payslips/internal/repositories/overtime"
	"payslips/internal/repositories/payroll"
	payslipsummary "payslips/internal/repositories/payslip_summary"
	reimbursmentrepo "payslips/internal/repositories/reimbursment"
	"payslips/internal/repositories/users"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestReimbursement_Create_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsers := users.NewMockUsers(ctrl)
	mockAttendance := attendance.NewMockAttendance(ctrl)
	mockOvertime := overtime.NewMockOvertime(ctrl)
	mockReimbursement := reimbursmentrepo.NewMockReimbursment(ctrl)
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

	b := reimbursment.NewBusiness(mockrepo)

	ctx := common.SetUserCtx(context.Background(), &entity.Claim{
		UserID:   "user-1",
		Username: "user1",
	})

	input := entity.ReimbursementCreate{
		Amount:      500000,
		Description: "Transport",
	}

	mockReimbursement.EXPECT().
		Create(gomock.Any(), gomock.AssignableToTypeOf(presentations.Reimbursement{})).
		Return(nil)

	result, err := b.Create(ctx, input)
	assert.NoError(t, err)
	assert.Equal(t, input.Amount, result.Amount)
	assert.Equal(t, input.Description, result.Description)
	assert.Equal(t, "user-1", result.UserID)
}

func TestReimbursement_Update_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsers := users.NewMockUsers(ctrl)
	mockAttendance := attendance.NewMockAttendance(ctrl)
	mockOvertime := overtime.NewMockOvertime(ctrl)
	mockReimbursement := reimbursmentrepo.NewMockReimbursment(ctrl)
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

	b := reimbursment.NewBusiness(mockrepo)

	ctx := common.SetUserCtx(context.Background(), &entity.Claim{
		UserID:   "user-1",
		Username: "user1",
	})

	mockReimbursement.EXPECT().
		Detail(ctx, "reimb-1").
		Return(&presentations.Reimbursement{
			ID:     "reimb-1",
			UserID: "user-1",
		}, nil)

	mockReimbursement.EXPECT().
		Update(ctx, gomock.AssignableToTypeOf(presentations.Reimbursement{})).
		Return(nil)

	input := entity.ReimbursementUpdate{
		Id:          "reimb-1",
		Amount:      300000,
		Description: "Makan siang",
	}

	result, err := b.Update(ctx, input)
	assert.NoError(t, err)
	assert.Equal(t, input.Amount, result.Amount)
}

func TestReimbursement_Detail_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsers := users.NewMockUsers(ctrl)
	mockAttendance := attendance.NewMockAttendance(ctrl)
	mockOvertime := overtime.NewMockOvertime(ctrl)
	mockReimbursement := reimbursmentrepo.NewMockReimbursment(ctrl)
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

	b := reimbursment.NewBusiness(mockrepo)

	ctx := common.SetUserCtx(context.Background(), &entity.Claim{
		UserID:   "user-1",
		Username: "user1",
	})

	mockReimbursement.EXPECT().
		Detail(ctx, "reimb-1").
		Return(&presentations.Reimbursement{
			ID:     "reimb-1",
			UserID: "user-1",
			Amount: 500000,
		}, nil)

	result, err := b.Detail(ctx, "reimb-1")
	assert.NoError(t, err)
	assert.Equal(t, "reimb-1", result.ID)
}
