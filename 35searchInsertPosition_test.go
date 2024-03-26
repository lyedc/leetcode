package leetcode

import (
	"fmt"
	"testing"
)

/*
leetcode 搜索插入位置
给定一个排序数组和一个目标值，在数组中找到目标值，如果找不到，则返回可以将其插入的位置以确保其仍然有序。

你可以假设数组中无重复元素。

示例 1:

输入: nums = [1,3,5,6], target = 5
输出: 2

示例 2:

输入: nums = [1,3,5,6], target = 2
输出: 1

示例 3:

输入: nums = [1,3,5,6], target = 7
输出: 4
*/

// 采用二分查找进行处理,如果大于中间值那么就在后半部分,如果小于中间值,就在前半部分,然后继续采用
// 二分查找的方式.左右指针不断变换前后的位置

func SearchInsertPosition(intList []int, target int)int  {
	if len(intList) == 0{
		return 0
	}
	left,right := 0, len(intList) -1
	for left<=right{
		mid := left + (right - left)/2
		if intList[mid] ==target{
			return mid
		}
		// 说明在钱半部分,这个时候right=mid
		if intList[mid] > target{
			right = mid - 1 // 中间的位置已经判断过了.再减去1
		}else{
			left = mid + 1
		}
	}
	return left // 这里返回left,或者right都一样
	
}

func TestSearchInsertPosition(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	target := 5
	result := SearchInsertPosition(nums, target)
	fmt.Println(result) // 输出: 2

	target = 2
	result = SearchInsertPosition(nums, target)
	fmt.Println(result) // 输出: 1

	target = 7
	result = SearchInsertPosition(nums, target)
	fmt.Println(result) // 输出: 4
}