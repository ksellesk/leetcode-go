package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var sum, left int
	head := new(ListNode)
	tmp := new(ListNode)
	tmp = head

	if l1 == nil || l2 == nil {
		return head.Next
	}

	for l1 != nil && l2 != nil {
		newnode := new(ListNode)
		sum = l1.Val + l2.Val + left
		newnode.Val = sum % 10
		left = sum / 10
		tmp.Next = newnode
		tmp, l1, l2 = newnode, l1.Next, l2.Next
	}

	l := new(ListNode)
	if l1 != nil || l2 != nil {
		if l1 == nil {
			l = l2
		} else {
			l = l1
		}

		for l != nil {
			newnode := new(ListNode)
			sum = l.Val + left
			newnode.Val = sum % 10
			left = sum / 10
			tmp.Next = newnode
			tmp, l = newnode, l.Next
		}

	}

	if left != 0 {
		newnode := new(ListNode)
		newnode.Val = left
		tmp.Next = newnode
	}

	return head.Next
}
