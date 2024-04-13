package leetcode

import (
	"fmt"
	"testing"
)

/*
第6题是“合并两个有序链表”（Merge Two Sorted Lists）。

题目描述如下：

将两个升序链表合并为一个新的升序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例：

输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]

输入：l1 = [], l2 = []
输出：[]

输入：l1 = [], l2 = [0]
输出：[0]

提示：

两个链表的节点数目范围是 [0, 50]
-100 <= Node.val <= 100
l1 和 l2 均按非递减顺序排列
*/

type node struct {
	value int
	next *node
}

func mergeTwoLists(node1,node2 *node)*node  {
	//result := &node{}
	// 如果node1遍历完了，那么就返回node2，直接连接到链表后面即可
	if node1 == nil{
		return node2
	}
	// 如果node遍历完了，直接返回node1，直接连接到链表后面即可。
	if node2 == nil{
		return node1
	}
	if node1.value <= node2.value{
		//result.next = node1
		// 直接返回node1即可，不用再用一个变量完成，
		// 当node1中的第一个比node2小的时候，比较node1.next和node2比较看谁小，然后返回这个节点
		node1.next = mergeTwoLists(node1.next, node2)
		return node1
	}else{
		//result.next = node2
		node2.next = mergeTwoLists(node2.next, node1)
		return node2
	}
	// 这里最终返回的也是node1或者node2，直接在上面返回，必用在增加一个变量了。
	//return result.next
}

// 使用迭代法完成
func mergeTwoLists2(node1, node2 *node)*node  {
	result := &node{}
	current := result
	// 持续循环找小值，增加链表长度。
	for node1 != nil && node2 != nil{
		if node1.value <= node2.value{
			current.next = node1
			node1 = node1.next
		}else{
			current.next = node2
			node2 = node2.next
		}
		current = current.next
	}
	// 判断谁没有了，就把对方的后续节点放到链表上
	if node1 == nil{
		current.next = node2
	}
	if node2 == nil{
		current.next = node1
	}
	return result.next
}

func printLinkList(node *node)  {
	for node != nil{
		fmt.Print(node.value)
		fmt.Print(" ")
		node = node.next
	}
	fmt.Println()
}

func TestMergeTwoLikn(t *testing.T) {
	l1 := &node{value: 1, next: &node{value: 2,next:&node{4, nil}}}
	l2 := &node{value: 1, next: &node{value: 3,next:&node{4, nil}}}
	result := mergeTwoLists(l1, l2)
	printLinkList(result)
	//l11 :=
	// l1 = [], l2 = [0]  l1表示的是一个nil
	l22 := &node{value: 0, next:nil}
	result22 := mergeTwoLists(nil, l22)
	printLinkList(result22)

}

func TestMergeTwoLikn2(t *testing.T) {
	l1 := &node{value: 1, next: &node{value: 2,next:&node{4, nil}}}
	l2 := &node{value: 1, next: &node{value: 3,next:&node{4, nil}}}

	result2 := mergeTwoLists2(l1, l2)
	printLinkList(result2)

	l22 := &node{value: 0, next:nil}
	result22 := mergeTwoLists(nil, l22)
	printLinkList(result22)
}