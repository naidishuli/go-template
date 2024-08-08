package validator

import (
    "strconv"
    "strings"
    "time"

    "backend/utils"
    "github.com/go-playground/validator/v10"
)

func dataOfBirth(fl validator.FieldLevel) bool {
    var birthTime time.Time
    var ok bool

    birthTime, ok = fl.Field().Interface().(time.Time)
    if !ok {
        date := fl.Field().Interface().(utils.Date)
        birthTime = date.Time
    }

    params := fl.Param() // Get the parameters from the tag

    var minAge, maxAge *int

    // Split the parameters to get the min and max age
    ageRange := strings.Split(params, "-")

    if len(ageRange) > 0 {
        val, err := strconv.Atoi(ageRange[0])
        if err != nil {
            return false
        }

        minAge = &val
    }

    if len(ageRange) > 1 {
        val, err := strconv.Atoi(ageRange[1])
        if err != nil {
            return false
        }

        maxAge = &val
    }

    now := time.Now()
    years := now.Year() - birthTime.Year()

    // Check if the current date is before the birth date in terms of month and day
    if now.Month() < birthTime.Month() || (now.Month() == birthTime.Month() && now.Day() < birthTime.Day()) {
        years--
    }

    minAgeCheck, maxAgeCheck := true, true
    if minAge != nil {
        minAgeCheck = years >= *minAge
    }

    if maxAge != nil {
        maxAgeCheck = years <= *maxAge
    }

    return minAgeCheck && maxAgeCheck
}
