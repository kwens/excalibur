/**
 * @Author: kwens
 * @Date: 2023-08-23 14:51:13
 * @Description:
 */
package controller

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

var counter atomic.Int32

func TestConcurrencyController(t *testing.T) {
	ctl := NewConcurrencyController(200)
	var totalNum = 10000
	var wg sync.WaitGroup
	var hold sync.WaitGroup
	wg.Add(1)
	for i := 0; i <= totalNum; i++ {
		if i == totalNum {
			wg.Done()
		}
		go func(n int) {
			wg.Wait()
			hold.Add(1)
			ctl.AddAndWait()
			defer func() {
				ctl.Release()
				hold.Done()
			}()
			do(n)
		}(i)
	}
	hold.Wait()
}

func do(i int) {
	counter.Add(1)
	time.Sleep(time.Millisecond * 100)
	fmt.Printf("do....n:%d\n", i)
	fmt.Printf("counter...n:%d\n", counter.Load())
}
