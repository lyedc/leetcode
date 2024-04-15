package host100

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode 94 二叉树中序遍历

/*
执行的结果就是看根节点的位置是什么时候执行的，中序就是中间执行。
输入：root = [1,null,2,3]
输出：[1,3,2]
示例 2：

输入：root = []
输出：[]
示例 3：

输入：root = [1]
输出：[1]
*/

func InorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	res = append(res, InorderTraversal(root.Left)...)
	res = append(res, root.Val)
	res = append(res, InorderTraversal(root.Right)...)
	return res
}

// leetcode 104 二叉树的最大深度

// 这个也是深度优先算法。。。
func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDeep := MaxDepth(root.Left) // 所以每次递归后的结果都是1
	rightDeep := MaxDepth(root.Right)
	return 1 + max2(leftDeep, rightDeep)
}

func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// leetcode 226 翻转二叉树

func InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 先把最上层的左右交换，递归每个叶子节点的左右交换
	root.Left, root.Right = root.Right, root.Left
	InvertTree(root.Left)
	InvertTree(root.Right)
	return root
}

// leetcode 101 对称二叉树
func IsSymmetric(root *TreeNode) bool {
	// 用go实现
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

func isMirror(left, right *TreeNode) bool {
	// 递归
	if left == nil && right == nil {
		return true
	}
	// 左边等于nil，但是右边不等于nil
	if left == nil || right == nil {
		return false
	}
	// 左右两边的值不想等
	if left.Val != right.Val {
		return false
	}
	// 递归。。。
	return isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
}

// leetcode 102 二叉树的层次遍历
func LevelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}
	var levelOrder func(node *TreeNode, level int)
	levelOrder = func(node *TreeNode, level int){
		if node == nil{
			return
		}
		// 如果是一个新的层级,就创建一个数组...
		if len(result) == level{
			result = append(result, []int{})
		}
		// 一个层级一个数组,并把层级的值加入到数组中.
		result[level] = append(result[level], node.Val)
		levelOrder(node.Left, level +1)
		levelOrder(node.Right,level +1)
	}
	levelOrder(root, 0)
	return result
}

// levelOrder 使用队列进行层次遍历
func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var result [][]int
	var queue []*TreeNode
	queue = append(queue, root)

	for len(queue) > 0 {
		levelSize := len(queue)
		level := make([]int, levelSize)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			// 永远取出队列最上面的叶子节点.
			queue = queue[1:]
			level[i] = node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, level)
	}
	return result
}

// leetcode 108 将有序数组转换为二叉搜索树
/*
要将一个有序数组转换为二叉搜索树，我们可以使用以下步骤：

选择数组中间的元素作为根节点。
将数组分为两部分，左边的元素都小于根节点的值，右边的元素都大于根节点的值。
递归地创建左子树和右子树。
*/
func SortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0{
		return nil
	}
	mid := len(nums) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = SortedArrayToBST(nums[:mid])
	root.Right = SortedArrayToBST(nums[mid+1:])
	return root
}

// leetcode 98 验证二叉搜索树
func isValidBST(root *TreeNode) bool {

}