package leetcode

import (
	"fmt"
	"testing"
)

// 环形链表：给定一个链表，判断链表中是否有环。

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow := head
	fast := head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			return true
		}
	}
	return false
}

// test

func TestHasCycle(t *testing.T) {
	head := &ListNode{value: 3}
	head.next = &ListNode{value: 2}
	head.next.next = &ListNode{value: 0}
	head.next.next.next = &ListNode{value: -4}
	head.next.next.next.next = head.next
	res := hasCycle(head)
	fmt.Print(res)
}
