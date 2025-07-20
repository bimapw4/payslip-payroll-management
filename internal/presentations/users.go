package presentations

import (
	"payslips/internal/common"
	// usercore "payslips/internal/provider/user_core"
	"time"
)

const (
	ErrUserNotExist     = common.Error("err users not exist")
	ErrUserAlreadyExist = common.Error("err users already exist")
)

type Users struct {
	ID        string    `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Username  string    `db:"username" json:"username"`
	Password  string    `db:"password" json:"-"`
	Salary    int       `db:"salary" json:"salary"`
	IsAdmin   bool      `db:"is_admin" json:"is_admin"`
	IsActive  bool      `db:"is_active" json:"is_active"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	CreatedBy string    `db:"created_by" json:"created_by"`
	UpdatedBy string    `db:"updated_by" json:"updated_by"`
}
