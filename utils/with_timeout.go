package utils

import (
	"errors"
	"time"
)

var ErrTimeout = errors.New("timeout: process took longer than expected")

func WithTimeout(d time.Duration, f func() error) error {
	done := make(chan bool, 1)
	go func() {
		f()
		done <- true
	}()

	select {
	case <-time.After(d):
		return ErrTimeout
	case <-done:
		return nil
	}
}
