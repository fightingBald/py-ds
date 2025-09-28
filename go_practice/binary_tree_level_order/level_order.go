package levelorder

// Problem: Given the root of a binary tree, return the level order traversal of its nodes' values from left to right, level by level.
// This is also known as breadth-first search (BFS) traversal.
//
// Example:
//   Input: root = [3,9,20,null,null,15,7]
//   Output: [[3],[9,20],[15,7]]
//
// Implement the levelOrder function below so it returns the correct traversal without modifying the provided tests.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LevelOrder(root *TreeNode) [][]int {
	panic("TODO: implement LevelOrder")
}
