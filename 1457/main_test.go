package main

import (
	"testing"
)

func Test_pseudoPalindromicPaths(t *testing.T) {
	test_pseudoPalindromicPaths(t, pseudoPalindromicPaths)
}

func Test_pseudoPalindromicPaths_bit(t *testing.T) {
	test_pseudoPalindromicPaths(t, pseudoPalindromicPaths_bit)
}
func test_pseudoPalindromicPaths(t *testing.T, f func(*TreeNode) int) {
	inout := []struct {
		in  *TreeNode
		out int
	}{
		{&TreeNode{Val: 2,
			Left: &TreeNode{Val: 3,
				Left: &TreeNode{Val: 3,
					Left:  nil,
					Right: nil},
				Right: &TreeNode{Val: 1,
					Left:  nil,
					Right: nil}},
			Right: &TreeNode{Val: 1,
				Left: nil,
				Right: &TreeNode{Val: 1,
					Left:  nil,
					Right: nil}}}, 2},

		{&TreeNode{Val: 2,
			Left: &TreeNode{Val: 1,
				Left: &TreeNode{Val: 1,
					Left:  nil,
					Right: nil},
				Right: &TreeNode{Val: 3,
					Left: nil,
					Right: &TreeNode{Val: 1,
						Right: nil,
						Left:  nil}}},
			Right: &TreeNode{Val: 1,
				Left:  nil,
				Right: nil}}, 1},

		{&TreeNode{Val: 1,
			Left: &TreeNode{Val: 9,
				Left: nil,
				Right: &TreeNode{Val: 1,
					Left:  nil,
					Right: nil}},
			Right: &TreeNode{Val: 1,
				Left: nil,
				Right: &TreeNode{Val: 1,
					Left: &TreeNode{Val: 7,
						Left:  nil,
						Right: nil},
					Right: nil}}}, 1},
	}

	for _, v := range inout {

		if out := f(v.in); v.out != out {
			t.Errorf("(\"%v\") should be %v,but %v", v.in, v.out, out)
		} else {
			t.Logf("PASS:(\"%v\") is %v ,and %v", v.in, v.out, out)
		}
	}
}
