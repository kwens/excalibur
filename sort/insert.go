/**
 * @Author: kwens
 * @Date: 2023-08-03 09:33:44
 * @Description:
 */
package sort

// 插入排序函数
func InsertionSort[T sortType](arr []T) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1

		// 将比 key 大的元素向后移动
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}
