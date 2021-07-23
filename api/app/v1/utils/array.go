package utils

// IsIn checks if an array have a value
func IsIn(arr []string, value string) bool {
	for _, item := range arr {
		if item == value {
			return true
		}
	}
	return false
}