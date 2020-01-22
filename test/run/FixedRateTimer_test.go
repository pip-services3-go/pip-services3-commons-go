package test_run

import (
	"testing"
	"time"

	"github.com/pip-services3-go/pip-services3-commons-go/run"
	"github.com/stretchr/testify/assert"
)

func TestTimerWithCallback(t *testing.T) {
	counter := 0

	timer := run.NewFixedRateTimerFromCallback(
		func() { counter++ },
		100, 0,
	)

	timer.Start()
	time.Sleep(time.Millisecond * 500)
	timer.Stop()

	assert.True(t, counter > 3)
}
