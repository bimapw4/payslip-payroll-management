package presentations

import (
	"payslips/internal/common"
	"time"
)

const (
	ErrAuthNotExist       = common.Error("err auth not exist")
	ErrAuthAlreadyExist   = common.Error("err auth already exist")
	ErrAuthAlreadyRevoked = common.Error("err auth already revoked")
)

type AuthorizationResp struct {
	AccessToken string `json:"access_token" db:"access_token"`
	// RefreshToken string `json:"refresh_token" db:"refresh_token"`
}

type Authorization struct {
	ID     string `json:"id" db:"id"`
	UserID string `json:"user_id" db:"user_id"`
	AuthorizationResp
	IsActive  bool      `json:"is_active" db:"is_active"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
