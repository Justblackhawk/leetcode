package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	var cur *ListNode
	cur = head
	for cur != nil && cur.Next != nil{
		tmp := cur.Val
		cur.Val = cur.Next.Val
		cur.Next.Val = tmp
		cur = cur.Next.Next
	}
	return head
}

func swapPairs1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	var dummy *ListNode = &ListNode{0,head}
	dummy.Next = head
	for p := dummy;p.Next !=nil && p.Next.Next!=nil;{
		a := p.Next
		b := a.Next
		p.Next = b
		p = a
		a.Next = b.Next
		b.Next = a
	}
	return dummy.Next
}