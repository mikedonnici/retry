package retry

import "time"

type Attempts struct {
	max   uint
	count uint
	wait  time.Duration
}

func Count(n uint) *Attempts {
	return &Attempts{max: n}
}

func CountWithWait(max uint, wait time.Duration) *Attempts {
	return &Attempts{max: max, wait: wait}
}

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
