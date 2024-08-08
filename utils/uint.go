package utils

import "strconv"

func ParseUint(str string) (uint, error) {
    i, err := strconv.ParseUint(str, 10, 32)
    return uint(i), err
}
