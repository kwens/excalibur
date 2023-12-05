/**
 * @Author: kwens
 * @Date: 2023-08-02 17:57:41
 * @Description:
 */
package sort


// 冒泡排序函数（优化版）
func BubbleSort[T sortType](arr []T) {
	n := len(arr)
	// 外层循环控制冒泡次数
	for i := 0; i < n-1; i++ {
		swapped := false // 标记是否有交换
		// 内层循环进行相邻元素比较并交换
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j] // 交换元素位置
				swapped = true
			}
		}
		// 如果一次冒泡没有任何交换，表示数组已经有序，结束排序
		if !swapped {
			break
		}
	}
}
