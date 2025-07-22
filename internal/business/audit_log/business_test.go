package auditlog_test

import (
	"context"
	"encoding/json"
	auditlog "payslips/internal/business/audit_log"
	"payslips/internal/presentations"
	"payslips/internal/repositories"
	"payslips/internal/repositories/attendance"
	auditlogrep "payslips/internal/repositories/audit_log"
	"payslips/internal/repositories/overtime"
	"payslips/internal/repositories/payroll"
	payslipsummary "payslips/internal/repositories/payslip_summary"
	"payslips/internal/repositories/reimbursment"
	"payslips/internal/repositories/users"
	"payslips/pkg/meta"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAuditLog_Create_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsers := users.NewMockUsers(ctrl)
	mockAttendance := attendance.NewMockAttendance(ctrl)
	mockOvertime := overtime.NewMockOvertime(ctrl)
	mockReimbursement := reimbursment.NewMockReimbursment(ctrl)
	mockPayslipSummary := payslipsummary.NewMockPayslipSummary(ctrl)
	mockPayroll := payroll.NewMockPayroll(ctrl)
	mockAuditLog := auditlogrep.NewMockAuditLog(ctrl)

	repo := &repositories.Repository{
		Users:          mockUsers,
		Attendance:     mockAttendance,
		Payroll:        mockPayroll,
		Overtime:       mockOvertime,
		Reimbursement:  mockReimbursement,
		PayslipSummary: mockPayslipSummary,
		AuditLog:       mockAuditLog,
	}

	b := auditlog.NewBusiness(repo)

	ctx := context.Background()

	input := presentations.AuditLog{
		ID:        uuid.NewString(),
		UserID:    "user-1",
		RequestID: uuid.NewString(),
		Path:      "/api/v1/payroll",
		Payload:   json.RawMessage(`{"test":"request"}`),
		Response:  json.RawMessage(`{"test":"response"}`),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		CreatedBy: "admin",
	}

	mockAuditLog.EXPECT().
		Create(ctx, input).
		Return(nil)

	err := b.Create(ctx, input)
	assert.NoError(t, err)
}

func TestAuditLog_List_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsers := users.NewMockUsers(ctrl)
	mockAttendance := attendance.NewMockAttendance(ctrl)
	mockOvertime := overtime.NewMockOvertime(ctrl)
	mockReimbursement := reimbursment.NewMockReimbursment(ctrl)
	mockPayslipSummary := payslipsummary.NewMockPayslipSummary(ctrl)
	mockPayroll := payroll.NewMockPayroll(ctrl)
	mockAuditLog := auditlogrep.NewMockAuditLog(ctrl)

	repo := &repositories.Repository{
		Users:          mockUsers,
		Attendance:     mockAttendance,
		Payroll:        mockPayroll,
		Overtime:       mockOvertime,
		Reimbursement:  mockReimbursement,
		PayslipSummary: mockPayslipSummary,
		AuditLog:       mockAuditLog,
	}

	b := auditlog.NewBusiness(repo)

	ctx := context.Background()

	expected := []presentations.AuditLog{
		{
			ID:        uuid.NewString(),
			UserID:    "user-1",
			RequestID: uuid.NewString(),
			Path:      "/api/v1/payroll",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			CreatedBy: "admin",
		},
	}

	mockAuditLog.EXPECT().
		List(ctx, gomock.Any()).
		Return(expected, nil)

	result, err := b.List(ctx, &meta.Params{})
	assert.NoError(t, err)
	assert.Equal(t, len(expected), len(result))
}
