package host100

import (
	"fmt"
	"testing"
)

// leetcode 160 相交链表
/*
判断两个俩表是否相交。

给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null 。

图示两个链表在节点 c1 开始相交：



题目数据 保证 整个链式结构中不存在环。

注意，函数返回结果后，链表必须 保持其原始结构 。

自定义评测：

评测系统 的输入如下（你设计的程序 不适用 此输入）：

intersectVal - 相交的起始节点的值。如果不存在相交节点，这一值为 0
listA - 第一个链表
listB - 第二个链表
skipA - 在 listA 中（从头节点开始）跳到交叉节点的节点数
skipB - 在 listB 中（从头节点开始）跳到交叉节点的节点数
评测系统将根据这些输入创建链式数据结构，并将两个头节点 headA 和 headB 传递给你的程序。如果程序能够正确返回相交节点，那么你的解决方案将被 视作正确答案 。



示例 1：



输入：intersectVal = 8, listA = [4,1,8,4,5], listB = [5,6,1,8,4,5], skipA = 2, skipB = 3
输出：Intersected at '8'
解释：相交节点的值为 8 （注意，如果两个链表相交则不能为 0）。
从各自的表头开始算起，链表 A 为 [4,1,8,4,5]，链表 B 为 [5,6,1,8,4,5]。
在 A 中，相交节点前有 2 个节点；在 B 中，相交节点前有 3 个节点。
— 请注意相交节点的值不为 1，因为在链表 A 和链表 B 之中值为 1 的节点 (A 中第二个节点和 B 中第三个节点) 是不同的节点。换句话说，它们在内存中指向两个不同的位置，而链表 A 和链表 B 中值为 8 的节点 (A 中第三个节点，B 中第四个节点) 在内存中指向相同的位置。


示例 2：



输入：intersectVal = 2, listA = [1,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1
输出：Intersected at '2'
解释：相交节点的值为 2 （注意，如果两个链表相交则不能为 0）。
从各自的表头开始算起，链表 A 为 [1,9,1,2,4]，链表 B 为 [3,2,4]。
在 A 中，相交节点前有 3 个节点；在 B 中，相交节点前有 1 个节点。
示例 3：



输入：intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2
输出：null
解释：从各自的表头开始算起，链表 A 为 [2,6,4]，链表 B 为 [1,5]。
由于这两个链表不相交，所以 intersectVal 必须为 0，而 skipA 和 skipB 可以是任意值。
这两个链表不相交，因此返回 null 。

*/

/*

思路：
开两个指针分别遍历这两个链表，在第一次遍历到尾部的时候，指向另一个链表头部继续遍历，这样会抵消长度差。

如果链表有相交，那么会在中途相等，返回相交节点；

如果链表不相交，那么最后会 nil == nil，返回 nil

作者：ElliotXX
链接：https://leetcode.cn/problems/intersection-of-two-linked-lists/solutions/5770/60msgo-shi-xian-by-elliotxx/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

在这个特定的例子中，链表 A 和 B 的相交点是从值为 8 的节点开始的，而不是值为 1 的节点。这是因为题目中的“相交”意味着从某个节点开始，两个链表的节点完全相同，并且共享相同的后续节点。

链表 A: [4,1,8,4,5] 链表 B: [5,6,1,8,4,5]

在这个例子中，链表 A 和 B 都有一个节点值为 8，并且这个节点后面跟着 [4,5]。这意味着从值为 8 的节点开始，两个链表是相同的。因此，值为 8 的节点是相交的起始节点。

值为 1 的节点在两个链表中都存在，但它们不是同一个节点。在链表 A 中，值为 1 的节点后面跟着 [8,4,5]，而在链表 B 中，值为 1 的节点后面跟着 [8,4,5]。只有当两个链表的节点完全相同，并且它们是共享的（即同一个节点实例）时，我们才说链表从那个点开始相交。

因此，正确的相交节点是值为 8 的节点，因为这是两个链表中第一个完全相同且共享相同后续节点的节点。

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}

func createList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	for i := 1; i < len(nums); i++ {
		current.Next = &ListNode{Val: nums[i]}
		current = current.Next
	}
	return head
}

func printList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Printf("%d -> ", current.Val)
		current = current.Next
	}
	fmt.Println("nil")
}

// test

func TestGetIntersectionNode(t *testing.T) {
	listA := createList([]int{4, 1, 8, 4, 5})
	listB := createList([]int{5, 6, 1, 8, 4, 5})
	printList(listA)
	printList(listB)
	t.Log(getIntersectionNode(listA, listB))
}

// leetcode 206 反转链表
/*
给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。

输入：head = [1,2,3,4,5]
输出：[5,4,3,2,1]

输入：head = [1,2]
输出：[2,1]
示例 3：

输入：head = []
输出：[]
*/

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var preNode *ListNode
	current := head
	for current != nil {
		tmpNode := current.Next
		current.Next = preNode
		preNode = current
		current = tmpNode
	}

	return preNode
}

// 使用递归方式实现
func reverseListRecursiveDiGui(head *ListNode) *ListNode {
	if head == nil && head.Next == nil {
		return nil
	}
	p := reverseListRecursiveDiGui(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}

// test
func TestReverseList(t *testing.T) {
	list := createList([]int{1, 2, 3, 4, 5})
	printList(list)
	printList(reverseList(list))
}

// leetcode 234 回文链表
func isPalindrome(head *ListNode) bool {
	return true
}
