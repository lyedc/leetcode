package study

import (
	"math"
)

// 有 N家人，不能连续头两家，问最多能偷到多少钱
func dajiajieshe(num []int) int {
	dp := make([]int, len(num))
	if len(num) == 0 {
		return 0
	}
	if len(num) == 1 {
		return num[0]
	}
	dp[0] = num[0]
	// 这里如果连续偷两家，就只能取最大的值。
	dp[1] = max(num[0], num[1])
	for i := 2; i < len(num); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+num[i])
	}
	return dp[len(num)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// 给个整数n，用最少的完全平方数组成
func wanquanpingfangshu(n int) int {
	// dp[i] = min(dp[i], dp[i - j*j] +1)
	dp := make([]int, n+1)
	for i := 0; i < n; i++ {
		dp[i] = n + 1
	}
	// for i:=1;i<
	dp[0] = 0
	for i := 1; i < n; i++ {
		for j := 1; j < i; j++ {
			if j*j < i {
				//dp[i]上面已经赋值过了，所以不用再赋值了
				dp[i] = min(dp[i], dp[i-j*j]+1)
			}
		}
	}
	return dp[n]

}

// dui hua ling qian
func duihualingqian(num []int, n int) int {
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = math.MaxInt
	}
	dp[0] = 0
	for i := 1; i < n; i++ {
		for j := 0; j < len(num); j++ {
			if num[j] < i {
				//dp[i]上面已经赋值过了，所以不用再赋值了
				dp[i] = min(dp[i], dp[i-num[j]]+1)
			}
		}
	}
	if dp[n] == math.MaxInt {
		return -1
	}
	return dp[n]
}

func dancichaifen(s string, words []string) bool {
	dp := make([]bool, len(s)+1)
	wordMap := make(map[string]bool)
	for _, item := range words {
		wordMap[item] = true
	}
	dp[0] = true
	for i := 1; i < len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordMap[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

func zuichangdizeng(num []int) int {
	dp := make([]int, len(num))
	for i := 0; i < len(num); i++ {
		dp[i] = 1
	}
	dp[0] = 0

}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
