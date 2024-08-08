package utils

import (
    "strings"
    "time"
)

func BeforeOrEqual(t1, t2 time.Time) bool {
    return t1.Equal(t2) || t1.Before(t2)
}

func AfterOrEqual(t1, t2 time.Time) bool {
    return t1.Equal(t2) || t1.After(t2)
}

type Date struct {
    time.Time
}

// implement an interface used by Fiber when parsign from query of the request to field in struct
func (d *Date) UnmarshalText(data []byte) error {
    s := strings.Trim(string(data), "\"") // Remove the double quotes around the JSON value
    if s == "null" || s == "" {
        d.Time = time.Time{} // Set to zero time if null or empty
        return nil
    }
    t, err := time.Parse("2006-01-02", s)
    if err != nil {
        return err
    }
    d.Time = t
    return nil
}

func (d *Date) UnmarshalJSON(data []byte) error {
    s := strings.Trim(string(data), "\"") // Remove the double quotes around the JSON value
    if s == "null" || s == "" {
        d.Time = time.Time{} // Set to zero time if null or empty
        return nil
    }

    t, err := time.Parse("2006-01-02", s)
    if err != nil {
        return err
    }

    d.Time = t.Truncate(24. * time.Hour)
    return nil
}
