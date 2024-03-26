package leetcode

import (
	"fmt"
	"testing"
)

// 203. 移除链表元素

/*
Remove all elements from a linked list of integers that have value val.

Example:

Input: 1->2->6->3->4->5->6, val = 6 Output: 1->2->3->4->5
*/

type listNode struct {
	val int
	next *listNode
}

func removeEle(head *listNode, target int)*listNode{
	// 做一个虚拟节点,便于处理边界问题,也就是头节点如果是target的话,便于处理
	dumpHead := &listNode{next: head}
	current := dumpHead
	for current.next != nil{
		// 判断头部节点是否需要删除.
		tmp := current.next
		if tmp.val == target{
			// 头部的下一个节点指向下一个节点.
			current.next = tmp.next
		}else {
			// current 指针移动到下一个节点. 下一个节点就是tmp
			current = tmp

		}
	}
	// 返回头节点
	return dumpHead.next
}

// test
func Test_removeEle(t *testing.T) {

	head := &listNode{
		val: 1,
		next: &listNode{
			val: 2,
			next: &listNode{
				val: 6,
				next: &listNode{
					val: 3,
					next: &listNode{
						val: 4,
						next: &listNode{
							val: 5,
							next: &listNode{
								val: 6,
								next: nil,
							},
						},
					},
				},
			},
		},
	}
	printListNode(head)
	result := removeEle(head, 6)
	fmt.Println("============")
	printListNode(result)
}

func printListNode(node *listNode)  {
	for node!= nil{
		fmt.Print(node.val)
		node = node.next
	}
}