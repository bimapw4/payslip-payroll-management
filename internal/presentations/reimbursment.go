package presentations

import (
	"payslips/internal/common"
	"time"
)

const (
	ErrReimbursmentNotExist     = common.Error("err reimbursment not exist")
	ErrReimbursmentAlreadyExist = common.Error("err reimbursment already exist")
)

type Reimbursement struct {
	ID          string    `db:"id" json:"id"`
	UserID      string    `db:"user_id" json:"user_id"`
	Amount      int       `db:"amount" json:"amount"`
	Description string    `db:"description" json:"description"`
	Attachment  string    `db:"attachment" json:"attachment"`
	PayrollID   *string   `db:"payroll_id" json:"payroll_id,omitempty"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
	CreatedBy   string    `db:"created_by" json:"created_by"`
	UpdatedBy   string    `db:"updated_by" json:"updated_by"`
}
