/**
 * @Author: kwens
 * @Date: 2023-08-15 14:47:05
 * @Description:
 */
package string

import (
	"strconv"
	"strings"
)

func MustToSliceFloat64(s string) []float64 {
	var result []float64
	if len(s) == 0 {
		return result
	}
	sliceS := strings.Split(s, ",")
	for _, v := range sliceS {
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			continue
		}
		result = append(result, f)
	}
	return result
}

func FormatSliceFloat64(f []float64) string {
	var result []string
	if len(f) == 0 {
		return ""
	}
	for _, v := range f {
		result = append(result, strconv.FormatFloat(v, 'f', -1, 64))
	}
	return strings.Join(result, ",")
}
