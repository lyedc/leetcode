package leetcode

import (
	"fmt"
	"testing"
)

/*
给定两个非空的链表，用来表示两个非负的整数。其中，它们各自的位数是按照逆序的方式存储的，并且它们的每一个节点只能存储一位数字。

请你将这两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数字都不会以零开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

*/

// 采用遍历的方式

type Node struct {
	value int
	next *Node
}

func newNode(value int)*Node  {
	return &Node{
		value: value,
		next:  nil,
	}
}

func twoNumber(node1, node2 *Node)*Node  {
	total := 0
	next := 0
	result := &Node{}
	current := result
	for{
		if node1 == nil || node2 == nil{
			break
		}
		total = node1.value + node2.value + next
		next = total/10   // 取
		current.next = newNode(total%10) // 这里10 的话的到的是余数,也就是想加的想加的第一个数,
		/*
		12 %10 = 2 个位
		12/10 =1 十位
		5 % 10 = 5
		5 / 10 =0
		*/
		node1 = node1.next
		node2 = node2.next
		current = current.next
	}
	// 如果node2遍历完了,那就接着遍历node1剩余的,和next想加
	for{
		if node1 == nil{
			break
		}
		total = node1.value + next
		current.next = newNode(total%10)
		next = total/10
		node1 = node1.next
		current = current.next
	}
	// 如果node1遍历完了,继续遍历node2剩余的和next的值.
	for{
		if node2 == nil{
			break
		}
		total = node2.value + next
		current.next = newNode(total%10)
		next = total/10
		node2 = node2.next
		current = current.next
	}
	
	return result.next
}


// 改进版
func addTwoNumbers(node1, node2 *Node)*Node  {
	result := &Node{}
	current := result
	next := 0
	for node1 == nil || node2 == nil{
		total := next
		// 这里的思路是逐个的去node上的值加到总和上,因为总和每次循环的时候都进行了初始化,都等于node.value + node2.value 的于数
		// 所以相对于第一种是在开头就直接加上了上一次的于数,
		if node1 != nil{
			total += node1.value
			node1 = node1.next
		}
		if node2 != nil{
			total += node2.value
			node2 = node2.next
		}
		current.next = newNode(total%10)
		next = total/10
		current = current.next
	}
	return result.next
}

// 使用递归的方式进行,没有成功...
func diguiTwoSum(node1, node2 *Node)*Node  {
	total := 0
	result := &Node{}
	current := result
	next := 0
	if node1.value ==0 && node2.value == 0{
		return result.next
	}
	total = node1.value + node2.value + next
	next = total/10
	current.next = newNode(total%10)
	current = current.next
	node1 = node1.next
	node2 = node2.next
	if node1 == nil{
		node1 = newNode(0)
	}
	if node2 == nil{
		node2 = newNode(0)
	}
	return twoNumber(node1, node2)
}


func printList(head *Node) {
	for head != nil {
		fmt.Printf("%d ", head.value)
		head = head.next
	}
	fmt.Println()
}

func Test_twoNumber(t *testing.T) {
	l1 := &Node{2, &Node{4, &Node{3, nil}}}
	l2 := &Node{5, &Node{6, &Node{4, nil}}}
	result := twoNumber(l1, l2)
	printList(result)
	 result2 := twoNumber(l1, l2)
	 printList(result2)
	 result3 := diguiTwoSum(l1, l2)
	 printList(result3)
}