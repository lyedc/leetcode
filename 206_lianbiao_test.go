package leetcode

import (
	"fmt"
	"testing"
)

// 206. 反转链表
/*
Reverse a singly linked list.

Example:

Input: 1->2->3->4->5->NULL Output: 5->4->3->2->1->NULL

*/


type listNode2 struct {
	val  int
	next *listNode2
}

func ReverseLinkeList(head *listNode2)*listNode2  {
	var pre *listNode2
	current := head
	for current != nil{
		nextTemp := current.next // 保存当前节点的下一个节点
		current.next = pre      // 反转当前节点的指针
		pre = current           // 更新prev为当前节点
		current = nextTemp       // 移动到下一个节点
	}
	return pre
}

// test case
func Test_ReverseLinkeList(t *testing.T){
	head := &listNode2{1,nil}
	head.next = &listNode2{2,nil}
	head.next.next = &listNode2{3,nil}
	head.next.next.next = &listNode2{4,nil}
	head.next.next.next.next = &listNode2{5,nil}
	printListNode2(head)
	result := ReverseLinkeList(head)
	fmt.Println("==============")
	printListNode2(result)
}

func printListNode2(node *listNode2)  {
	for node != nil{
		fmt.Print(node.val)
		node = node.next
	}
}