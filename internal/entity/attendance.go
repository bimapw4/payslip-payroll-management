package entity

import (
	"payslips/internal/consts"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type AttendanceInput struct {
	Type     string    `json:"type"`
	Datetime time.Time `json:"date_time"`
}

func (v *AttendanceInput) Validation() error {
	return validation.ValidateStruct(
		v,
		validation.Field(&v.Type, validation.Required, validation.In(consts.AttendanceCheckin, consts.AttendanceCheckout).Error("must fill with check_in or check_out")),
		validation.Field(&v.Datetime, validation.Required),
	)
}
