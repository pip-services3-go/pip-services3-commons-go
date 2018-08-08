package run

import (
	"time"
)

type FixedRateTimer struct {
	task     INotifiable
	callback func()
	delay    int
	interval int
	ticker   *time.Ticker
}

func NewFixedRateTimer() *FixedRateTimer {
	return &FixedRateTimer{}
}

func NewFixedRateTimerFromCallback(callback func(), interval int, delay int) *FixedRateTimer {
	return &FixedRateTimer{
		callback: callback,
		interval: interval,
		delay:    delay,
	}
}

func NewFixedRateTimerFromTask(task INotifiable, interval int, delay int) *FixedRateTimer {
	c := &FixedRateTimer{
		interval: interval,
		delay:    delay,
	}
	c.SetTask(task)
	return c
}

func (c *FixedRateTimer) Task() INotifiable {
	return c.task
}

func (c *FixedRateTimer) SetTask(value INotifiable) {
	c.task = value
	c.callback = func() {
		c.task.Notify("timer", NewEmptyParameters())
	}
}

func (c *FixedRateTimer) Callback() func() {
	return c.callback
}

func (c *FixedRateTimer) SetCallback(value func()) {
	c.callback = value
	c.task = nil
}

func (c *FixedRateTimer) Delay() int {
	return c.delay
}

func (c *FixedRateTimer) SetDelay(value int) {
	c.delay = value
}

func (c *FixedRateTimer) Interval() int {
	return c.interval
}

func (c *FixedRateTimer) SetInterval(value int) {
	c.interval = value
}

func (c *FixedRateTimer) IsStarted() bool {
	return c.ticker != nil
}

func (c *FixedRateTimer) Start() {
	// Todo: lock to avoid concurrency issues
	// Stop previously set timer
	c.Stop()

	// Exit if interval is not defined
	if c.interval <= 0 {
		return
	}

	// Introducing delay
	delay := c.delay - c.interval
	ticker := time.NewTicker(time.Millisecond * time.Duration(c.interval))
	c.ticker = ticker

	go func() {
		if delay > 0 {
			time.Sleep(time.Millisecond * time.Duration(delay))
		}

		for range ticker.C {
			callback := c.callback
			if callback != nil {
				callback()
			}
		}
	}()
}

func (c *FixedRateTimer) Stop() {
	// Todo: lock to avoid concurrency
	ticker := c.ticker
	if ticker != nil {
		ticker.Stop()
		c.ticker = nil
	}
}

func (c *FixedRateTimer) Close(correlationId string) error {
	c.Stop()
	return nil
}
