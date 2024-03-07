package leetcode

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

/*
LeetCode 49. 字母异位词分组

给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。

示例:

输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
输出:
[
["ate","eat","tea"],
["nat","tan"],
["bat"]
]

*/

//采用排序方式
// 对字符串进行切片,然后排序,然后存入到map中,最后转换
func GroupAnagrams(list []string)[][]string  {

	result := make(map[string][]string)
	for _, value := range list{
		tmp := strings.Split(value, "")
		sort.Strings(tmp)
		tmp2 := strings.Join(tmp, "")
		result[tmp2] = append(result[tmp2], value)
	}
	// 转换map为二位切片返回
	result2 := [][]string{}
	for _, item := range result{
		result2 = append(result2, item)
	}
	return result2
}

func TestGroupAnagrams(t *testing.T) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	result := GroupAnagrams(strs)
	for _, group := range result {
		fmt.Println(group)
	}
}