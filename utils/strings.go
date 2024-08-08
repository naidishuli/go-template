package utils

import (
    "reflect"
    "regexp"
    "strings"
    "unicode/utf8"

    "github.com/google/uuid"
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

func ToLowerPtr(s *string) *string {
    if s == nil {
        return nil
    }

    *s = strings.ToLower(*s)
    return s
}

func IsUUID(str string) bool {
    _, err := uuid.Parse(str)
    return err == nil
}

func Truncate[T string | *string](s T, maxBytes int) T {
    var str string

    // Check if the input is a pointer or a string
    v := reflect.ValueOf(s)
    if v.Kind() == reflect.Ptr {
        if v.IsNil() {
            return s // Return nil pointer as is
        }
        str = *v.Interface().(*string)
    } else {
        str = v.Interface().(string)
    }

    if len(str) <= maxBytes {
        return s
    }

    truncated := str[:maxBytes]

    // Ensure we do not cut a multibyte character in half
    for !utf8.ValidString(truncated) {
        truncated = truncated[:len(truncated)-1]
    }

    if v.Kind() == reflect.Ptr {
        truncatedPtr := new(string)
        *truncatedPtr = truncated
        return any(truncatedPtr).(T)
    }
    return any(truncated).(T)
}
