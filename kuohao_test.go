package leetcode

import (
	"fmt"
	"strings"
	"testing"
)

/*
题目 20: 有效的括号 (Valid Parentheses)

给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:

输入: "()"
输出: true

输入: "()[]{}"
输出: true

输入: "(]"
输出: false

*/

type NewStack struct {
	list []string
}

func ( stack *NewStack)Push( item string)  {
	stack.list = append(stack.list, item)
}

func (stack *NewStack)Pop()string  {
	v := stack.list[len(stack.list)-1]
	stack.list = stack.list[:len(stack.list) - 1]
	return v
}

func (stack *NewStack)IsEmpty()bool  {
	return len(stack.list) == 0
}

func kuohao(item string)bool  {
	stack := NewStack{[]string{}}
	list := strings.Split(item, "")
	for _, value := range list{
		switch value {
		case "{":
			stack.Push(value)
		case "[":
			stack.Push(value)
		case "(":
			stack.Push(value)
		case "}":
			// 表示匹配上了继续下一个
			if "{" == stack.Pop(){
				continue
			}
			// 没有匹配上
			return false
		case "]":
			if "[" == stack.Pop(){
				continue
			}
			// 没有匹配上
			return false
		case ")":
			if "(" == stack.Pop(){
				continue
			}
			// 没有匹配上
			return false
		}
	}
	// 没有完全匹配
	if len(stack.list) !=0{
		return false
	}
	return true
}

func TestKuohao(t *testing.T) {
	demo1 := "{{}}[]()"
	fmt.Println(kuohao(demo1))

	demo2 := "{{}}[()"
	fmt.Println(kuohao(demo2))

	demo3 := "{{}}[()]"
	fmt.Println(kuohao(demo3))
}