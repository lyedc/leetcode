package host100

import (
	"math"
	"testing"
)

// leetcode 53 最大子数组和

/*
给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

子数组
是数组中的一个连续部分


示例 1：

输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
示例 2：

输入：nums = [1]
输出：1
示例 3：

输入：nums = [5,4,-1,7,8]
输出：23

如果维护了一个最大的前缀值得话，当出现负值的时候，就会出现 [1. -1] 的情况，导致计算结果出错。
	nums = []int{-1}
	// 测试结果
	result = maxSubArray(nums)
	t.Log(result)
https://leetcode.cn/problems/maximum-subarray/solutions/2533977/qian-zhui-he-zuo-fa-ben-zhi-shi-mai-mai-abu71/
*/

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	prefixSum := 0
	minPrefixSum := 0
	// 因为已经维护了最小的值为0，如果出现负值的情况况，相减能获取到最大的值
	// 如果没有负值的情况下，永远是前n项和最大。
	// maxprefixSum := 0
	result := math.MinInt
	for _, num := range nums {
		prefixSum += num
		// 计算最小的前缀和。
		// 如果没有判断数组的数量，那么这里就需要放到下面计算，当只有一个数据的时候，就没有最小的计算了。
		if prefixSum < minPrefixSum {
			minPrefixSum = prefixSum
		}
		// if prefixSum > maxprefixSum {
		// 	maxprefixSum = prefixSum
		// }
		// 取结果，和最大前缀-最小前缀的值得最大的一个。
		if prefixSum-minPrefixSum > result {
			result = prefixSum - minPrefixSum
		}
		// else {
		// 	result = prefixSum
		// }

	}
	return result
}

// test
func TestPuTongShuZu(t *testing.T) {
	// 测试用例
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	// 测试结果
	result := maxSubArray(nums)
	t.Log(result)

	nums = []int{5, 4, -1, 7, 8}
	// 测试结果
	result = maxSubArray(nums)
	t.Log(result)

	nums = []int{-1}
	// 测试结果
	result = maxSubArray(nums)
	t.Log(result)

	nums = []int{1, 2, 3, 4, 5}
	// 测试结果
	result = maxSubArray(nums)
	t.Log(result)

}
