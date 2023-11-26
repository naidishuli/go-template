package utils

import (
	"regexp"
	"strings"
)

// Strip remove all duplicates and surrounding white spaces from a given string
func Strip(s string) string {
	ms := regexp.MustCompile(`\s+`)
	s = strings.TrimSpace(s)
	s = ms.ReplaceAllString(s, " ")
	return s
}

// ToSnakeCase return the snake_case version of a string from CamelCase
func ToSnakeCase(s string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(s, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func GetFromStringArray(arr []string, i int) string {
	if arr == nil {
		return ""
	}

	if i < 0 || i >= len(arr) {
		return ""
	}

	return arr[i]
}

func FetchString(value, def string) string {
	if value != "" {
		return value
	}

	return def
}
