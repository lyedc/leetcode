package leetcode

import (
	"container/heap"
	"fmt"
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

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] } // 小顶堆
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}
func (h *IntHeap) Pop() interface{} {
	a := *h
	v := a[a.Len()-1]
	*h = a[0 : a.Len()-1]
	return v
}

func findKthLargest5(nums []int, k int) int {
	tmp := IntHeap(nums)
	heap.Init(&tmp)
	for i := 0; i < k-1; i++ {
		heap.Pop(&tmp)
	}
	return tmp[0]
}

// test
func TestFindKthLargest5(t *testing.T) {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	res := findKthLargest5(nums, k)
	println(res)
	nums = []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k = 4
	fmt.Println(findKthLargest5(nums, k)) // 输出: 4
}

// 692 前 K 个高频单词

/*
给定一个单词列表 words 和一个整数 k ，返回前 k 个出现次数最多的单词。

返回的答案应该按单词出现频率由高到低排序。如果不同的单词有相同出现频率， 按字典顺序 排序
*/

/*

 */

func topKFrequent(words []string, k int) []string {
	tmpMap := make(map[string]int)
	for _, item := range words {
		tmpMap[item] += 1
	}
	// map 中保存的是每个单词出现的次数
	fmt.Print(tmpMap)
	return []string{}
}

// test

func Test_topKFrequent(t *testing.T) {
	words := []string{"i", "love", "leetcode", "i", "love", "coding"}
	k := 2
	fmt.Println(topKFrequent(words, k))
}
