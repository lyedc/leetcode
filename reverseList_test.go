package leetcode

import (
	"fmt"
	"testing"
)

/*
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

你不能改变节点中的值，只能交换节点本身。

示例：

给定 1->2->3->4，你应该返回 2->1->4->3

func reverseList(head *ListNode) *ListNode {}
*/

type ListNode struct {
	value int
	next  *ListNode
}

// reverseList reverses a linked list and returns the new head.
func SwapNode(head *ListNode) *ListNode {
	result := &ListNode{next: head}
	current := result // nil-1-2-3-4-nil
	for head != nil && head.next != nil {
		first := head // first:= 1-2-3-4-nil     // first := 3-4-nil
		second := head.next // second:= 2-3-4-nil  // sencond:= 4-nil

		// 交换
		first.next = second.next // first:= 1-3-4   // first := 3-nil
		second.next = first // sencond:= 2 -1-3-4   // second := 4-3-nil

		current.next = second  // current:= nil - 2-1-3-4  == result = nil- 2-1-3-4  // current: 1-4-3-nil  ==result: 1-4-3

		// 把current移动到一下对
		current = current.next.next // current:= 1-3-4-nil   //current:= 3-nil
		// 更新 heade 4
		head = first.next  //head:= 3-4-nil  // head: nil
	}

	return result.next
}

func swapPrintLink(node *ListNode) {
	for node != nil {
		fmt.Println(node.value)
		node = node.next
		fmt.Print("")
	}
}

func TestSwapNode(t *testing.T) {
	l1 := &ListNode{value: 1, next: &ListNode{value: 2, next: &ListNode{value: 3, next: &ListNode{value: 4, next: nil}}}}
	swapPrintLink(l1)
	result := SwapNode(l1)
	swapPrintLink(result)
}
