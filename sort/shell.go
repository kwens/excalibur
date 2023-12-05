/**
 * @Author: kwens
 * @Date: 2023-08-03 09:16:55
 * @Description: 希尔排序
 */
package sort

// 希尔排序函数
func ShellSort[T sortType](arr []T) {
	n := len(arr)
	// 定义增量序列
	for gap := n / 2; gap > 0; gap /= 2 {
		// 对每个子序列进行插入排序
		for i := gap; i < n; i++ {
			temp := arr[i]
			j := i
			for j >= gap && arr[j-gap] > temp {
				arr[j] = arr[j-gap]
				j -= gap
			}
			arr[j] = temp
		}
	}
}
