package code

type listNode struct {
	val  int
	next *listNode
}

func detectCycle(head *listNode) *listNode {
	if head == nil || head.next == nil {
		return nil
	}
	slow, fast := head, head
	flag := true
	for slow != fast {
		slow = slow.next
		fast = fast.next.next
		if fast == nil || fast.next == nil {
			flag = false
		}
	}
	if !flag {
		return nil
	}
	fast = head
	for fast != slow {
		fast = fast.next
		slow = slow.next
	}
	return slow
}
