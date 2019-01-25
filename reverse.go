package stringsext

// Reverse reverses the order of a string.
// source: https://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go
// This does not work with combining characters like \u0301
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
