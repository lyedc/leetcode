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
	if head == nil || head.Next == nil {
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
/*
1. 先使用快慢指针找到链表的中间节点
2. 反转后半部分的链表
3. 比较前半部分和后半部分的链表
4. 恢复链表

*/
func isPalindrome(head *ListNode) bool {
	if head == nil && head.Next == nil {
		return true
	}
	// 使用快慢指针找到链表的中间节点
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		// 慢指针移动一步
		// 快指针移动两步
		// 如果是奇数个节点，快指针会到达最后一个节点
		// 如果是偶数个节点，快指针会到达倒数第二个节点
		// 慢指针和快指针相等时，慢指针到达中间节点
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 反转后半部分链表
	reversedSecondHalf := reverseList2(slow.Next)
	//循环遍历后半部分的链表,和前面进行对比
	for reversedSecondHalf != nil {
		if head.Val != reversedSecondHalf.Val {
			return false
		}
		reversedSecondHalf = reversedSecondHalf.Next
		head = head.Next
	}
	// 还原链表
	//slow.Next = reverseList2(reversedSecondHalf)
	return true
}

func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := reverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}

// test
func TestIsPalindrome(t *testing.T) {
	list := createList([]int{1, 2, 3, 4, 5})
	printList(list)
	t.Log(isPalindrome(list))
	printList(list)
	list2 := createList([]int{1, 2, 2, 1})
	printList(list2)
	t.Log(isPalindrome(list2))
	printList(list2)
}

// leetcode 141 环形链表

func hasCycle(head *ListNode) bool {
	// 表明没有节点
	if head == nil || head.Next == nil {
		return false
	}
	// 设置快慢指针
	slow, fast := head, head.Next
	// 如果两个指针最终没有相遇的话，如果快指针先链表尾部的话，表示没有环形
	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

func hasCycle2(head *ListNode) bool {
	// 表明没有节点
	if head == nil || head.Next == nil {
		return false
	}
	// 设置快慢指针
	slow, fast := head, head.Next
	// 如果两个指针最终没有相遇的话，如果快指针先链表尾部的话，表示没有环形
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}
	return false
}

// test
func TestHasCycle(t *testing.T) {
	list := createList([]int{1, 2, 3, 4, 5})
	printList(list)
	t.Log(hasCycle(list))
}

// leetcode 142 环形链表 II 返回环的位置

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow, fast := head, head
	loopExists := false
	// 判断是否存在环形
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			loopExists = true
			break
		}
	}
	// 不能存在环形就直接返回
	if !loopExists {
		return nil
	}

	// 快慢指针相遇后，快指针回到链表头部
	fast = head
	// 快慢指针以相同速度移动，再次相遇时即为环的入口
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return slow
}

// test
func TestDetectCycle(t *testing.T) {
}

// leetcode 21 合并两个有序链表
/*
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]
示例 2：

输入：l1 = [], l2 = []
输出：[]
示例 3：

输入：l1 = [], l2 = [0]
输出：[0]
*/
/*
递归： 时间复杂度： o(m +n)
空间复杂度：O（m + n）
*/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 采用递归的方式
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}

// 迭代法
/*
时间复杂度为：n(m +n)
空间复杂度为：O(1)
*/
func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	tmp := &ListNode{}
	// 赋值之后，cur是一直在移动的，这样tmp就是这个链表的整体了。
	cur := tmp
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		// 移动cur
		cur = cur.Next
	}
	if list1 == nil {
		cur.Next = list2
	}
	if list2 == nil {
		cur.Next = list1
	}
	return tmp.Next

}

// test
func TestMergeTwoLists(t *testing.T) {
	// 测试用例
	list1 := createList([]int{1, 2, 4})
	list2 := createList([]int{1, 3, 4})
	printList(mergeTwoLists(list1, list2))

}

// leetcode 2 两数相加
/*
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0 开头

*/
/*

时间复杂度：O(max(m,n))
空间复杂度：O(1）
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	tmp := result
	next := 0
	// 这里需要计算next ！=0 的情况。，这个next的判断是为了最后一位的进位计算的。，防止丢失，如果等于0的话
	// 表示链表刚好等于原来的长度，如果不等于0，链表就要增加一个节点。
	for l1 != nil || l2 != nil || next != 0 {
		total := next
		if l1 != nil {
			total += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			total += l2.Val
			l2 = l2.Next
		}
		tmp.Next = &ListNode{Val: total % 10}
		tmp = tmp.Next
		next = total / 10
	}
	return result.Next
}

func ddd(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	tmp := result
	next := 0
	for l1 != nil || l2 != nil || next != 0 {
		total := next
		if l1 != nil {
			total += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			total += l2.Val
			l2 = l2.Next
		}
		tmp.Next = &ListNode{Val: total % 10}
		tmp = tmp.Next
		next = total / 10
	}
	return result.Next
}

// leetcode 19 删除链表的倒数第 N 个结点
/*
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	result := &ListNode{Next: head}
	fast, slow := result, result
	// 计算长度。得出前N个数据的位置
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	// fast 移动到尾部的时候，slow刚好是N节点的前一个数据，因为多了个头部需要fast多走一步，slow少走一步
	// 这样刚好计算出倒数第N个节点的前一个节点
	// 总长是M， fast继续走了 m-n的距离。slow，也走了m-n的距离。这个时候，slow就到了倒数第N个节点的前一个节点。
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return result.Next
}

// test

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	//result := &ListNode{Next: head}
	// 可以不用增加一个虚拟节点.直接操作heade节点即可.
	fast, slow := head, head
	// 计算长度。得出前N个数据的位置
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	// fast 移动到尾部的时候，slow刚好是N节点的前一个数据，因为多了个头部需要fast多走一步，slow少走一步
	// 这样刚好计算出倒数第N个节点的前一个节点
	// 总长是M， fast继续走了 m-n的距离。slow，也走了m-n的距离。这个时候，slow就到了倒数第N个节点的前一个节点。
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return head
}

func TestRemoveNthFromEnd(t *testing.T) {
	list := createList([]int{1, 2, 3, 4, 5})
	printList(removeNthFromEnd2(list, 2))
}

// leetcode 24 两两交换链表中的节点

/*
给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。
你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）

输入：head = [1,2,3,4]
输出：[2,1,4,3]
示例 2：

输入：head = []
输出：[]
示例 3：

输入：head = [1]
输出：[1]

// 一对一对的进行处理，先处理第一对，然后再处理下一对，直到处理完毕。

*/

func swapPairs(head *ListNode) *ListNode {
	result := &ListNode{Next: head}
	cur := result
	for head != nil && head.Next != nil {
		// 保存下一个节点
		next := head.Next
		cur.Next = next
		// 交换当前节点到下一个节点
		// 对前两个节点进行交换
		head.Next = next.Next
		next.Next = head
		// 相当于是: 0-1-2-3-4
		// 0-1-2 变成： 0-2-1-3-4
		// 1-3-4： 0-2-1-4-3
		// 移动到下一对节点
		cur = head
		head = head.Next
	}
	return result.Next
}

// 采用递归的形式完成。
func swapPairs123(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	head.Next = swapPairs(newHead.Next)
	newHead.Next = head
	return newHead
}

// leetcode 138 复制带随机指针的链表

// leetcode 148 排序链表
/*

采用归并排序的方式，把链表从中间分开，分别排序，然后递归。。。
最后采用合并两个有序链表的方式合并链表。


*/

func sortListNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 找到中间节点， 一个走一步，一个走两步这样slow就是中间节点。
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// slow 就是中间节点
	// 排序中间节点的右边
	right := sortListNode(slow.Next)
	// 断开中间节点
	slow.Next = nil
	// 排序中间节点的左边
	left := sortListNode(head)
	// 合并两个有序链表
	return mergeTwoLists2(left, right)
}

// test

func TestSortListNode(t *testing.T) {
	list := createList([]int{4, 2, 1, 3})
	printList(sortListNode(list))
}

// leetcode 146 LRU 缓存机制

/*
请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
实现 LRUCache 类：
LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
*/

type DlinkedMode struct {
	Key   int
	Value int
	Prev  *DlinkedMode
	Next  *DlinkedMode
}

type LRUCache struct {
	// 容量
	capacity int
	// 缓存数据
	hashMap map[int]*DlinkedMode
	// 缓存数据个数
	head *DlinkedMode
	tail *DlinkedMode
}

func Constructor(capacity int) LRUCache {
	head := &DlinkedMode{}
	tail := &DlinkedMode{}
	head.Next = tail
	tail.Prev = head
	return LRUCache{
		capacity: capacity,
		hashMap:  make(map[int]*DlinkedMode),
		head:     head,
		tail:     tail,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.hashMap[key]; ok {
		this.moveToHead(node)
		return node.Value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	// 存在这个key，就更新之，并放到最前面去
	if node, ok := this.hashMap[key]; ok {
		node.Value = value
		this.moveToHead(node)
		return
	}
	// 没有的话，就创建一个新的节点，并放入到最前面去
	newNode := &DlinkedMode{Key: key, Value: value}
	this.hashMap[key] = newNode
	this.addToHead(newNode)
	// 判读是否满了。
	if len(this.hashMap) > this.capacity {
		// 删除后面的数据
		removed := this.removeTail()
		delete(this.hashMap, removed.Key)
	}

}

func (this *LRUCache) addToHead(node *DlinkedMode) {
	node.Prev = this.head
	node.Next = this.head.Next
	this.head.Next.Prev = node
	this.head.Next = node
}

func (this *LRUCache) removeNode(node *DlinkedMode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (this *LRUCache) moveToHead(node *DlinkedMode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *DlinkedMode {
	node := this.tail.Prev
	this.removeNode(node)
	// 这里需要把这个node返回回去，如果不返回回去，后面删除的时候，会报错。
	return node
}

// func copyRandomList(head *Node) *Node {}
/*

链表总结：
当时找数据的时候：需要有一个空的头部节点
*/
