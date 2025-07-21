package common

import "time"

func CountWorkingDays(start, end time.Time) int {
	if start.After(end) {
		return 0
	}

	workingDays := 0
	current := start

	for !current.After(end) {
		weekday := current.Weekday()
		if weekday != time.Saturday && weekday != time.Sunday {
			workingDays++
		}
		current = current.AddDate(0, 0, 1)
	}

	return workingDays
}
