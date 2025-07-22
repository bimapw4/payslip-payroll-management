package entity

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Overtime struct {
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

func (v *Overtime) Validation() error {
	return validation.ValidateStruct(
		v,
		validation.Field(&v.StartTime, validation.Required),
		validation.Field(&v.EndTime, validation.Required),
	)
}

func (o *Overtime) GetDuration() float64 {
	duration := o.EndTime.Sub(o.StartTime)
	return duration.Hours()
}
