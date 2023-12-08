package utils

// containsString checks if a string slice contains the given string.
// It iterates through the slice and returns true if the string is found.
func ContainsString(slice []string, str string) bool {
	for _, item := range slice {
		if item == str {
			return true
		}
	}
	return false
}
