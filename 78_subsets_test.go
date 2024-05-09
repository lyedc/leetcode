package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

/*
原题：LeetCode 78. 子集

给定一个可能包含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

说明：解集不能包含重复的子集。

示例:

输入: [1,2,2]

[]
[1]
[1, 2]
[1, 2, 2]
[2]
[2, 1] // 这个表示的是重复的
[2, 1, 2]  // 这个也是重复的.
[2, 2]
输出:
[
[],
[1],
[1,2],
[1,2,2],
[2],
[2,2]
]

*/

func Subsets(nums []int) [][]int {
	var result [][]int
	var current []int
	sort.Ints(nums) // 现对数组进行排序,保证重复元素相邻
	backtrackSubetes(nums, 0, current, &result)
	return result
}

func backtrackSubetes(num []int, index int, current []int, result *[][]int) {
	// 把当前子集加入到结果集中
	tmp := make([]int, len(current))
	copy(tmp, current)
	*result = append(*result, tmp)
	for i := index; i < len(num); i++ {
		// 跳过重复元素,避免生成重复的子集
		// 下一个元素和上一个元素相同的话,就不用再进行循环,因为不能有重复的元素
		if i > index && num[i] == num[i-1] {
			continue
		}
		current = append(current, num[i])
		backtrackSubetes(num, i+1, current, result)
		// 回溯(满足了就往上回溯..只要满足就回溯)
		current = current[:len(current)-1]
	}
}

func kuo_zhan_fa(nums []int) [][]int {
	var current []int
	var result [][]int
	result = append(result, current)
	for _, item := range nums {
		tmpResult := [][]int{}
		// 遍历每个结果集中的数据,把当前的数字加入到之前的结果中.
		for _, dd := range result {
			tmp := dd
			tmp = append(tmp, item)
			tmpResult = append(tmpResult, tmp)
		}
		// 把新配置成的结果,放入到最终的结果集中
		result = append(result, tmpResult...)
	}
	return result
}

// 深度优先算法
func dfs() {

}

func TestSubsets(t *testing.T) {
	nums := []int{1, 2, 2}
	subsetsResult := Subsets(nums)
	for _, subset := range subsetsResult {
		fmt.Println(subset)
	}
	fmt.Println("===================")
	subsetsResult2 := kuo_zhan_fa(nums)
	for _, item := range subsetsResult2 {
		fmt.Println(item)
	}
}
