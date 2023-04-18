package main

import (
	"fmt"
)

func main() {

	list := sliceToLinkedList([]int{1, 2, 4, 3, 1})
	// fmt.Println(tail)
	// fmt.Println(tail.Next)
	// fmt.Println(tail.Next.Next)
	// fmt.Println(tail.Next.Next.Next)
	isPalindrome2(list)
	fmt.Println(list.itrString())

}

func sliceToLinkedList(nums []int) *ListNode {

	root := new(ListNode)
	parent := root
	for i, v := range nums {
		parent.Val = v

		// except last node
		if i+1 < len(nums) {
			parent.Next = new(ListNode)
			parent = parent.Next
		}
	}
	return root
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// スライスに変換して両端から順に一致しているか確認 time:O(n),space:O(n)
func isPalindrome(head *ListNode) bool {
	sl := make([]int, 0)

	for node := head; node != nil; node = node.Next {
		sl = append(sl, node.Val)

	}

	len := len(sl)
	for i := 0; i < len/2; i++ {
		if sl[i] != sl[len-1-i] {
			return false
		}
	}
	return true
}

// headに破壊的変更あり time:O(n),space:O(1)
func reverse(head *ListNode) *ListNode {
	cur, prev := head, (*ListNode)(nil)

	for cur != nil {
		next := cur.Next
		cur.Next = prev

		prev = cur
		cur = next
	}
	return prev

}

func (head *ListNode) itrString() string {
	s := ""

	for node := head; node != nil; node = node.Next {
		// s += fmt.Sprintf("%v %p", node.Val, node.Next)
		s += fmt.Sprintf("%v\n", node)

	}
	return s
}

// headに破壊的変更あり time:O(n),space:O(1)
func isPalindrome2(head *ListNode) bool {
	// sl := make([]int, 0)
	half := 0
	step2x, step1x := head, head
	for ; step2x != nil && step2x.Next != nil; step2x, step1x = step2x.Next.Next, step1x.Next {
		half++
	}

	tail := reverse(step1x)

	for i, headside, tailside := 0, head, tail; i < half; i++ {
		if headside.Val != tailside.Val {
			return false
		}
		headside, tailside = headside.Next, tailside.Next
	}

	return true
}
