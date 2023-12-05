/**
 * @Author: kwens
 * @Date: 2022/7/14 13:45
 * @Description:
 */

package token

import "time"

type TokenOption interface {
	apply(option *tokenOption)
}

type tokenOption struct {
	Exp time.Duration
}

var defaultOption = tokenOption{
	Exp: TOKEN_EXP,
}

var (
	TOKEN_EXP = time.Hour * 24
)

type expTime time.Duration

func (et expTime) apply(opt *tokenOption) {
	opt.Exp = time.Duration(et)
}

func WithExp(exp time.Duration) TokenOption {
	return expTime(exp)
}
