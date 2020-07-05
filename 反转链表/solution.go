package 反转链表

type ListNode struct{
	value int
	next *ListNode
}

func reverseList(head *ListNode) *ListNode{
	if head == nil || head.next == nil{
		return head
	}
	// pre 是重置的头指针
	var prev *ListNode = nil
	// cur 指针是整个反转过程中保持迭代前进的指针
	cur := head
	// tmp 临时存储cur下一个节点,以便cur下一轮找到要逆置的那个节点
	var tmp *ListNode
	for cur != nil {
		// 1.（保存一下前进方向）保存下一跳
		tmp = cur.next
		// 2.斩断过去,不忘前事
		cur.next = prev
		// 3.前驱指针的使命在上面已经完成，这里需要更新前驱指针
		prev = cur
		// 当前指针的使命已经完成，需要继续前进了
		cur = tmp
	}
	return prev
}