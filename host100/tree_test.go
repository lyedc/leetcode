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
	levelOrder = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		// 如果是一个新的层级,就创建一个数组...
		if len(result) == level {
			result = append(result, []int{})
		}
		// 一个层级一个数组,并把层级的值加入到数组中.
		result[level] = append(result[level], node.Val)
		levelOrder(node.Left, level+1)
		levelOrder(node.Right, level+1)
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

// TreeNode 定义二叉搜索树的节点结构

// BST 定义二叉搜索树结构，根节点为私有成员
type BST struct {
	root *TreeNode
}

// NewBST 创建一个新的二叉搜索树
func NewBST() *BST {
	return &BST{}
}

// Insert 向二叉搜索树中插入一个新值
func (bst *BST) Insert(val int) {
	bst.root = insert(bst.root, val)
}

// insert 是 Insert 方法的辅助函数，用于递归地插入节点
func insert(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return &TreeNode{Val: val}
	}
	// 小于根节点，插入到左边，并跟node的左边进行比较
	if val < node.Val {
		node.Left = insert(node.Left, val)
		// 如果大于根节点，插入到右边，并跟node的右边进行比较
	} else if val > node.Val {
		node.Right = insert(node.Right, val)
	}
	// 如果 val == node.Val，什么也不做（BST中不允许有重复元素）
	return node
}

/*

二叉搜索树（BST）是一种特殊的二叉树，其中每个节点的值都大于或等于其左子树上所有节点的值，并且小于或等于其右子树上所有节点的值。以下是二叉搜索树的一个例子：

        8
       / \
      3   10
     / \    \
    1   6    14
       /  \  /
      4   7  13

*/

// leetcode 108 将有序数组转换为二叉搜索树
// 搜索二叉树的定义: 左子树的所有节点的值都小于根节点的值，右子树的所有节点的值都大于根节点的值。
/*
要将一个有序数组转换为二叉搜索树，我们可以使用以下步骤：

选择数组中间的元素作为根节点。
将数组分为两部分，左边的元素都小于根节点的值，右边的元素都大于根节点的值。
递归地创建左子树和右子树。
*/
func SortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = SortedArrayToBST(nums[:mid])
	root.Right = SortedArrayToBST(nums[mid+1:])
	return root
}

// leetcode 98 验证搜索二叉树
func IsValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return helper(root, nil, nil)
}

func helper(root *TreeNode, min, max *int) bool {
	if root == nil {
		return true
	}
	if min != nil && root.Val <= *min {
		return false
	}
	if max != nil && root.Val >= *max {
		return false
	}
	// 这里不能直接的传入nil， 一定要根节点的，根节点传入到递归函数中，所以这里应该是min和max
	// 左边是最小值，左边的值都要小于root节点
	if !helper(root.Left, min, &root.Val) {
		return false
	}
	// 右边的都要大于root节点
	if !helper(root.Right, &root.Val, max) {
		return false
	}
	return true
}

func isValidBST2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isValidBSTHelper(root, nil, nil)
}

func isValidBSTHelper(node *TreeNode, lower, upper *int) bool {
	if node == nil {
		return true
	}
	if lower != nil && node.Val <= *lower {
		return false
	}
	if upper != nil && node.Val >= *upper {
		return false
	}
	if !isValidBSTHelper(node.Left, lower, &node.Val) {
		return false
	}
	if !isValidBSTHelper(node.Right, &node.Val, upper) {
		return false
	}
	return true
}

// leetcode 230 二叉搜索树中第K小的元素
/*
二叉搜索树，左边的都小于根节点，右边都大于根节点
*/
func kthSmallest(root *TreeNode, k int) int {
	var res int
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		k--
		if k == 0 {
			res = node.Val
			return
		}
		inorder(node.Right)
	}
	inorder(root)
	return res
}

// 使用栈的方式完成
/*
先遍历全部的左边，然后再遍历右边。

*/
func kthSmallest2(root *TreeNode, k int) int {
	stack := []*TreeNode{}
	current := root
	for current != nil || len(stack) > 0 {
		// 把左子树全部入栈，这样栈顶的元素就是最小的元素
		if current != nil {
			stack = append(stack, current)
			current = current.Left
		} else {
			// 如果左边都入栈了，那么久开始比较大小了。弹出栈顶的数据就是树的左子树的小值
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			k--
			// 达到最小的数据了，就表示找到了。因为是从tree的底部开始遍历的。
			if k == 0 {
				return root.Val
			}
			// 判断最后的左子树的右边的值，走到这里表示最小的最子树不满足，需要判断右边的数据。
			// 到达最后的时候root=nil，但是stack不等于空，所以继续循环。这个时候就会弹出上一个节点，然后root就不是ni了。并且root.right不为nil
			current = current.Right
		}
	}
	// 表示没有最小的元素
	return 0
}

// leetcode  199 二叉树的右视图
// 使用广度优先算法 bfs进行层次遍历,保留最后一个节点即可
/*
在这个Go语言实现中，我们同样使用了一个切片queue来模拟队列，
用于存储每一层的节点。在每一层的开始，
我们先计算出这一层的节点数量levelLength，
然后遍历这些节点，将它们的子节点加入队列中。
对于每一层，我们只保留最后一个节点的值，这就是右视图中的一个节点
*/
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	// 使用队列进行bfs广度优先遍历
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		// 第一次表示是第一层,第一层就一个节点.遍历完第一层后,就会把第一层的叶子节点加入到队列中.
		leveLength := len(queue)
		// 每次的一个for循环就是在循环一个层级。
		for i := 0; i < leveLength; i++ {
			// 队列是先进先出,弹出队列最前面的节点.第一次弹出的左边的节点,
			//最后一个弹出的是右边的节点.
			// 比如第二个节点,第一次弹出的是左边的节点,第二次弹出的是右边的节点.
			node := queue[0]
			queue = queue[1:]
			// 如果是当前层的最后一个节点,就加入到结果中。因为右视图看到的是tree的最右侧的一个分支。
			if i == leveLength-1 {
				result = append(result, node.Val)
			}
			// 加入第一层的叶子节点,
			if node.Left != nil {
				queue = append(queue, node)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

		}

	}
	return result
}

// leetcode 114 二叉树展开为链表
/*
先前序遍历排好序
然后把节点连接起来,原有的tree的left都为nil,右边的逐个链接起来.
*/
func flatten(root *TreeNode) {
	list := treeqianxu(root)
	for i := 1; i < len(list); i++ {
		pre, cur := list[i-1], list[i]
		pre.Left, pre.Right = nil, cur
	}
}

func treeqianxu(root *TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}
	var result []*TreeNode
	result = append(result, root)
	result = append(result, treeqianxu(root.Left)...)
	result = append(result, treeqianxu(root.Right)...)
	return result
}

// 通过前驱节点逐步的链接,画图比较容易理解.

func flatten2(root *TreeNode) {
	curr := root
	for curr != nil {
		if curr.Left != nil {
			next := curr.Left
			predecessor := next
			// 找到左子树最右边的节点找到最右一个.
			for predecessor.Right != nil {
				predecessor = predecessor.Right
			}
			// 把左边的最右节点的right的值,替换为 root的right.
			predecessor.Right = curr.Right
			curr.Left, curr.Right = nil, next
		}
		// 逐步的获取每个节点的右节点作为当前节点.继续下一个节点的判断.
		curr = curr.Right
	}
}
