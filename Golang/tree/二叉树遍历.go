package main

import "fmt"

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

// 前序遍历
func Preorder(head *TreeNode) {
	if head == nil {
		return
	}
	// 先访问根节点
	print(head.Val)
	// 再访问左节点
	Preorder(head.Left)
	// 最后访问右节点
	Preorder(head.Right)
}

// 中序遍历
func Inorder(head *TreeNode) {
	if head == nil {
		return
	}
	// 先访问左节点
	Inorder(head.Left)
	// 再访问根节点
	print(head.Val)
	// 最后访问右节点
	Inorder(head.Right)
}

// 后序遍历
func Postorder(head *TreeNode) {
	if head == nil {
		return
	}
	// 先访问左节点
	Postorder(head.Left)
	// 再访问右节点
	Postorder(head.Right)
	// 最后访问根节点
	print(head.Val)
}

// 层序遍历
func LevelOrder(root *TreeNode) (vals []int) {
	vals = append(vals, 1)
	fmt.Println(vals)
	return
}

func main() {
	// 构建二叉树
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 7}
	// 前序遍历
	Preorder(root)
	fmt.Println("\n")
	// 中序遍历
	Inorder(root)
	fmt.Println("\n")
	// 后序遍历
	Postorder(root)

	// 层序遍历
	fmt.Println("层序遍历\n")
	LevelOrder(root)
}
