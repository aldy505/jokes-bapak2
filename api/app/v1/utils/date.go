package utils

import "time"

// IsToday checks if a date is in fact today or not.
func IsToday(date string) (bool, error) {
	parse, err := time.Parse(time.UnixDate, date)
	if err != nil {
		return false, err
	}

	now := time.Now()

	if parse.Equal(now) {
		return true, nil
	}
	return false, nil
}
