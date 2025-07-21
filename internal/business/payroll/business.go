package payroll

import (
	"context"
	"payslips/internal/common"
	"payslips/internal/entity"
	"payslips/internal/presentations"
	"payslips/internal/repositories"
	"payslips/pkg/meta"

	"github.com/google/uuid"
)

type Contract interface {
	CreatePayroll(ctx context.Context, payload entity.Payroll) error
	RunningPayroll(ctx context.Context, payrollID string) error
	GeneratePayslip(ctx context.Context, payrollID string) (*presentations.PayslipResponse, error)
	ListSummary(ctx context.Context, m *meta.Params, payrollId string) ([]presentations.PayslipSummary, error)
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

	users, err := b.repo.Users.GetAllUsers(ctx)
	if err != nil {
		return err
	}

	for _, v := range users {
		calc, err := b.calculatePayslip(ctx, v.ID, v.Salary, ext)
		if err != nil {
			return err
		}

		err = b.repo.PayslipSummary.Create(ctx, presentations.PayslipSummary{
			ID:                 uuid.NewString(),
			PayrollID:          ext.ID,
			UserID:             v.ID,
			BaseSalary:         v.Salary,
			ProratedSalary:     calc.ProratedSalary,
			OvertimePay:        int(calc.OvertimePay),
			ReimbursementTotal: int(calc.TotalReimb),
			TakeHomePay:        calc.TotalTakeHome,
			CreatedBy:          userctx.Username,
		})
		if err != nil {
			return err
		}
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

	calc, err := b.calculatePayslip(ctx, user.ID, user.Salary, payroll)
	if err != nil {
		return nil, err
	}

	response := presentations.PayslipResponse{
		PayrollID: payrollID,
		Period: presentations.Period{
			Start: payroll.PeriodStart,
			End:   payroll.PeriodEnd,
		},
		Attendance: presentations.AttendanceBreakdown{
			WorkingDays:    calc.WorkingDays,
			PresentDays:    calc.PresentDays,
			AbsentDays:     calc.WorkingDays - calc.PresentDays,
			ProratedSalary: calc.ProratedSalary,
		},
		Overtime: presentations.OvertimeBreakdown{
			TotalHours:  calc.TotalOvertimeHours,
			OvertimePay: int(calc.OvertimePay),
		},
		Reimbursements: calc.Reimbursements,
		TotalTakeHome:  calc.TotalTakeHome,
	}

	return &response, nil
}

func (b *business) calculatePayslip(ctx context.Context, userID string, salary int, payroll *presentations.Payroll) (*presentations.CalculatePayslip, error) {

	attendance, err := b.repo.Attendance.FindByPayrollID(ctx, userID, payroll.ID)
	if err != nil {
		return nil, err
	}

	reimbursements, err := b.repo.Reimbursement.FindByPayrollID(ctx, userID, payroll.ID)
	if err != nil {
		return nil, err
	}

	overtimes, err := b.repo.Overtime.FindByPayrollID(ctx, userID, payroll.ID)
	if err != nil {
		return nil, err
	}

	// attendance
	workingDays := common.CountWorkingDays(payroll.PeriodStart, payroll.PeriodEnd)
	proratedSalary := (salary / workingDays) * len(attendance)

	// reimbursement
	totalReimb := presentations.SumReimbursement(reimbursements)

	//overtime
	totalOvertimeHours := presentations.SumOvertime(overtimes)
	overtimeRate := float64(salary) / float64(workingDays) / 8 * 2
	overtimePay := overtimeRate * totalOvertimeHours

	totalTakeHome := int(float64(proratedSalary) + overtimePay + totalReimb)

	return &presentations.CalculatePayslip{
		WorkingDays:        workingDays,
		ProratedSalary:     proratedSalary,
		TotalReimb:         totalReimb,
		TotalOvertimeHours: totalOvertimeHours,
		OvertimePay:        overtimePay,
		TotalTakeHome:      totalTakeHome,
		PresentDays:        len(attendance),
		Reimbursements:     reimbursements,
	}, nil
}

func (b *business) ListSummary(ctx context.Context, m *meta.Params, payrollId string) ([]presentations.PayslipSummary, error) {
	list, err := b.repo.PayslipSummary.List(ctx, m, payrollId)
	if err != nil {
		return nil, err
	}
	return list, nil
}
