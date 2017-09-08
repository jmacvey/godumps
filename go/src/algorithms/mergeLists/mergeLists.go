package mergeLists

// ListNode Struct
type ListNode struct {
	Value interface{}
	Next  *ListNode
}

func InitializeList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := &ListNode{Value: arr[0], Next: nil}
	curr := head
	for _, v := range arr[1:] {
		curr.Next = &ListNode{Value: v, Next: nil}
		curr = curr.Next
	}
	return head
}

func mergeTwoLinkedLists(l1 *ListNode, l2 *ListNode) *ListNode {
	newHead := &ListNode{Value: nil, Next: nil}

	// 1 -> 2 -> 3 -> 4
	// 1 -> 2 -> 3 -> 4
	// 1 -> 1 -> 2 -> 2 -> 3 -> 3 -> 4 -> 4

	// For each iteration, compare the currA to currB
	// branch with the current on the other
	// While A && B != nil
	// If A <= B, push value A & set currA forward
	// otherwise push value B and set currB forward

	// while A != nil
	// push remaining A
	// while B != nil
	// push remaining B

	curr1 := l1
	curr2 := l2
	currNew := newHead

	for curr1 != nil && curr2 != nil {
		if curr1.Value.(int) <= curr2.Value.(int) {
			currNew.Next = &ListNode{Value: curr1.Value, Next: nil}
			curr1 = curr1.Next
		} else {
			currNew.Next = &ListNode{Value: curr2.Value, Next: nil}
			curr2 = curr2.Next
		}
		currNew = currNew.Next
	}

	for curr1 != nil {
		currNew.Next = &ListNode{Value: curr1.Value, Next: nil}
		curr1 = curr1.Next
		currNew = currNew.Next
	}

	for curr2 != nil {
		currNew.Next = &ListNode{Value: curr2.Value, Next: nil}
		curr2 = curr2.Next
		currNew = currNew.Next
	}

	return newHead.Next
}
