/**
 * @Author: kwens
 * @Date: 2023-07-27 11:43:21
 * @Description:
 */
package rand

import (
	"math/rand"
	"time"
)

// RandInt 随机返回n位数字
func RandInt(n int) int {
	// 使用当前时间初始化随机数生成器
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// 计算最小值和最大值，确保生成n位数字
	minValue := intPow(10, n-1)
	maxValue := intPow(10, n) - 1

	// 生成一个介于minValue和maxValue之间的随机数
	randomNumber := r.Intn(maxValue-minValue+1) + minValue
	return randomNumber
}

// 辅助函数：计算x的y次方
func intPow(x, y int) int {
	result := 1
	for i := 0; i < y; i++ {
		result *= x
	}
	return result
}
