package leetcode

import (
	"container/heap"
	"testing"
)

/*
堆：一种二叉树的结构--》 完全二叉树
每个节点的值都大于等于左右子节点的值（最大堆）根节点是最大值（堆顶元素）
最小堆：每个节点的值都小于等于左右子节点的值（最小堆） 根节点是最小值。（堆顶元素）

在搜索一般是找堆顶元素，所以是O(1)
添加：logN
删除： logN 一般删除的都是堆顶元素，不是堆里面的任何一个元素。。。
添加删除都可能需要树上的元素。。。需要满足堆的特性。。。
*/

// 215 数组中的第K个最大元素
/*
给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。

请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。
*/

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}

func (h *IntHeap) Pop() interface{} {
	a := *h
	n := len(a)
	v := a[n-1]  // 数组中的最后一个
	*h = a[:n-1]  // 对堆进行重新的赋值操作.
	return v
}


func findKthLargest2(nums []int, k int) int {
	h := IntHeap(nums)
	heap.Init(&h)  // 初始化一个堆
	// 构建一个大顶堆
	for i := 0; i < k; i++ {
		heap.Push(&h, heap.Pop(&h).(int))
	}
	// 返回对顶元素
	return heap.Pop(&h).(int)
}

// test
func Test_findKthLargest2(t *testing.T) {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	t.Log(findKthLargest2(nums, k))
}

type Heap struct {
	data []int
	size int
}

func NewHeap(n int) *Heap {
	return &Heap{
		data: make([]int, n+1),
		size: 0,
	}
}


func (h *Heap) Add(v int) {
	h.data[h.size+1] = v
	h.size++
	// 添加元素后，需要调整堆
	for i := h.size; i > 1; i /= 2 {
		if h.data[i] > h.data[i/2] {
			h.data[i], h.data[i/2] = h.data[i/2], h.data[i]
		} else {
			break
		}
	}
}

func (h *Heap) Remove() {
	if h.size == 0 {
		return
	}
	h.data[1], h.data[h.size] = h.data[h.size], h.data[1]
	h.size--
	// 删除元素后，需要调整堆
	for i := 1; i*2 <= h.size; {
		if h.data[i*2] > h.data[i] {
			h.data[i], h.data[i*2] = h.data[i*2], h.data[i]
		}
	}
}

func (h *Heap) Peek() int {
	return h.data[1]
}

func findKthLargest(nums []int, k int) int {
	// 创建一个堆
	h := NewHeap(len(nums))
	for _, v := range nums {
		h.Add(v)
	}
	for i := 0; i < k-1; i++ {
		h.Remove()
	}
	return h.Peek()
}

// 692 寻找前K个高频元素

func Test_findKthLargest(t *testing.T) {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k := 2
	t.Log(findKthLargest(nums, k))

}
