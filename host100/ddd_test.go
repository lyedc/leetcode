package host100

import "testing"

func quickSort2(arr []int){
	if len(arr) <=1{
		return
	}
	QuickSort2(arr, 0, len(arr) - 1)
}

func QuickSort2(arr []int, left, right int){
	if left >= right{
		return
	}
	index := partition2(arr, left, right)
	QuickSort2(arr, index +1, right)
	QuickSort2(arr, left, index -1)
}

func partition2(arr []int, left, right int)int{
	tmp := arr[right]
	i := left
	for j := left; j < right; j++{
		if arr[j] < tmp{
			arr[i], arr[j] = arr[j], arr[i]
			i ++
		}
	}
	arr[i], arr[right] = arr[right], arr[i]
	return i
}

func TestQuickSort2(t *testing.T) {
	arr := []int{0, 2, 0, 1, 5, 6}
	quickSort2(arr)
	for _,item := range arr{
		t.Log(item)
	}
}