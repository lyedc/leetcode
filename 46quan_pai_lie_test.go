package leetcode

import (
	"fmt"
	"testing"
)

/*
原题：LeetCode 46. 全排列

给定一个不含重复元素的数组 nums ，返回其所有可能的全排列。你可以按任意顺序返回答案。

示例 1：

输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,2,1],[3,1,2]]
示例 2：

输入：nums = [0,1]
输出：[[0,1],[1,0]]
使用 Go 语言解题的一种方法是使用回溯算法。回溯算法通过递归和剪枝的方式，
在搜索空间中探索所有可能的解

*/

func permute(nums []int)[][]int  {
	var result [][]int
	var current []int
	visited := make([]bool, len(nums))
	var backtrack func(first int)
	backtrack = func(first int) {
		// 递归结束条件, 所有元素都遍历完
		if first == len(nums) {
			// 将当前解添加到结果中
			tmp := make([]int, len(nums))
			copy(tmp, current)
			result = append(result, tmp)
		}
		for i := 0; i < len(nums); i++ {
			// 如果当前元素已经访问过，则跳过
			if visited[i]{
				continue
			}
			// 加入到结果集中,表示做选择
			current = append(current, nums[i])
			fmt.Println(current)
			visited[i] =true
			// 递归 进入下一层进行数据的筛选
			backtrack(first + 1)
			// 回溯，撤销选择
			// 什么是回溯？
			// 回溯就是从当前的解中撤销选择，重新选择下一个元素
			// 我们递归调用 backtrack 函数来生成下一个位置的排列。
			// 在递归返回后，我们需要撤销选择，即将当前数字从 current 切片中移除，
			// 并标记为未使用，以便进行下一次的尝试。
			current = current[:len(current) - 1]
			visited[i] = false
		}
	}
	backtrack(0)
	return result
}


func Test_permute(t *testing.T) {
	nums := []int{1, 2, 3}
	permutations := permute(nums)
	for _, p := range permutations {
		fmt.Println(p)
	}
}