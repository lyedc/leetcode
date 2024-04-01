package host100

import (
	"testing"
)

// leetcode 3

/*

给定一个字符串 s ，请你找出其中不含有重复字符的 最长
子串的长度。
示例 1:

输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
*/

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	start := 0                    // 滑动窗口的起始位置
	maxLength := 0                // 最长子串的长度
	charMap := make(map[byte]int) // 字符映射表，用于记录每个字符最近出现的位置

	for end := 0; end < len(s); end++ {
		char := s[end]

		// 如果当前字符已经在 charMap 中，且其位置大于等于 start（即在当前窗口内）
		if lastIndex, ok := charMap[char]; ok && lastIndex >= start {
			// 更新 start 为重复字符的下一个位置
			// 也就是滑动窗口的起始位置
			start = lastIndex + 1
		}

		// 更新当前字符的最新位置
		charMap[char] = end

		// 更新最长子串的长度
		// end 因为是数组的长度，所以真实的物理长度是需要增加1的。
		if end-start+1 > maxLength {
			maxLength = end - start + 1
		}
	}

	return maxLength
}

func TestHuaDong(t *testing.T) {
	s := "abcabcbb"
	t.Log(lengthOfLongestSubstring(s))
	// s = "bbbbb"
	// t.Log(lengthOfLongestSubstring(s))
	// s = "pwwkew"
	// t.Log(lengthOfLongestSubstring(s))
}
