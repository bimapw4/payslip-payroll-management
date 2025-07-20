package entity

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Payroll struct {
	PeriodStart time.Time `json:"period_start"`
	PeriodEnd   time.Time `json:"period_end"`
}

func (v *Payroll) Validation() error {
	return validation.ValidateStruct(
		v,
		validation.Field(&v.PeriodStart, validation.Required),
		validation.Field(&v.PeriodEnd, validation.Required),
	)
}
