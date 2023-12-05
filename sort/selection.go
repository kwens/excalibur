/**
 * @Author: kwens
 * @Date: 2023-08-03 09:11:08
 * @Description:
 */
package sort

// 选择排序函数
func SelectionSort[T sortType](arr []T) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		// 假设当前元素为最小值
		minIndex := i
		// 在未排序序列中查找最小值
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		// 将最小值与当前元素交换位置
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}
