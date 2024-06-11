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

// 438 找到字符串中所有字母异位词
/*
给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。

异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。



示例 1:

输入: s = "cbaebabacd", p = "abc"
输出: [0,6]
解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。
 示例 2:

输入: s = "abab", p = "ab"
输出: [0,1,2]
解释:
起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。


findAnagrams函数接受两个字符串参数s和p，返回一个整数切片，其中包含s中所有p的字母异位词的起始索引。

我们首先创建一个空的整数切片result用于存储结果。

如果s的长度小于p的长度，那么不可能存在p的字母异位词，直接返回空结果。

创建一个map pCounter 用于记录p中每个字符的出现次数。遍历p字符串，统计每个字符出现的次数。

初始化滑动窗口的起始位置 start 和结束位置 end，以及计数器 counter， counter 表示当前窗口内包含的p中字符的个数。

进入主要的滑动窗口循环。循环的结束条件是end达到s字符串的末尾。

在每次循环中，首先检查s[end]字符是否在p中，如果是则counter减一。

更新pCounter中s[end]字符的计数，因为它已经在窗口内了。

将end向右移动一位。

检查counter是否为0，如果为0表示窗口内包含了p的全部字符，这时我们将当前窗口的起始位置start添加到结果中。

检查当前窗口的大小是否等于p的长度，如果等于则说明窗口已经包含了p的长度，需要移动start。

如果s[start]在p中，那么我们需要将counter加一，因为我们将要移除窗口左边的这个字符。

更新pCounter中s[start]字符的计数，因为它将要被移出窗口了。

将start向右移动一位，缩小窗口。

最后返回结果切片result，其中包含了所有符合条件的字母异位词的起始索引

*/

func findAnagrams(s string, p string) []int {
	var result []int

	if len(s) < len(p) {
		return result
	}

	pCounter := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		// 记录p字符串每个字母出现的次数，记录的是字符的ascii码
		pCounter[p[i]]++
	}
	// 初始化滑动窗口的起始和结束索引，以及counter用于记录当前窗口内还需要找到的字符数。

	// 开始滑动窗口的计算方案
	start := 0
	end := 0
	counter := len(p)

	for end < len(s) {
		// 如果当前窗口内的字符在p中出现过，则counter减一。

		// 判断窗口内每个字符在p字符串出现的次数，如果大于0，说明是异位词，counter减1
		if pCounter[s[end]] > 0 {
			// counter 是p字符的长度，如果counter减到0，说明找到了一个异位词
			//counter 表示当前窗口内包含的p中字符的个数
			counter--
		}
		// end位置的字符在p字符串中出现的次数减一，因为end位置的字符已经不在窗口内了
		// 更新pCounter中当前字符的计数
		pCounter[s[end]]--
		// 开始判断字符中的第二个位置。
		// 将end索引向右移动，扩大窗口。
		end++
		// 表示在一个窗口内已经比较完了。刚好是一个异位词。那么就记录结果。
		if counter == 0 {
			result = append(result, start)
		}
		// 如果窗口的长度等于p的长度，需要移动start索引，缩小窗口。

		// 表示窗口内字符的数量已经达到p字符串的长度，需要移动start位置，
		// 也就是说一个滑动窗口已经满了，需要开始滑动了。。
		if end-start == len(p) {
			// 如果移动start后，窗口内移除的字符在p中出现次数仍大于等于0，则counter加一。

			if pCounter[s[start]] >= 0 {
				counter++
			}
			// 更新pCounter中移动start索引对应的字符计数。
			// 逆向更新计数：当窗口向右滑动时，意味着原本位于窗口内的字符s[start]现在要移出窗口。由于之前在进入窗口时我们对s[start]进行了计数减一的操作（pCounter[s[end]]--），现在它离开窗口，我们需要“回退”这一操作，即对该字符的计数进行加一，以便准确反映不在当前窗口内时该字符应统计的次数
			pCounter[s[start]]++
			start++
		}
	}

	return result
}

// test
func TestFindAnagrams(t *testing.T) {
	s := "cbaebabacd"
	words := "abc"
	t.Log(findAnagrams(s, words))
	s = "abab"
	words = "ab"
	t.Log(findAnagrams(s, words))
}
