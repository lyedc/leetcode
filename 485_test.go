package leetcode

import (
	"fmt"
	"testing"
)

// 485. 最大连续1的个数

func FindMaxConsecutiveOnes(nums []int) int {
	max := 0
	count := 0
	for _, v := range nums {
		if v == 1 {
			count++
		} else {
			if count > max {
				max = count
			}
			count = 0
		}
	}
	if count > max {
		max = count
	}
	return max
}

func FindMaxConsecutiveOnes2(nums []int) int {
	result := 0
	tmpResult := 0

	for _, item := range nums {
		// 只要等于1就增加临时变量，如果不等于1就表示不连续了，就重置临时变量
		// 每次增加临时变量，如果临时变量大于结果，就重置结果
		if item == 1 {
			tmpResult++
			if tmpResult > result {
				result = tmpResult
			}
		} else {
			tmpResult = 0
		}
	}

	// for i := 0; i < len(nums)-1; i++ {
	// 	if nums[i] == nums[i+1] {
	// 		tmpResult++
	// 		if tmpResult > result {
	// 			result = tmpResult
	// 		}
	// 	} else {
	// 		tmpResult = 0
	// 	}
	// }
	// if tmpResult > result {
	// 	result = tmpResult
	// }
	return result
}

// test
func TestFindMaxConsecutiveOnes(t *testing.T) {
	nums := []int{1, 1, 0, 1, 1, 1}
	//findMaxConsecutiveOnes(nums)
	fmt.Println(nums)
	fmt.Println(FindMaxConsecutiveOnes(nums))
	fmt.Println("==========")
	fmt.Println(FindMaxConsecutiveOnes2(nums))
}
