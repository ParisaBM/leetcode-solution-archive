package main
import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func max_path_sum_helper(root *TreeNode) (int, int) {
	left_max, left_max_to_root := -1001, 0
	if root.Left != nil {
		left_max, left_max_to_root = max_path_sum_helper(root.Left)
	}
	right_max, right_max_to_root := -1001, 0
	if root.Right != nil {
		right_max, right_max_to_root = max_path_sum_helper(root.Right)
	}
	max_to_root := max(0, root.Val + max(left_max_to_root, right_max_to_root))
	max_overall := max(max(left_max, right_max), root.Val + left_max_to_root + right_max_to_root)
	return max_overall, max_to_root
}

func max_path_sum(root *TreeNode) int {
	result, _ := max_path_sum_helper(root)
	return result
}

func main() {
	fmt.Println(max_path_sum(&TreeNode{1, &TreeNode{Val: 2}, &TreeNode{Val: 3}}))
	fmt.Println(max_path_sum(&TreeNode{-10, &TreeNode{Val: 9}, &TreeNode{20, &TreeNode{Val: 15}, &TreeNode{Val: 7}}}))
}
