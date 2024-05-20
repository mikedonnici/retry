package retry_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/regrowag/ses/pkg/retry"
)

func TestAttemptsCount(t *testing.T) {
	a := retry.Count(5)
	for i := 0; i < 5; i++ {
		assert.True(t, a.Next())
	}
	assert.False(t, a.Next())
}

func TestAttemptsCountWithWait(t *testing.T) {
	a := retry.CountWithWait(5, 100*time.Millisecond)
	t1 := time.Now()
	assert.True(t, a.Next())
	t2 := time.Now()
	assert.True(t, t2.Sub(t1) >= 100*time.Millisecond)
}
