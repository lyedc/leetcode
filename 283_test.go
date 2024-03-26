package leetcode

import "testing"

// 283. 移动零 原题目：283. Move Zeroes

func moveZeroes(nums []int) {

}

// test case
func TestMoveZeroes(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	t.Log(nums)
}
