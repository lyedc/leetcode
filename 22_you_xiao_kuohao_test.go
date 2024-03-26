package leetcode

import (
"fmt"
"testing"
)

/*
题目描述
给定一个整数 n，生成所有有效的表达式，其中包含 n 个左括号和 n 个右括号。

示例：

输入：n = 3
输出：[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]

*/

func GenerateParenthesis(n int) []string {
	var result []string
	backtrack(&result, "", 0, 0, n)
	return result
}

func backtrack(result *[]string, current string, open int, close int, max int) {
	if len(current) == max*2 {
		*result = append(*result, current)
		return
	}

	if open < max {
		backtrack(result, current+"(", open+1, close, max)
	}
	if close < open {
		backtrack(result, current+")", open, close+1, max)
	}
}


func TestGenerateParenthesis(t *testing.T) {
	n := 3
	result := GenerateParenthesis(n)
	for _, val := range result {
		fmt.Println(val)
	}
}