package leetcode

import (
	"fmt"
	"strings"
	"testing"
)

// 217. 存在重复元素

/*
Given an array of integers, find if the array contains any duplicates.

Your function should return true if any value appears at least twice in the array, and it should return false if every element is distinct.

Example 1:

Input: [1,2,3,1] Output: true

Example 2:

Input: [1,2,3,4] Output: false

Example 3:

Input: [1,1,1,3,3,4,3,2,4,2] Output: true

*/

func containsDuplicate(nums []int) bool {
	tmpMap := make(map[int]bool)
	for _, item := range nums {
		if value := tmpMap[item]; value {
			return true
		}
		tmpMap[item] = true
	}
	return false
}

// test

func TestContainsDuplicate(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	res := containsDuplicate(nums)
	t.Log(res)

	nums = []int{1, 2, 3, 4}
	res = containsDuplicate(nums)
	t.Log(res)

	nums = []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	res = containsDuplicate(nums)
	t.Log(res)
}

// 389 找不同
/*
给定两个字符串 s 和 t ，它们只包含小写字母。

字符串 t 由字符串 s 随机重排，然后在随机位置添加一个字母。

请找出在 t 中被添加的字母。



示例 1：

输入：s = "abcd", t = "abcde"
输出："e"
解释：'e' 是那个被添加的字母。
示例 2：

输入：s = "", t = "y"
输出："y"
*/
func findTheDifference(s string, t string) string {

	sTmp := strings.Split(s, "")
	tTmp := strings.Split(t, "")
	// 这里使用int，是防止出现添加的是相同的字符，所以来计算每个字符出现的次数
	tmpMap := make(map[string]int)
	for _, item := range sTmp {
		// if _, ok := tmpMap[item]; ok {
		// 	tmpMap[item] += 1
		// } else {
		// 	tmpMap[item] = 1
		// }
		// 下面的更加简洁。。。
		tmpMap[item] += 1
	}
	for _, item := range tTmp {
		// 如果没有出现，直接返回
		if value, ok := tmpMap[item]; !ok {
			return item
		} else {
			// 如果出现了，并且次数已经为0了，返回当前值,表示添加的数据是一个重复的数据。
			// 这里一定是小于0 这样才能表示出现的次数比原来的多，表示添加了一重复的数据
			if value-1 < 0 {
				return item
			}
			tmpMap[item] = value - 1
		}
	}
	return ""

}

// test
func Test_findTheDifference(g *testing.T) {
	s := "abcd"
	t := "abcda"
	fmt.Println(findTheDifference(s, t))

	s = ""
	t = "a"
	fmt.Println(findTheDifference(s, t))
}

// 496 下一个更大元素 这个已经在496中实现了，采用map加上栈的方式实现。。。
