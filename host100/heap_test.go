package host100

import (
	"container/heap"
	"fmt"
	"testing"
)

// leetcode 215 数组中的第K个最大元素
func findKthLargest(nums []int, k int) int {
	// 这里使用的是 lens(nums) -k 表示的是倒数第k个元素，如果这里传入的是k的话，下面的快速排序
	// 就需要大于基准值。这样大的数据就会排在前面，小的排在后面。
	quickSelect(nums, 0, len(nums)-1, len(nums)-k)
	return 0
}
// 使用快排的方式实现
func quickSelect(arr []int, start ,end, k int)int{
	pivot := arr[end]
	i := start
	for j := start; j<end; j++{
		if arr[j] <pivot{
			arr[i],arr[j] = arr[j],arr[i]
			i++
		}
	}
	arr[i], arr[end] = arr[end], arr[i]
	if i == k{
		return arr[i]
	}
	if i < k{
		return quickSelect(arr, i +1, end, k)
	}else{
		return quickSelect(arr, start, i-1, k)
	}
}


// 先实现一个堆

type IntHeap []int
func (h IntHeap) Len() int {
	return len(h)
}
// 这里是小顶堆，最上面是最小值。
func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0:n-1]
	return x
}

func findKthLargest2(nums []int, k int) int {
	return heapSelect(nums, k)
}

// 使用堆的形式实现。
func heapSelect(arr []int, k int)int{
	newHeap := IntHeap(arr)
	heap.Init(&newHeap)
	// 把堆顶的元素弹出k-1个数据这样堆顶的元素就是第K大的元素
	// 如果是大顶堆的话，就需要弹出
	for i :=0; i<k-1;i++{
		heap.Pop(&newHeap)
	}
	return newHeap[0]
}

type stringHeap []string

func (h stringHeap)Len()int{
	return len(h)
}
func (h stringHeap)Less(i, j int)bool{
	return h[i] < h[j]
}

func (h stringHeap)Swap(i, j int){
	h[i], h[j] = h[j], h[i]
}

func (h *stringHeap)Push(x interface{}){
	*h = append(*h, x.(string))
}

func (h *stringHeap)Pop()interface{}{
	//x := h[len(*h) -1]
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0:n-1]
	return x

}

	type elem struct {
		num int
		value int
	}

	type elemHeap []elem

	func (h elemHeap)Len()int{
		return len(h)
	}

	// 大顶堆
	func (h elemHeap)Less(i, j int)bool{
		return h[i].value> h[j].value
	}

	func (h elemHeap)Swap(i, j int){
		h[i], h[j] = h[j], h[i]
	}
	func (h *elemHeap)Push(x interface{}){
		*h = append(*h, x.(elem))
	}

	func (h *elemHeap)Pop()interface{}{
		old := *h
		n := len(old)
		x := old[n-1]
		*h = old[0:n-1]
		return x
	}


	// leetcode 347 前K个高频元素
	func topKFrequent(nums []int, k int) []int {
		tmpMap := make(map[int]int)
		for _, item := range nums{
			tmpMap[item]++
		}
		newHeap := elemHeap{}
		heap.Init(&newHeap)
		// 最小堆顶元素是最小的。如果需要k个频率高的，那么久需要pop出k-1个元素
		for key, value := range tmpMap{
			heap.Push(&newHeap, elem{num: key, value: value})
		}
		result := []int{}
		for i:=0;i<k;i++{
			tmp := heap.Pop(&newHeap).(elem)
			result = append(result, tmp.num)
		}
		return result
	}

// test
func TestHeap(t *testing.T){
	tmp1 := []int{1,1,1,2,2,3}
	k := 2
	fmt.Println(topKFrequent(tmp1, k))

	tmp1 = []int{1}
	k = 1
	fmt.Println(topKFrequent(tmp1, k))
}