package utils

import (
	"log"
	"time"
)

// RetryFunction retries a given function multiple times
func RetryFunction(attempts int, sleep time.Duration, fn func() error) error {
	var err error
	for i := 0; i < attempts; i++ {
		err = fn()
		if err == nil {
			return nil
		}
		log.Printf("Retry %d/%d failed: %v", i+1, attempts, err)
		time.Sleep(sleep)
	}
	return err
}

// FormatTime formats time in a readable string
func FormatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// Min returns the minimum of two integers
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Max returns the maximum of two integers
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
