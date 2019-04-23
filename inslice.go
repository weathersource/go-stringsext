package stringsext

// InSlice checks for a string in a slice.
func InSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}