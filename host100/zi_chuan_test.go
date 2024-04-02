package host100

import (
	"fmt"
	"testing"
)

// leetcode 560 和为k的子数组
/*
给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。

子数组是数组中元素的连续非空序列。



示例 1：

输入：nums = [1,1,1], k = 2
输出：2
示例 2：

输入：nums = [1,2,3], k = 3
输出：2


前缀和（Prefix Sum）是一种常见的技巧，用于快速计算数组中某个范围内元素的和。前缀和数组中的每个元素表示从数组开始到当前位置的所有元素的总和。

例如，给定数组 nums = [1, 2, 3, 4, 5]，它的前缀和数组为 prefixSum = [1, 3, 6, 10, 15]。其中 prefixSum[i] 表示数组 nums 中前 i+1 个元素的和。

前缀和的计算方式如下：

初始化一个额外的数组 prefixSum，其长度为原数组长度加1。
将 prefixSum[0] 设为0，表示前0个元素的和为0。
遍历原数组，依次计算累积和，并存储到 prefixSum 中。
利用前缀和，可以在O(1)的时间复杂度内计算出任意范围内元素的和。假设要计算数组 nums 中下标 i 到下标 j（包括 i 和 j）之间的元素的和，只需使用 prefixSum[j+1] - prefixSum[i] 即可。

前缀和的一个主要应用是在子数组问题中，可以快速地求解子数组的和、平均值等问题

*/
func subarraySum(nums []int, k int) int {
	count := 0
	sumMap := make(map[int]int)
	// 其中 prefixSum[i] 表示数组 nums 中前 i+1 个元素的和
	prefixSum := 0
	// 1 + 1= 2
	sumMap[0] = 1 // 初始化前缀和为0的出现次数为1

	for _, num := range nums {
		prefixSum += num
		// 意味的是： prefixSum[j] - k = prefixSum[i], 那么就表示 i:j之间的数据的和为k
		// 计算之前的前缀和，也就是  prefixSum[j] - profixSum[i] = k是否存在，j表示后面的前缀和，i表示前面的前缀和。
		count += sumMap[prefixSum-k]
		// 前缀和存在次数++， 存储起来每次的前缀和
		sumMap[prefixSum]++
	}

	return count
}

// test

func TestSubarraySum(t *testing.T) {
	nums := []int{1, 1, 1}
	k := 2
	count := subarraySum(nums, k)
	println(count)

	nums2 := []int{1, 2, 3}
	k = 3
	count = subarraySum(nums2, k)
	println(count)
}

// leetcode 239 滑动窗口最大值

/*

给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

返回 滑动窗口中的最大值 。
输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
输出：[3,3,5,5,6,7]

针对单调递减的解释：
在这个问题中，我们维护一个单调递减的窗口，意味着窗口内的元素应该按照递减的顺序排列。这样做的目的是为了使得窗口的最大值始终位于窗口的最左侧。

让我们来看一下为什么维护一个单调递减的窗口是有用的：

当我们将一个新的元素添加到窗口中时，我们希望尽可能地将较小的元素从窗口中删除。因为这些较小的元素不太可能成为最大值。如果我们发现新元素比窗口中的某些元素更大，我们可以将这些较小的元素从窗口中移除。

维护一个单调递减的窗口可以保证窗口中的最大值始终位于窗口的最左侧。这样，当我们需要记录当前窗口的最大值时，只需要检查窗口的最左侧元素即可。

在代码中，我们使用一个切片 window 来表示窗口。每当我们遇到一个新的元素时，我们将其索引添加到 window 中。然后，我们从窗口的末尾开始，删除所有比当前元素小的元素，直到窗口内的元素保持单调递减。最终，窗口的最左侧元素即为当前窗口的最大值。

*/

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}
	result := []int{}
	windown := make([]int, 0)
	for i, num := range nums {
		// // 维护窗口，保持窗口大小为k，并且窗口中的元素单调递减
		for len(windown) > 0 && nums[windown[len(windown)-1]] < num {
			// 删除滑动窗口中所有比新进入的值小的值，这样保持最大的值，一直是在滑动窗口的最左侧
			windown = windown[:len(windown)-1]
		}
		windown = append(windown, i)
		// 当滑动窗口的左边界大于窗口长度，则移除窗口左边界的元素
		// 也就是 i + 0 > k 表示已经出了滑动窗口了
		// // 移除窗口左边界元素，如果它已经不在当前窗口内
		// 表示当前的所有 j,和之前的索引i的差值，j - i = k 表示滑动窗口的长度。等于长度的时候，就要去掉最左边的。
		// 例如：i = 2 j = 5 那么i 和 J之间已经存储了4个数字了。超过滑动窗口的长度，所以要移除最左边的。
		if windown[0] == i-k {
			windown = windown[1:]
		}
		// 表示窗口已经达到大小K，记录当前窗口的最大值
		// 表示已经达到一个滑动窗口了，需要取出窗口中的最大值。
		if i >= k-1 {
			result = append(result, nums[windown[0]])
		}
	}
	return result

}

// test
func TestMaxSlidingWindow(t *testing.T) {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	result := maxSlidingWindow(nums, k)
	fmt.Println(result)
}
