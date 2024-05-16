package host100

// leetcode 35 搜索插入位置
/*
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。
如果目标值不存在于数组中，返回它将会被按顺序插入的位置
请必须使用时间复杂度为 O(log n) 的算法。
*/
func searchInsert(nums []int, target int) int {
	start, end := 0, len(nums)-1
	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return start
}

// leetcode 74 搜索二维矩阵
func searchMatrix(matrix [][]int, target int) bool {
	// m 表示多少行，n表示多少列。
	m, n := len(matrix), len(matrix[0])
	start, end := 0, m*n-1
	for start <= end {
		mid := start + (end-start)/2
		// 根据中间值，计算出对应的行和列
		cow, rol := mid/n, mid%n
		if matrix[cow][rol] == target {
			return true
		}
		if matrix[cow][rol] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return false
}

// leetcode 34 在排序数组中查找元素的第一个和最后一个位置
/*
给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。
请你找出给定目标值在数组中的开始位置和结束位置。

如果数组中不存在目标值 target，返回 [-1, -1]。

你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题
*/
func searchRange(nums []int, target int) []int {
	start, end := 0, len(nums)-1
	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			left, right := mid, mid
			for left >= 0 && nums[left] == target {
				left--
			}
			for right < len(nums) && nums[right] == target {
			}
		}
	}
	return nil
}
