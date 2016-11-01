package main

import (
	"testing"
)

func initTwoListNodes() (*ListNode, *ListNode) {
	l1 := &ListNode{Val:2}
	l2 := &ListNode{Val:5}
	l1.Next = &ListNode{Val:4}
	l2.Next = &ListNode{Val:6}
	l1.Next.Next = &ListNode{Val:3}
	l2.Next.Next = &ListNode{Val:4}
	return l1, l2
}

func Test_addTwoNumbers(t *testing.T) {
	l1, l2 := initTwoListNodes()
	l := addTwoNumbers(l1, l2)
	result := []int{7, 0, 8}
	for _, v := range result {
		if l.Val != v {
			t.Error("Test case 0 failed")
		}
		l = l.Next
	}
}
