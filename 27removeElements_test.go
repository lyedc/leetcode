package leetcode

import (
	"fmt"
	"testing"
)

/*
LeetCode 第 27 题是 "Remove Element"，题目描述如下：

给定一个数组 nums 和一个值 val，你需要原地移除所有数值等于 val 的元素，使得每个新元素都无需改变其相对顺序。返回移除后数组的新长度，不需要考虑数组中超出新长度后面的元素。

不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。

给定 nums = [3,2,2,3], val = 3,
函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。
你不需要考虑数组中超出新长度后面的元素。

给定 nums = [0,1,2,2,3,0,4,2], val = 2,
函数应该返回新的长度 5, 并且 nums 中的前五个元素为 0, 1, 3, 0, 4。
注意这五个元素可为任意顺序。
你不需要考虑数组中超出新长度后面的元素。


*/
// 采用左右指针的方式,左边找等于val的值,右边找不等于val,l和r相等结束.这样val就都会跑到右边.
func RemoveElements(intList []int, val int)(int, []int)  {
	if len(intList) == 0{
		return 0, nil
	}
	l,r := 0, len(intList) -1
	// 只要不相等就一直循环
	for l != r{
		// 从左边找到第一个等于val的值 l不能大于r,只能小于等于r,判断当前值是什么
		// 左边和右边只能有一个计算中间位置...否则就会造成多加或者多减的操作.造成死循环
		for intList[l] != val && l <=r {
			l ++
		}
		// 从右边找到第一个不等于val的值, r只能大于l,如果等于了的话,就会多判断一次,导致r的值减小,多减了一次
		for intList[r] == val && l<r{
			r --
		}
		// 双方交换位置
		intList[l], intList[r] = intList[r], intList[l]
	}
	if intList[l] == val{
		return l, intList
	}
	return l, intList
}

// 采用快慢指针的方式,
func RemoveElements2(intList []int, val int)(int, []int)  {
	if len(intList) == 0{
		return 0, nil
	}

	l := 0  // 慢指针
	// 快指针
	for r := 0 ;r<len(intList);r++{
		// 只要不等于val就把值放到慢指针指向的位置,也就是从数组的开头替换
		if intList[r] != val{
			// 只要不相等就可以进行数据的交换,这样val目标的数据都会被移动到后面
			intList[l], intList[r] = intList[r], intList[l]
			l ++
		}
	}
	return l, intList
}

func TestRemoveElements(t *testing.T) {
	l1 := []int{3,2,3,2,3}
	val := 3
	len1, result := RemoveElements(l1, val)
	fmt.Println(result)
	fmt.Println(len1)
	l2 := []int{0,1,2,2,3,0,4,2}
	val = 2
	len2, result2 := RemoveElements(l2, val)
	fmt.Println(len2)
	fmt.Println(result2)
}

func TestRemoveElements2(t *testing.T) {
	l1 := []int{3,2,3,2,3}
	val := 3
	len1, result := RemoveElements2(l1, val)
	fmt.Println(result)
	fmt.Println(len1)
	l2 := []int{0,1,2,2,3,0,4,2}
	val = 2
	len2, result2 := RemoveElements2(l2, val)
	fmt.Println(len2)
	fmt.Println(result2)
}