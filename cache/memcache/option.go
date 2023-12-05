/**
 * @Author: kwens
 * @Date: 2022-08-29 16:44:13
 * @Description:
 */
package memcache

import "time"

type MemOption interface {
	apply(*memOption)
}

type memOption struct {
	Timeout time.Duration
}

var defaultMemOption = memOption{}

type memTimeout time.Duration

func (mt memTimeout) apply(opt *memOption) {
	opt.Timeout = time.Duration(mt)
}

func WithMemTimeout(t time.Duration) MemOption {
	return memTimeout(t)
}
