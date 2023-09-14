package main

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

type ListNode = kit.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	curr := dummyHead
	carry := 0

	p, q := l1, l2
	for p != nil || q != nil {
		x := 0
		if p != nil {
			x = p.Val
			p = p.Next
		}

		y := 0
		if q != nil {
			y = q.Val
			q = q.Next
		}

		_sum := x + y + carry
		carry = _sum / 10
		curr.Next = &ListNode{Val: _sum % 10}
		curr = curr.Next
	}

	if carry > 0 {
		curr.Next = &ListNode{Val: carry}
	}

	return dummyHead.Next
}

func main() {
	//var err error
	//
	//fmt.Println("Enter values for the list 1:")
	//_, err =
}
