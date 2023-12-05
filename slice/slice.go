/**
 * @Author: kwens
 * @Date: 2023-05-04 16:30:13
 * @Description:
 */
package slice

import "strconv"

func String2Int(s []string) []uint {
	var tmp []uint
	for _, v := range s {
		iv, err := strconv.Atoi(v)
		if err != nil {
			return nil
		}
		tmp = append(tmp, uint(iv))
	}
	return tmp
}

