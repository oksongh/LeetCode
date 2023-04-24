package main

import (
// "fmt"
)

func main() {

}

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type Stack []int

func (stack *Stack) PushBack(i int) {
	*stack = append(*stack, i)

}

func (stack *Stack) PopBack() int {
	n := len(*stack) - 1
	defer func() { *stack = (*stack)[:n] }()
	return (*stack)[n]
}

func isPseudoPalindrom(stack Stack) bool {
	nums := make([]int, 10)
	allowed_odd := 0
	if len(stack)%2 == 1 {
		allowed_odd = 1
	}

	for _, v := range stack {
		nums[v]++
	}
	for _, v := range nums {
		if v%2 == 1 {
			allowed_odd--
		}
		if allowed_odd < 0 {
			return false
		}
	}
	return true

}

func pseudoPalindromicPaths(root *TreeNode) int {

	count := 0
	var stack Stack
	var dfs func(*TreeNode)

	dfs = func(root *TreeNode) {
		stack.PushBack(root.Val)
		defer stack.PopBack()

		if root.Left == nil && root.Right == nil {
			if isPseudoPalindrom(stack) {
				count++
			}
			return
		}

		if root.Right != nil {
			dfs(root.Right)
		}
		if root.Left != nil {
			dfs(root.Left)
		}
	}
	dfs(root)
	return count
}

func pseudoPalindromicPaths_bit(root *TreeNode) int {

	ans := 0
	var dfs func(*TreeNode, int)

	dfs = func(root *TreeNode, count int) {
		count ^= 1 << root.Val
		if root.Left == nil && root.Right == nil {
			if count&(count>>1) == 0 {
				ans++
			}
			return
		}

		if root.Right != nil {
			dfs(root.Right, count)
		}
		if root.Left != nil {
			dfs(root.Left, count)
		}
	}
	dfs(root, 0)
	return ans
}
