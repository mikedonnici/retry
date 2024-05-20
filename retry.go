package retry

import "time"

// Attempts represents a number of attempts to perform an operation.
type Attempts struct {
	max   uint
	count uint
	wait  time.Duration
}

// Count returns a pointer to an Attempts value with the specified number of attempts.
func Count(n uint) *Attempts {
	return &Attempts{max: n}
}

// CountWithWait returns a pointer to an Attempts value with the specified number of attempts and wait time.
func CountWithWait(max uint, wait time.Duration) *Attempts {
	return &Attempts{max: max, wait: wait}
}

// Next returns true if the number of attempts has not been exceeded.
func (a *Attempts) Next() bool {
	if a.count >= a.max {
		return false
	}
	a.count++
	if a.wait > 0 {
		time.Sleep(a.wait)
	}
	return true
}
