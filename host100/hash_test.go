package host100

import (
	"sort"
	"strings"
	"testing"
)

// leetcode 1 两数之和
func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return nil
	}
	tmpMap := make(map[int]int)
	for index, item := range nums {
		tmp := target - item
		if tmpIndex, ok := tmpMap[tmp]; ok {
			return []int{index, tmpIndex}
		}
		tmpMap[item] = index
	}
	return nil
}

func TestTwoHumbers(t *testing.T) {
	t.Log(twoSum([]int{2, 7, 11, 15}, 9))
	t.Log(twoSum([]int{2, 7, 11, 15}, 26))
}

// 49 字母异位词分组

/*
给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。

字母异位词 是由重新排列源单词的所有字母得到的一个新单词。

示例 1:

输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
示例 2:

输入: strs = [""]
输出: [[""]]
示例 3:

输入: strs = ["a"]
输出: [["a"]]
*/

func groupAnagrams(strs []string) [][]string {
	tmpMap := make(map[string][]string)
	result := [][]string{}
	for _, item := range strs {
		tmpStr := strings.Split(item, "")
		sort.Strings(tmpStr)
		newItem := strings.Join(tmpStr, "")
		tmpMap[newItem] = append(tmpMap[newItem], item)
	}
	for _, value := range tmpMap {
		result = append(result, value)
	}
	return result
}

func Test_GroupAnagrams(t *testing.T) {
	t.Log(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	t.Log(groupAnagrams([]string{"", ""}))
	t.Log(groupAnagrams([]string{"a"}))
}

// 128 最长连续序列

/*
给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。

请你设计并实现时间复杂度为 O(n) 的算法解决此问题。



示例 1：

输入：nums = [100,4,200,1,3,2]
输出：4
解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
示例 2：

输入：nums = [0,3,7,2,5,8,4,6,0,1]
输出：9
*/

func longestConsecutive(nums []int) int {
	tmpMap := make(map[int]bool)
	for _, item := range nums {
		tmpMap[item] = true
	}
	result := 0
	for value := range tmpMap {
		// 判断相邻的数据是否在map中，如果存在那么就跳过这个数据了，因为在下面的循环中会判断连续的数据的。
		if tmpMap[value-1] {
			continue
		}
		// 这里为什么是1呢，因为即便没有连续的数据它本身的长度也是1
		tmpResult := 1
		currentNum := value
		// 表示连续的数据也在map中。
		// 拿着第一个数据，开始循环判断相邻的数据是否也存在于集合中。 current就是当前的数据，
		// for循环表示去判断集合中其他的字段，
		for tmpMap[currentNum+1] {
			tmpResult++
			currentNum++
		}
		// 当找到不在集合中的数据时候，就更新结果。
		if tmpResult > result {
			result = tmpResult
		}
	}
	return result
}

// test
func Test_LongestConsecutive(t *testing.T) {
	t.Log(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	t.Log(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
}
