package trydo

import (
	"errors"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTryTimes(t *testing.T) {
	do := newTestDoFn()

	err := TryTimes(do, 2)
	assert.NotNil(t, err)

	err = TryTimes(do, 3)
	assert.Nil(t, err)
}

func TestTryWithIntervals(t *testing.T) {
	do := newTestDoFn()

	err := TryWithIntervals(do, time.Second, time.Second, time.Second)
	assert.Nil(t, err)
	// PASS: TestTryWithIntervals (2.00s)
}

func newTestDoFn() DoFn {
	var i int32 = 0
	do := func() error {
		c := atomic.AddInt32(&i, 1)
		if c < 3 {
			return errors.New("err")
		}
		return nil
	}
	return do
}
