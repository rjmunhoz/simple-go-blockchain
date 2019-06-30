package domain

import (
	"fmt"
	"time"
)

// Repeat repeats fn * times
func Repeat(fn func(i int), times int) {
	counter := times

	for i := 0; i < counter; i++ {
		fn(i)
	}
}

// GetTimeTracker returns a callback that should be called when the timer ends
func GetTimeTracker(what string, startTime *time.Time) func() string {
	return func() string {
		return fmt.Sprintf(what+". Took: %v", time.Since(*startTime))
	}
}
