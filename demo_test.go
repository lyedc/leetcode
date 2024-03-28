package leetcode


import (
	"container/heap"
	"fmt"
	"testing"
)

// 定义一个结构体来表示堆中的元素
type IntHeap2 []int

// Len 方法实现了 heap.Interface 接口
func (h IntHeap2) Len() int {
	return len(h)
}

// Less 方法实现了 heap.Interface 接口
func (h IntHeap2) Less(i, j int) bool {
	//return h[i] > h[j]  // 定义大腚对象
	return h[i] < h[j]  // 定义小顶堆
}

// Swap 方法实现了 heap.Interface 接口
func (h IntHeap2) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push 方法实现了 heap.Interface 接口
func (h *IntHeap2) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop 方法实现了 heap.Interface 接口
func (h *IntHeap2) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findKthLargest3(nums []int, k int) int {
	// 将数组转换为 IntHeap
	h := IntHeap2(nums)
	heap.Init(&h)

	// 构建大顶堆
	//for i := 1; i < k; i++ {
	//	heap.Push(&h, heap.Pop(&h).(int))
	//}
	// 因为是要找第k 大的元素，所以需要 k-1 次弹出堆顶元素,比如找第 2 大的元素，则需要弹出 1 次堆顶元素
	for i := 0; i < k-1; i++ {
		heap.Pop(&h)
	}

	// 返回堆顶元素，即第 k 个最大元素
	return h[0]
}

func Test_test(t *testing.T) {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 6
	fmt.Println(findKthLargest3(nums, k)) // 输出: 5

	nums = []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k = 9
	fmt.Println(findKthLargest3(nums, k)) // 输出: 4
}
