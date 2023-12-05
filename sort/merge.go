/**
 * @Author: kwens
 * @Date: 2023-08-03 09:18:26
 * @Description:
 */
package sort

// 归并排序函数
func MergeSort[T sortType](arr []T) []T {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return merge(left, right)
}

// 归并两个有序数组
func merge[T sortType](left, right []T) []T {
	result := make([]T, 0, len(left)+len(right))
	l, r := 0, 0

	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}

	result = append(result, left[l:]...)
	result = append(result, right[r:]...)

	return result
}
