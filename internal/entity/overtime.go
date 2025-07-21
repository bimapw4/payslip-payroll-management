package entity

import "time"

type Overtime struct {
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

func (o *Overtime) GetDuration() float64 {
	duration := o.EndTime.Sub(o.StartTime)
	return duration.Hours()
}
