package leetcode

import (
	"fmt"
	"testing"
)

/*
原题：

给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

你可以按任意顺序返回答案。

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
*/

// 使用哈希表的方式,map中存储的是列表数据的value和index， for循环list，然后taget-value = 期待的值，从map中找其他的值，然后获取索引
func twoNumberSum(nums []int, target int) []int  {
	indexMap := make(map[int]int)
	for index,value := range nums{
		two := target - value
		if oneIndex, ok := indexMap[two]; ok{
			return []int{index, oneIndex}
		}
		indexMap[value] = index
	}
	return []int{}
}

// 暴力破解
func twoNumberSumBaoli(nums []int, target int) []int  {
	for i :=0; i< len(nums); i++  {
		// 初始值大于了target就直接下一次
		if  nums[i] >= target{
			continue
		}
		for j := i + 1 ; j<len(nums);j++{
			// 初始值大于了target就直接下一次
			if nums[j] >= target{
				continue
			}
			if nums[j] + nums[i] == target{
				return []int{i, j}
			}
		}
	}
	return nil
}


func Test_twoNumberSum(t *testing.T)  {
	nums := []int{2, 7, 3, 6}
	target := 9
	result := twoNumberSum(nums, target)
	fmt.Println(result)

	result2 := twoNumberSumBaoli(nums, target)
	fmt.Println(result2)
}
