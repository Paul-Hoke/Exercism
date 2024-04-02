package logs

import "unicode/utf8"
import "strings"

// Application identifies the application emitting the given log.
func Application(log string) string {
	
	myRune := '0'
	
	for _, char := range log {
		if char == '❗' || char == '🔍' || char == '☀' {
			myRune = char
			break
		}
	}
	
	switch myRune {
	case '❗':
		return "recommendation"
	case '🔍':
		return "search"
	case '☀':
		return "weather"
	default:
		return "default"
	}
	
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	return strings.ReplaceAll(log, string(oldRune), string(newRune))
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
