/**
 * @Author: kwens
 * @Date: 2023-08-23 14:40:31
 * @Description: 并发执行控制器
 */
package controller

import (
	"sync"
	"sync/atomic"
	"time"
)

type ConcurrencyOption func(*concurrencyOptions)

type concurrencyOptions struct {
	stepDuration time.Duration
}

func WithConcurrencyStepDuration(sd time.Duration) ConcurrencyOption {
	return func(co *concurrencyOptions) {
		co.stepDuration = sd
	}
}

type ConcurrencyController struct {
	ch           chan struct{}
	counter      atomic.Int32
	wg           sync.WaitGroup
	isWait       atomic.Bool
	limit        int32
	stepDuration time.Duration
}

func NewConcurrencyController(concurrencyNum int32, opts ...ConcurrencyOption) *ConcurrencyController {
	var opt = &concurrencyOptions{}
	for _, o := range opts {
		o(opt)
	}
	return &ConcurrencyController{
		ch:           make(chan struct{}, concurrencyNum),
		limit:        concurrencyNum,
		stepDuration: opt.stepDuration,
	}
}

func (c *ConcurrencyController) AddAndWait() {
	c.ch <- struct{}{}
	if c.isWait.Load() {
		c.wg.Wait()
	}
	n := c.counter.Add(1)
	if n >= c.limit && c.isWait.CompareAndSwap(false, true) {
		c.wg.Add(1)
		if c.stepDuration > 0 {
			time.Sleep(c.stepDuration)
		}
	}
}

func (c *ConcurrencyController) Release() {
	defer func() {
		<-c.ch
	}()
	n := c.counter.Add(-1)
	if n == 0 && c.isWait.CompareAndSwap(true, false) {
		c.wg.Done()
	}
}
