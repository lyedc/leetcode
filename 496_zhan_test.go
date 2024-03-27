package leetcode

import (
	"testing"
)

type stack struct {
	data []int
}

func NewStack1() *stack {
	return &stack{data: make([]int, 0)}
}

func (s *stack) getTop() int {
	return s.data[len(s.data)-1]
}

func (s *stack) push(item int) {
	s.data = append(s.data, item)
}

func (s *stack) pop() int {
	if s.isEmpty() {
		return -1
	}
	item := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return item
}

func (s *stack) isEmpty() bool {
	return len(s.data) == 0
}

// 496 下一个更大元素 I
// 这样做是限定为只能找挨着的元素，如果没有这个限定条件的话，这样实现就不行了。
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return nil
	}
	tmpMap := make(map[int]int)
	for i := 0; i < len(nums2); i++ {
		// 因为没有重复的元素，直接放入到map就即可，
		// 遍历到最后一个元素的时候，因为最右边没有元素了，直接复制-1即可。
		if i < len(nums2)-1 {
			tmpMap[nums2[i]] = nums2[i+1]
		} else {
			tmpMap[nums2[i]] = -1
		}
	}
	for i := 0; i < len(nums1); i++ {
		tmp := nums1[i]
		if value := tmpMap[tmp]; value > tmp {
			nums1[i] = value
		} else {
			nums1[i] = -1
		}
	}

	return nums1
}

// 使用一个栈来实现
func nextGreaterElement2(nums1 []int, nums2 []int) []int {
	tmpMap := make(map[int]int)
	// 栈中
	tmpStack := NewStack1()
	for _, item := range nums2 {
		// 判断栈中是否为空，并且栈中的元素是否大于即将入栈的元素，如果是的话
		// 就弹出栈顶元素，并把要入栈的数据放入到map中，
		// 其实跟上面的实现的思路是一样的。他这样的方式是直接对map中的数据进行好了大小的初始化
		if !tmpStack.isEmpty() && tmpStack.getTop() < item {
			// 把栈顶的元素弹出来，放入到map中，并把大于栈顶的元素当作map的value
			tmpMap[tmpStack.pop()] = item
		}
		// 把当前元素压入栈中
		tmpStack.push(item)
	}
	// 栈中剩余的元素，就是在数据中没有比它大的相邻的元素了。
	// 这样就把栈中的元素直接放入map中，key为栈中的元素，value为-1
	for !tmpStack.isEmpty() {
		tmpMap[tmpStack.pop()] = -1
	}
	// 和map中的元素进行比较，因为map中的元素已经对应好数据的大小了。直接拿map的value即可。
	for index, item := range nums1 {
		nums1[index] = tmpMap[item]
	}
	return nums1

}

// test nextGreaterElement
func TestNextGreaterElement(t *testing.T) {
	nums1 := []int{4, 0, 3}
	nums2 := []int{0, 3, 4, 7}
	t.Log(nextGreaterElement(nums1, nums2))
	nums3 := []int{4, 0, 3}
	t.Log(nextGreaterElement2(nums3, nums2))
}
