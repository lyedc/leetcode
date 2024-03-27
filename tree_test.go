package leetcode

import (
	"reflect"
	"testing"
)

/*

没有孩子的节点都是叶子节点

根节点:
叶子节点:
节点:


树的高度: 从下面算起,最下层是0,根节点是1,所以树的高度就是节点数-1

深度: 从上面算起,根节点是0,所以深度就是节点数-1

普通二叉树: 每个节点最多两个孩子
满二叉树: 每个节点最多两个孩子,并且每个节点都有两个孩子,并且每个节点的左孩子和右孩子都是满二叉树, 所有的叶子节点都在同一层...
完全二叉树: 从树的根节点,从上到下,从左到右依次填满节点形成的二叉树. 依次中间不能有空的...

满二叉树一定是完全二叉树,完全二叉树一定是满二叉

遍历的是根的位置:

前序遍历: 根左右 每个节点都先访问,然后是左子树,最后是右子树
中序遍历: 左根右 每个节点都先访问,然后是左子树,最后是右子树
后序遍历: 左右根 每个节点都先访问,然后是左子树,最后是右子树
每个子树都按照一个树的形式去访问: 也就是对树进行分拆...

*/

// 144 二叉树的前序遍历

/*
输入：root = [1,null,2,3]
输出：[1,2,3]
示例 2：

输入：root = []
输出：[]
示例 3：

输入：root = [1]
输出：[1]
*/

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	result = append(result, root.val)
	result = append(result, preorderTraversal(root.left)...)
	result = append(result, preorderTraversal(root.right)...)
	return result
}

func TestPreorderTraversal(t *testing.T) {
	// 测试用例1：空树
	var emptyTree *TreeNode
	expected := []int{}
	actual := preorderTraversal(emptyTree)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}

	// 测试用例2：单节点树
	SingleNode := &TreeNode{val: 1}
	expected = []int{1}
	actual = preorderTraversal(SingleNode)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}

	// 测试用例3：具有多个节点的完整二叉树
	root := &TreeNode{
		val: 1,
		left: &TreeNode{
			val:   2,
			left:  &TreeNode{val: 4},
			right: &TreeNode{val: 5},
		},
		right: &TreeNode{
			val:   3,
			left:  &TreeNode{val: 6},
			right: &TreeNode{val: 7},
		},
	}
	expected = []int{1, 2, 4, 5, 3, 6, 7}
	actual = preorderTraversal(root)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}

}

// 94 二叉树的中序遍历

func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	if root == nil{
		return  result
	}
	result = append(result, inorderTraversal(root.left)...)
	result = append(result, root.val)
	result = append(result, inorderTraversal(root.right)...)
	return result
}

// test
func TestInorderTraversal(t *testing.T) {
	// 测试用例1：空树
	var emptyTree *TreeNode
	expected := []int{}
	actual := inorderTraversal(emptyTree)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}

	// 测试用例2：单节点树
	SingleNode := &TreeNode{val: 1}
	expected = []int{1}
	actual = inorderTraversal(SingleNode)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
	// 测试用例3：具有多个节点的完整二叉树
	root := &TreeNode{
		val: 1,
		left: &TreeNode{
			val:   2,
			left:  &TreeNode{val: 4},
			right: &TreeNode{val: 5},
		},
		right: &TreeNode{
			val:   3,
			left:  &TreeNode{val: 6},
			right: &TreeNode{val: 7},
		},
	}
	expected = []int{4, 2, 5, 1, 6, 3, 7}
	actual = inorderTraversal(root)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

// 145 二叉树的后序遍历

func postorderTraversal (root *TreeNode) []int {
	result := []int{}
	if root == nil{
		return result
	}
	result = append(result, postorderTraversal(root.left)...)
	result = append(result, postorderTraversal(root.right)...)
	result = append(result, root.val)
	return result
}

// test
func TestPostorderTraversal(t *testing.T) {
	// 测试用例1：空树
	var emptyTree *TreeNode
	expected := []int{}
	actual := postorderTraversal(emptyTree)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}

	// 测试用例2：单节点树
	SingleNode := &TreeNode{val: 1}
	expected = []int{1}
	actual = postorderTraversal(SingleNode)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
	// 测试用例3：具有多个节点的完整二叉树
	root := &TreeNode{
		val: 1,
		left: &TreeNode{
			val:   2,
			left:  &TreeNode{val: 4},
			right: &TreeNode{val: 5},
		},
		right: &TreeNode{
			val:   3,
			left:  &TreeNode{val: 6},
			right: &TreeNode{val: 7},
		},
	}
	expected = []int{4, 5, 2, 6, 7, 3, 1}
	actual = postorderTraversal(root)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

