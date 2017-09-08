package rearrange

import "fmt"

type ListNode struct {
	Value interface{}
	Next  *ListNode
}

func DisplayList(l *ListNode) {
	for l != nil {
		fmt.Print(l.Value, "-> ")
		l = l.Next
	}
	fmt.Println("")
}

func InitializeList(arr []int) *ListNode {
	head := &ListNode{Value: arr[0], Next: nil}
	curr := head
	for _, v := range arr[1:] {
		curr.Next = &ListNode{Value: v, Next: nil}
		curr = curr.Next
	}
	return head
}

func RearrangeLastN(l *ListNode, k int) *ListNode {
	lengthOfList := 0
	end := l
	for curr := l; curr != nil; curr = curr.Next {
		lengthOfList++
		if curr.Next == nil {
			end = curr
		}
	}

	curr := l
	goTo := lengthOfList - k
	for i := 0; i < goTo; i++ {
		curr = curr.Next
	}

	toReturn := curr.Next
	curr.Next = nil

	return toReturn
}
