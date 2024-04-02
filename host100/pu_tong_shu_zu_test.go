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

// 思路： 因为是计算的最大子数组，可以利用前缀和，前缀和可以计算出子数组的和，利用当前的前缀和减去之前的前缀和，得到当前子数组的和，然后比较
// 但是需要注意不能维护一个最大的前缀和，因为前缀和是递增的，最大的前缀和永远是当前的值，所以需要维护一个最小的前缀和，这样减去最小的前缀和，就可以得到当前子数组的和，然后比较
// 同时需要注意只有一个数据的情况下。

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

// leetcode 56 合并区间

/*
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。



示例 1：

输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2：

输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]
解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。

*/
// 思路：
// 1. 先对数组进行排序
// 2. 遍历数组，如果当前区间的起始值小于等于前一个区间的结束值，则将前一个区间的结束值更新为当前区间的结束值
// 3. 如果当前区间的起始值大于前一个区间的结束值，则将当前区间添加到结果数组中
// 4. 返回结果数组
func merge(intervals [][]int) [][]int {
}
