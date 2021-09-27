package utils

import "time"

// IsToday checks if a date is in fact today or not.
func IsToday(date string) (bool, error) {
	if date == "" {
		return false, nil
	}

	parse, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return false, err
	}

	y1, m1, d1 := parse.Date()
	y2, m2, d2 := time.Now().Date()

	return y1 == y2 && m1 == m2 && d1 == d2, nil
}
