package payroll

import (
	"context"
	"payslips/internal/common"
	"payslips/internal/entity"
	"payslips/internal/presentations"
	"payslips/internal/repositories"

	"github.com/google/uuid"
)

type Contract interface {
	CreatePayroll(ctx context.Context, payload entity.Payroll) error
	RunningPayroll(ctx context.Context, payrollID string) error
	GeneratePayslip(ctx context.Context, payrollID string) (*presentations.PayslipResponse, error)
}

type business struct {
	repo *repositories.Repository
}

func NewBusiness(repo *repositories.Repository) Contract {
	return &business{
		repo: repo,
	}
}

func (b *business) CreatePayroll(ctx context.Context, payload entity.Payroll) error {

	userctx := common.GetUserCtx(ctx)

	err := b.repo.Payroll.Create(ctx, presentations.Payroll{
		ID:          uuid.NewString(),
		PeriodStart: payload.PeriodStart,
		PeriodEnd:   payload.PeriodEnd,
		CreatedBy:   userctx.Username,
	})
	if err != nil {
		return err
	}

	return nil
}

func (b *business) RunningPayroll(ctx context.Context, payrollID string) error {
	userctx := common.GetUserCtx(ctx)

	ext, _ := b.repo.Payroll.Detail(ctx, payrollID)
	if ext.RunPayroll {
		return common.Error("payroll has already been processed")
	}

	err := b.repo.Attendance.UpdatePayrollID(ctx, payrollID, userctx.Username, ext.PeriodStart, ext.PeriodEnd)
	if err != nil {
		return err
	}

	err = b.repo.Reimbursement.UpdatePayrollID(ctx, payrollID, userctx.Username, ext.PeriodStart, ext.PeriodEnd)
	if err != nil {
		return err
	}

	err = b.repo.Overtime.UpdatePayrollID(ctx, payrollID, userctx.Username, ext.PeriodStart, ext.PeriodEnd)
	if err != nil {
		return err
	}

	err = b.repo.Payroll.Update(ctx, presentations.Payroll{
		ID:         payrollID,
		RunPayroll: true,
		UpdatedBy:  userctx.Username,
	})
	if err != nil {
		return err
	}

	return nil
}

func (b *business) GeneratePayslip(ctx context.Context, payrollID string) (*presentations.PayslipResponse, error) {
	userctx := common.GetUserCtx(ctx)

	user, err := b.repo.Users.Detail(ctx, userctx.UserID)
	if err != nil {
		return nil, err
	}

	payroll, err := b.repo.Payroll.Detail(ctx, payrollID)
	if err != nil {
		return nil, err
	}

	attendance, err := b.repo.Attendance.FindByPayrollID(ctx, userctx.UserID, payrollID)
	if err != nil {
		return nil, err
	}

	reimbursements, err := b.repo.Reimbursement.FindByPayrollID(ctx, userctx.UserID, payrollID)
	if err != nil {
		return nil, err
	}

	overtimes, err := b.repo.Overtime.FindByPayrollID(ctx, userctx.UserID, payrollID)
	if err != nil {
		return nil, err
	}

	// attendance
	workingDays := common.CountWorkingDays(payroll.PeriodStart, payroll.PeriodEnd)
	proratedSalary := (user.Salary / workingDays) * len(attendance)

	// reimbursement
	totalReimb := presentations.SumReimbursement(reimbursements)

	//overtime
	totalOvertimeHours := presentations.SumOvertime(overtimes)
	overtimeRate := float64(user.Salary) / float64(workingDays) / 8 * 2
	overtimePay := overtimeRate * totalOvertimeHours

	totalTakeHome := int(float64(proratedSalary) + overtimePay + totalReimb)

	response := presentations.PayslipResponse{
		PayrollID: payrollID,
		Period: presentations.Period{
			Start: payroll.PeriodStart,
			End:   payroll.PeriodEnd,
		},
		Attendance: presentations.AttendanceBreakdown{
			WorkingDays:    workingDays,
			PresentDays:    len(attendance),
			AbsentDays:     workingDays - len(attendance),
			ProratedSalary: proratedSalary,
		},
		Overtime: presentations.OvertimeBreakdown{
			TotalHours:  totalOvertimeHours,
			OvertimePay: int(overtimePay),
		},
		Reimbursements: reimbursements,
		TotalTakeHome:  totalTakeHome,
	}

	return &response, nil
}
