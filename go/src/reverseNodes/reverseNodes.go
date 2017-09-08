package reverseNodes

import "fmt"

type ListNode struct {
	Value interface{}
	Next  *ListNode
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

// O(N) to reverse time, // O(1) space
func reverseList(l *ListNode) *ListNode {
	anchor := l
	newHead := l
	for anchor != nil && anchor.Next != nil {
		tmp := anchor.Next.Next
		anchor.Next.Next = newHead
		newHead = anchor.Next
		anchor.Next = tmp
	}
	l = newHead
	return l
}

func findNodeKIn(l *ListNode, k int) (*ListNode, *ListNode, bool) {
	first := l
	i := 0
	for i < (k-1) && l != nil {
		l = l.Next
		if l != nil {
			i++
		}

	}
	return first, l, i != (k - 1)
}

// Func - ReverseNodesInKGroups
func ReverseNodesInKGroups(l *ListNode, k int) *ListNode {

	var absoluteFirst, c, d *ListNode
	a, b := l, l
	isShort := false

	for a != nil {

		a, b, isShort = findNodeKIn(a, k)

		if b != nil {
			c = b.Next
			b.Next = nil
		} else {
			c = nil
		}

		if isShort {
			if d != nil {
				d.Next = a
			} else {
				absoluteFirst = a
			}
			break
		}

		b = reverseList(a)

		if d != nil {
			d.Next = b
		}

		// save the front absolutely if necessary
		if absoluteFirst == nil {
			absoluteFirst = b
		}
		d = a
		a.Next = c
		a = c
	}
	displayList(absoluteFirst)
	return absoluteFirst
}

func displayList(l *ListNode) {
	for l != nil {
		fmt.Print(l.Value, "-> ")
		l = l.Next
	}
	fmt.Println("")
}
