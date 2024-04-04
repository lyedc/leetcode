package host100

import (
	"math"
	"sort"
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
// 思路： 排序，加上循环遍历
// 1. 先对数组进行排序
// 2. 遍历数组，如果当前区间的起始值小于等于前一个区间的结束值，则将前一个区间的结束值更新为当前区间的结束值
// 3. 如果当前区间的起始值大于前一个区间的结束值，则将当前区间添加到结果数组中
// 4. 返回结果数组

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func merge(intervals [][]int) [][]int {
	if len(intervals) == 1 {
		return intervals
	}
	// 按照区间的起始位置进行排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	// 先把第一个数据放入到结果集中,这个数据成为比较的基础对象。
	merged := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		// 判断后续的是否有和第一个重复的区间，如果没有重复的就直接添加到结果集中。如果存在
		// 重复的，合并重复的区间。
		pre := merged[len(merged)-1]
		curr := intervals[i]
		if pre[len(pre)-1] >= curr[0] {
			// 取pre的最后一个和 curr的最后一个的最大值
			pre[len(pre)-1] = max(pre[len(pre)-1], curr[len(curr)-1])
		} else {
			merged = append(merged, curr)
		}
	}
	return merged

}

// test
func TestMerge(t *testing.T) {
	// 测试用例
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	t.Log(merge(intervals))

	intervals2 := [][]int{{1, 4}, {4, 5}}
	t.Log(merge(intervals2))

}

// leetcode 189 轮转数组

/*
给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。



示例 1:

输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右轮转 1 步: [7,1,2,3,4,5,6]
向右轮转 2 步: [6,7,1,2,3,4,5]
向右轮转 3 步: [5,6,7,1,2,3,4]
示例 2:

输入：nums = [-1,-100,3,99], k = 2
输出：[3,99,-1,-100]
解释:
向右轮转 1 步: [99,-1,-100,3]
向右轮转 2 步: [3,99,-1,-100]

*/

func rotate(nums []int, k int) {
	if len(nums) <= k {
		return
	}
	// 反转数组
	/*

		示例：
		1， 2， 3， 4， 5， 6， 7，

		7， 6， 5, 4, 3, 2, 1

		5， 6， 7， 4， 3, 2， 1

		5， 6， 7， 1， 2， 3， 4，


	*/

	k %= len(nums) // 将 k 取模，防止 k 大于数组长度的情况
	// 取模表示的就是取K的真实值。余数，只要K不大于数组的长度的情况下，这个数据就是K的真实值。
	// 先反转全部的数据
	reverse(nums, 0, len(nums)-1)
	// 反转前K个数据
	reverse(nums, 0, k-1)
	// 反转剩余的数据
	reverse(nums, k, len(nums)-1)

}

func reverse(nums []int, start, end int) {
	// 对数组内部的数据进行两两交换，注意因为是 start 和 end同时在动，所以相当于是使用双向指针的方式完成的
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

// test
func TestRotate(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(nums, k)
	t.Log(nums)

	nums = []int{-1, -100, 3, 99}
	k = 2
	rotate(nums, k)
	t.Log(nums)
}

// leetcode 238 乘积数组
/*
给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。

题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。

请 不要使用除法，且在 O(n) 时间复杂度内完成此题。



示例 1:

输入: nums = [1,2,3,4]
输出: [24,12,8,6]
示例 2:


输入: nums = [-1,1,0,-3,3]
输出: [0,0,9,0,0]
、

因为是要获取除了这个数之外所有的数据的乘积，所以可以演变成每个数左侧的乘积，* 每个数右侧的乘积，然后得到这个数据
之外所有数据的乘积了。。

创建一个长度为n的结果数组 result，用于存储最终的乘积结果。
首先，从左向右遍历一次数组，依次计算每个元素左侧的乘积，并存储在 result 中。
然后，从右向左遍历一次数组，依次计算每个元素右侧的乘积，并与之前存储的左侧乘积相乘，得到最终结果。
返回结果数组 result

第一次遍历：计算每个元素左侧所有元素的乘积。我们使用一个额外的数组或者变量来存储每个元素左侧所有元素的乘积。例如，对于数组 nums = [1, 2, 3, 4]，左侧乘积数组 leftProduct 为 [1, 1, 2, 6]，其中 leftProduct[i] 表示 nums[i] 左侧所有元素的乘积。

第二次遍历：计算每个元素右侧所有元素的乘积，并与第一次遍历得到的左侧乘积相乘。我们同样使用一个额外的数组或者变量来存储每个元素右侧所有元素的乘积。然后将右侧乘积数组与左侧乘积数组相乘即可得到最终的结果。

*/

func productExceptSelf(nums []int) []int {
	n := len(nums)
	result := []int{}
	left := 1
	// 先计算左边的乘积
	for i := 0; i < n; i++ {
		result[i] = left
		left *= nums[i]
	}
	// 在计算右边的乘积，并乘以最终的结果。
	right := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= right
		right *= nums[i]
	}

	return result
}

// test
func TestProductExceptSelf(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	result := productExceptSelf(nums)
	t.Log(result)

	nums = []int{-1, 1, 0, -3, 3}
	result = productExceptSelf(nums)
	t.Log(result)
}
