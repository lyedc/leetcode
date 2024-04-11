package host100

import "testing"

// 快排

func QuickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	quickSort(arr, 0, len(arr)-1)
}
func quickSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	pivot := partition(arr, left, right)
	// 排基准值坐标的
	quickSort(arr, left, pivot-1)
	// 排基准值右边的
	quickSort(arr, pivot+1, right)
	return
}

func partition(arr []int, left int, right int) int {
	// 选择最后一个元素作为基准值
	pivot := arr[right]
	// 新的基准值的位置
	i := left
	for j := left; j < right; j++ {
		// 将小于基准值的元素交换到前面
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	// 将基准值交换到正确的位置
	arr[i], arr[right] = arr[right], arr[i]
	// 返回基准值的索引
	return i
}

// test
// 测试用例
func TestQuickSort(t *testing.T) {
	arr := []int{5, 2, 9, 1, 5, 6}
	QuickSort(arr)
	for _,item := range arr{
		t.Log(item)
	}
}




