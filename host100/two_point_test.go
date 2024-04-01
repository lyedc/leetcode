package host100

import (
	"testing"
)

// leetcode 283 move zeroes
/*
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

请注意 ，必须在不复制数组的情况下原地对数组进行操作。

示例 1:

输入: nums = [0,1,0,3,12]
输出: [1,3,12,0,0]
示例 2:

输入: nums = [0]
输出: [0]
*/

// 采用覆盖的形式完成。
func moveZeroes(nums []int) {
	moveIndex := 0
	for _, value := range nums {
		if value != 0 {
			nums[moveIndex] = value
			moveIndex++
		}
	}
	//不是零的元素都移动到前面去了。 然后把数组剩余的位置补零
	// tmpNum := len(nums) - moveIndex
	for i := len(nums) - 1; i >= moveIndex; i-- {
		nums[i] = 0
	}
}

// 采用了左右双指针的形式移动了。
func moveZeroes2(nums []int) {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

// test
func TestMoveZeroes(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	t.Log(nums)

	nums1 := []int{1, 1, 2, 3, 12}
	moveZeroes(nums1)
	t.Log(nums1)

	nums2 := []int{0}
	moveZeroes(nums2)
	t.Log(nums2)
}

func TestMoveZeroes2(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes2(nums)
	t.Log(nums)

	nums1 := []int{1, 1, 2, 3, 12}
	moveZeroes2(nums1)
	t.Log(nums1)

	nums2 := []int{0}
	moveZeroes2(nums2)
	t.Log(nums2)
}

// leetcode 11 盛最多水的容器
/*
给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。

找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

返回容器可以储存的最大水量。

说明：你不能倾斜容器。
输入：[1,8,6,2,5,4,8,3,7]
输出：49
解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
示例 2：

输入：height = [1,1]
输出：1

*/
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 双指针，谁矮就移动谁。。。这样继续组成围城的空间。
func maxArea(height []int) int {
	left, right, ants := 0, len(height)-1, 0
	for left < right {
		tmpAnts := (right - left) * min(height[left], height[right])
		if tmpAnts > ants {
			ants = tmpAnts
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}

	}
	return ants
}

// test

func Test_maxArea123(t *testing.T) {
	// [1,8,6,2,5,4,8,3,7]
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	t.Log(maxArea(height))
}

// leetcode 15 三数之和

/*
给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请

你返回所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。

 示例 1：

输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
解释：
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
注意，输出的顺序和三元组的顺序并不重要。
示例 2：

输入：nums = [0,1,1]
输出：[]
解释：唯一可能的三元组和不为 0 。
示例 3：

输入：nums = [0,0,0]
输出：[[0,0,0]]
解释：唯一可能的三元组和为 0 。


*/

func threeSum(nums []int) [][]int {
	// sort.Ints(nums)
	// -1 + -1 + 2 = 0
	// 1= -1 + 2
	// 1 - (-1) =
	//  把三个数之和转换成一个固定的数据和两个数之和的形式计算
	// -1 + = 2 + -1
	res := [][]int{}
	for i := 0; i < len(nums)-1; i++ {
		tmpTarget := -1 * nums[i]
		tmpMap := make(map[int]int)
		// tmpResult := []int{nums[i]}
		for j := i + 1; j < len(nums)-1; j++ {
			tmp := tmpTarget - nums[j]
			// 找到符合的数据了，返回索引
			if index, ok := tmpMap[tmp]; ok {
				// 找到一组就添加到结果中。
				res = append(res, []int{nums[i], nums[j], nums[index]})
			}
			tmpMap[nums[j]] = j
		}

	}
	return res

}

// test
func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	t.Log(threeSum(nums))
}
