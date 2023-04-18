package main

import (
	"fmt"
	"testing"
)

func Test_sliceToLinkedList(t *testing.T) {
	inout := []struct {
		in   []int
		want *ListNode
	}{
		{[]int{1, 2, 2, 1},
			&ListNode{Val: 1,
				Next: &ListNode{Val: 2,
					Next: &ListNode{Val: 2,
						Next: &ListNode{Val: 1,
							Next: nil}}}}},
		{[]int{1, 2},
			&ListNode{Val: 1,
				Next: &ListNode{Val: 2,
					Next: nil}}},

		{[]int{1},
			&ListNode{Val: 1,
				Next: nil}},

		{[]int{1, 2, 3, 2, 1},
			&ListNode{Val: 1,
				Next: &ListNode{Val: 2,
					Next: &ListNode{Val: 3,
						Next: &ListNode{Val: 2,
							Next: &ListNode{Val: 1,
								Next: nil}}}}}},
		{[]int{1, 2, 3, 4, 2, 1},
			&ListNode{Val: 1,
				Next: &ListNode{Val: 2,
					Next: &ListNode{Val: 3,
						Next: &ListNode{Val: 4,
							Next: &ListNode{Val: 2,
								Next: &ListNode{Val: 1,
									Next: nil}}}}}}},
	}
	for _, v := range inout {
		got := sliceToLinkedList(v.in)
		for nodeg, nodew := got, v.want; ; nodeg, nodew = nodeg.Next, nodew.Next {
			if nodeg == nil || nodew == nil {
				if !(nodeg == nil && nodew == nil) {
					t.Errorf("length mismatch")
				}
				break
			}
			if nodeg.Val != nodew.Val {
				t.Errorf("got:%v,want:%v", nodeg.Val, nodew.Val)
			}
		}
	}

}

func Test_isPalindrome(t *testing.T) {
	test_isPalindromes(t, isPalindrome)
}

func Test_isPalindrome2(t *testing.T) {
	test_isPalindromes(t, isPalindrome2)
}

func test_isPalindromes(t *testing.T, target func(*ListNode) bool) {

	type Case struct {
		label string
		in    *ListNode
		out   bool
	}

	part := func(ints []int, b bool) Case {
		return Case{fmt.Sprintln(ints), sliceToLinkedList(ints), b}
	}

	inout := []struct {
		label string
		in    *ListNode
		out   bool
	}{
		part([]int{1, 2, 2, 1}, true),
		part([]int{1, 2}, false),
		part([]int{1}, true),

		part([]int{1, 2, 3, 2, 1}, true),
		part([]int{1, 2, 3, 4, 2, 1}, false),
		part([]int{1, 2, 3, 2, 3, 1}, false),

		{"1,2,2,1",
			&ListNode{Val: 1,
				Next: &ListNode{Val: 2,
					Next: &ListNode{Val: 2,
						Next: &ListNode{Val: 1,
							Next: nil}}}}, true},
		{"1,2",
			&ListNode{Val: 1,
				Next: &ListNode{Val: 2,
					Next: nil}}, false},

		{"1",
			&ListNode{Val: 1,
				Next: nil}, true},

		{"1,2,3,2,1",
			&ListNode{Val: 1,
				Next: &ListNode{Val: 2,
					Next: &ListNode{Val: 3,
						Next: &ListNode{Val: 2,
							Next: &ListNode{Val: 1,
								Next: nil}}}}}, true},
		{"1,2,3,4,2,1",
			&ListNode{Val: 1,
				Next: &ListNode{Val: 2,
					Next: &ListNode{Val: 3,
						Next: &ListNode{Val: 4,
							Next: &ListNode{Val: 2,
								Next: &ListNode{Val: 1,
									Next: nil}}}}}}, false},
		{"1,2,3,2,3,1",
			&ListNode{Val: 1,
				Next: &ListNode{Val: 2,
					Next: &ListNode{Val: 3,
						Next: &ListNode{Val: 2,
							Next: &ListNode{Val: 3,
								Next: &ListNode{Val: 1,
									Next: nil}}}}}}, false},
	}

	for _, v := range inout {

		// t.Log(v)
		if out := target(v.in); v.out != out {
			t.Errorf("%s should be %v,but %v", v.label, v.out, out)
		} else {
			t.Logf("PASS:%s is %v ,and %v", v.label, v.out, out)
		}
	}
}
