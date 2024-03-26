package leetcode

import (
	"testing"
)

/*
Given an integer array nums, move all 0’s to the end of it while maintaining the relative order of the non-zero elements.

Note that you must do this in-place without making a copy of the array.

Example 1:

Input: nums = [0,1,0,3,12] Output: [1,3,12,0,0]

Example 2:

Input: nums = [0] Output: [0]

 283. 移动零 原题目：283. Move Zeroes
*/

func moveZeroes(nums []int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			for j := i + 1; j < len(nums); j++ {
				if nums[j] != 0 {
					nums[i], nums[j] = nums[j], nums[i]
					break
				}
			}
		}
	}
}

func moveZeroes2(nums []int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			for j := i; j < len(nums); j++ {
				if nums[j] != 0 {
					nums[i], nums[j] = nums[j], nums[i]
					// 以为要保持相对顺序，所以要跳出循环
					break
				}
			}
		}
	}
}

func moveZeroes3(nums []int) {
	indexFirst := 0
	for _, item := range nums {
		if item != 0 {
			nums[indexFirst] = item
			// 先把所有不等于0的数据都排到最签名
			indexFirst++
		}
	}
	// 最后把indexFirst起止到nums最后的数据全部都填充为0
	for j := indexFirst; j < len(nums); j++ {
		nums[j] = 0
	}
}

// test case
func TestMoveZeroes(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	t.Log(nums)

	nums2 := []int{0, 1, 0, 3, 12}
	moveZeroes2(nums2)
	t.Log(nums2)

	nums3 := []int{0, 1, 0, 3, 12}
	moveZeroes3(nums3)
	t.Log(nums3)
}
