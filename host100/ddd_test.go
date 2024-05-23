package host100

import (
	"fmt"
	"testing"
)

func quickSort2(arr []int) {
	if len(arr) <= 1 {
		return
	}
	QuickSort2(arr, 0, len(arr)-1)
}

func QuickSort2(arr []int, left, right int) {
	if left >= right {
		return
	}
	index := partition2(arr, left, right)
	QuickSort2(arr, index+1, right)
	QuickSort2(arr, left, index-1)
}

func partition2(arr []int, left, right int) int {
	tmp := arr[right]
	i := left
	for j := left; j < right; j++ {
		// 小于基准值的放在左边边，大于基准值的放在右边。当大于基准值的时候，这个值得位置不变，并且i也不变，
		// 所以最后循环完后，i就是新的基准值的位置。j会比i走的更快一点的.所以最后就是小于基准值的都跑到左边了.
		if arr[j] < tmp {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[right] = arr[right], arr[i]
	fmt.Println(arr)
	return i
}

func TestQuickSort2(t *testing.T) {
	arr := []int{0, 2, 0, 1, 5, 6}
	quickSort2(arr)
	for _, item := range arr {
		t.Log(item)
	}
}

func quickSort3(num []int) {
	if len(num) < 2 {
		return
	}

}

func QuickSort3(num []int, left, right int) {
	if left > right {
		return
	}
	index := partition3(num, left, right)
	QuickSort3(num, index+1, right)
	QuickSort3(num, left, index-1)
}

func partition3(num []int, left, right int) int {
	mid := num[right]
	i := left
	for j := left; j < right; j++ {
		if num[j] < mid {
			num[i], num[j] = num[j], num[i]
			i++
		}
	}
	num[i], num[right] = num[right], num[i]
	return i
}
