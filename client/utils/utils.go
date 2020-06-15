package utils

import "strings"

// MakeLocalhostURL returns an URL with port on localhost
func MakeLocalhostURL(port string) string {
	return "http://localhost:" + port + "/"
}

// CleanseText clears \n and trims the input string
func CleanseText(text string) string {
	return strings.Trim(strings.Trim(strings.TrimSuffix(text, "\n"), "\""), " ")
}

// this above change was only made to account for "strings" within the arguments
// TODO: handle above note in a better manner
