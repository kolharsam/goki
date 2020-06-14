package utils

import "strings"

// MakeLocalhostURL returns an URL with port on localhost
func MakeLocalhostURL(port string) string {
	return "http://localhost:" + port + "/"
}

// CleanseText clears \n and trims the input string
func CleanseText(text string) string {
	return strings.Trim(strings.TrimSuffix(text, "\n"), " ")
}
