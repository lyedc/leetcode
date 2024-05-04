package host100

import (
	"fmt"
	"math"
	"testing"
)

// leetcode 打家劫舍
func rob(nums []int) int {
	if len(nums) == 0{
		return 0
	}
	if len(nums) == 1{
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i<len(nums); i++{
		// // 抢当前房子或不抢当前房子,如果抢了当前房子就只能获取到dp[i -2] + 当前房子中的金额
		// 如果不抢当前房子那么就直接获取dp[i-1]
		// 所以获取两种情况的最大值就可以了。

		dp[i] = max(dp[i-1], dp[i-2]+ nums[i])
	}
	return dp[len(nums)-1]
}

// leetcode  完全平方数 使用go 语言实现，并给出详细的解释
func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i<n;i++{
		dp[i] = n +1
	}
	dp[0] = 0
	for i := 1; i<n; i++{
		for j :=1; j< i ; j++{
			if j*j <i{
				// i - j*j 表示的是通过j*j的平方后，还需要多少中组合才能完成数据的组成
				// +1 表示的是加上 j*j的这种方案。
				dp[i] = min(dp[i], dp[i-j*j] +1)
			}
		}
	}
	return dp[n]
}

// leetcode  零钱兑换 使用go 语言实现，并给出详细的解释
// 有多少种组成方案的关键是: dp[i] = min(dp[i], dp[i-coins[j] +1])
// i - coins[j] 表示使用了j后，还需要多少情况才能组成，
// +1 表示的是要加上 coins[j] 这种方案
func coinChange(coins []int, amount int) int {
	// 核心原理是动态规划
	// 为什么需要两个循环,以为是金额的循环，第二个循环是使用多少硬币，能兑换出指定的金额。
	// 创建一个长度为amount+1 的切片 dp，用于存储每个金额对应的最小硬币数
	// // 初始值设为 math.MaxInt32，这是一个不可能的值，用于后续的比较和更新
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
	}
	// 表示0 金额的组成方式只有0个硬币
	dp[0] = 0
	// 循环要兑换的金额
	for i := 1; i< amount; i++{
		// 循环每种金额需要多少硬币才能兑换
		for j :=0 ; j<len(coins); j ++{
			if coins[j] <= i{
				// dp[i-coins[j] +1] 表示使用了con[j]个硬币后，还需要多少的硬币采可以组成指定的金额
				// +1 表示的是 coins[j]这个硬币(这个组成是需要加入到总的组成数中的。)
				dp[i] = min(dp[i], dp[i-coins[j] +1])
			}
		}
	}
	// 表示无法组成指定的金额，返回-1
	// 没有办法组成指定的金额的话，dp中的值就不会被更新
	if dp[amount] == math.MaxInt32{
		return -1
	}
	// 否则就返回指金额的组成部分
	return dp[amount]
}



// leetcode 单词拆分 使用go 语言实现，并给出解释
func wordBreak(s string, wordDict []string) bool {
	// 创建一个长度为 len(s)+1 的切片 dp，用于存储字符串前 i 个字符是否可以拆分
	// 初始值设为 false
	dp := make([]bool, len(s)+1)
	// base case，空字符串可以拆分
	dp[0] = true

	// 创建一个单词集合，用于快速判断单词是否在字典中
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}

	// 动态规划求解
	for i := 1; i <= len(s); i++ {
		// 对于每个子字符串 s[:i]，尝试拆分成 s[:j] 和 s[j:i]
		for j := 0; j < i; j++ {
			// 如果前 j 个字符可以拆分，并且 s[j:i] 是字典中的单词，则更新 dp[i]
			// 前j个字符被拆分成了s[:j]和s[j:i]，且s[j:i]是字典中的单词，所以dp[i]为true
			// 因为只要有一个符合要求的拆分，就可以认为整个字符串可以拆分，所以dp[i]为true
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	// dp[len(s)] 表示整个字符串是否可以拆分
	return dp[len(s)]
}


// leetcode 最长递增子序列 使用go 语言实现，并给出解释
func lengthOfLIS(nums []int) int {
	// 创建一个长度为 len(nums) 的切片 dp，用于存储以 nums[i] 结尾的最长递增子序列的长度
	// 初始值都设为 1，因为每个数字本身就是一个长度为 1 的递增子序列
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}

	// 动态规划求解
	for i := 1; i < len(nums); i++ {
		// 对于每个数字 nums[i]，遍历在它之前的所有数字 nums[j]
		for j := 0; j < i; j++ {
			// 如果 nums[i] > nums[j]，则尝试更新 dp[i]
			// 意思就是当前的数字num[i]都大于前面的每一个数字，那么就可以将dp[i]更新为dp[j]+1
			if nums[i] > nums[j] {
				// dp[j] 表示以 nums[j] 结尾的最长递增子序列的长度
				// dp[i] 取 dp[i] 和 dp[j]+1 的较大值
				// dp[j]表示的是第二次 循环中，dp[j]的值，如果dp[j]的值大于dp[i]，那么dp[i]的值就等于dp[j]+1
				// 相当于是每次循环都是更新 dp[i] 的值，直到dp[i]的值大于dp[j]+1
				dp[i] = max(dp[i], dp[j]+1)
				//if nums[i] == 101{
				//	fmt.Println(nums[j])
				//	//fmt.Println(dp[i])
				//}
			}
		}
	}

	// 遍历 dp 数组，找到最大的值，即为最长递增子序列的长度
	maxLength := 0
	for i := range dp {
		maxLength = max(maxLength, dp[i])
	}

	return maxLength
}

func TestLength(t *testing.T){
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Printf("The length of the longest increasing subsequence is %d\n", lengthOfLIS(nums))
}


//leetcode 152 乘积最大子数组
func maxProduct(nums []int) int {
	max ,min, maxproduct := nums[0], nums[0], nums[0]
	for i :=1 ; i < len(nums); i++{
		// 这里需要考虑如果是负数的时候，最小就变成了最大值了，最大值就变成最小值了。
		if nums[i] < 0{
			max, min = min, max
		}
		newMax := max * nums[i]
		newMin := min * nums[i]
		max = int(math.Max(float64(newMax), float64(nums[i])))
		min = int(math.Min(float64(newMin), float64(nums[i])))
		maxproduct = int(math.Max(float64(maxproduct), float64(max)))
	}
 	return maxproduct
}

// leetcode 416 分割等和子集
/*
为了解决这个问题，我们可以使用一个类似于“0-1背包问题”的动态规划方法。
我们需要找到一个子集，其总和为整个数组总和的一半。如果找到了这样的子集，
那么另一个子集也必然存在，并且和第一个子集的和相等。
*/
func canPartition(nums []int) bool {
	sum := 0
	// 先计算数组的总和
	for i :=0; i<len(nums); i++{
		sum += nums[i]
	}
	// 如果数组的综合是一个奇数 则不可能分割成两个相等的子集
	if sum % 2 != 0{
		return false
	}
	target := sum / 2
	// target 是我们要找到的目标子数组的和
	// 创建一个长度为target+1的数组，用于存储每个目标值是否可以找到一个子集
	dp := make([]bool, target+1)
	// 初始化dp数组，将dp[0]设为true，表示空集可以找到一个子集
	dp[0] = true
	// 遍历数组中的每个元素
	for i := 0; i < len(nums); i++ {
		// 遍历每个目标值
		for j := target; j >= nums[i]; j--{
			// 如果当前元素可以加入到子集，则更新dp[j]为true
			//  dp[j-nums[i]] 表示的是dp[j-nums[i]]为true，表示可以找到一个子集，并且这个子集的和为j-nums[i]
			// 后面的表示 j-nums[i]剩下和是否存在一个这样的子集，也就是计算了每个dp的可能性。
			dp[j] = dp[j] || dp[j-nums[i]]
		}
	}
	return dp[target]
}