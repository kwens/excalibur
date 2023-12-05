/**
 * @Author: kwens
 * @Date: 2023-08-02 18:02:24
 * @Description:
 */
package sort

import (
	"fmt"
	"testing"
)

func TestIntBubbleSort(t *testing.T) {

	arr := []int{11, 12, 22, 64, 34, 25, 90}
	fmt.Println("未排序前的数组:", arr)

	// 调用冒泡排序函数进行排序
	BubbleSort(arr)

	fmt.Println("排序后的数组:", arr)
}

func TestInt32BubbleSort(t *testing.T) {
	arr := []int32{11, 12, 22, 64, 34, 25, 90}
	fmt.Println("未排序前的数组:", arr)

	// 调用冒泡排序函数进行排序
	BubbleSort(arr)

	fmt.Println("排序后的数组:", arr)
}

func TestInt64BubbleSort(t *testing.T) {
	arr := []int64{11, 12, 22, 64, 34, 25, 90}
	fmt.Println("未排序前的数组:", arr)

	// 调用冒泡排序函数进行排序
	BubbleSort(arr)

	fmt.Println("排序后的数组:", arr)
}
func TestFloat64BubbleSort(t *testing.T) {
	arr := []float64{11.1, 12.1, 22.1, 64.1, 34.1, 25.1, 90.1}
	fmt.Println("未排序前的数组:", arr)

	// 调用冒泡排序函数进行排序
	BubbleSort(arr)

	fmt.Println("排序后的数组:", arr)
}
