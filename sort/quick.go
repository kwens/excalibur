/**
 * @Author: kwens
 * @Date: 2023-08-03 09:13:51
 * @Description: 快速排序
 */
package sort

func QuickSort[T sortType](arr []T) []T {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[0]
	left := make([]T, 0, len(arr))
	right := make([]T, 0, len(arr))
	equal := make([]T, 0, len(arr))

	for _, num := range arr {
		if num < pivot {
			left = append(left, num)
		} else if num > pivot {
			right = append(right, num)
		} else {
			equal = append(equal, num)
		}
	}

	left = QuickSort(left)
	right = QuickSort(right)

	result := append(append(left, equal...), right...)
	return result
}
