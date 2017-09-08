package LinkedList

import (
	"fmt"
	"node"
)

/*
 * Convention to call constructors New{ClassName} apparently
 * Normally they'd return a pointer since it's GC'd anyway,
 * but this is just POC
 * http://www.golangpatterns.info/object-oriented/constructors
 */
func NewLinkedList() LinkedList {
	return LinkedList{Head: nil, End: nil, Length: 0}
}

func InitializeList(initial []int) LinkedList {
	list := LinkedList{Head: nil, End: nil, Length: 0}
	// range based for loop
	for _, v := range initial {
		list.PushBack(v)
	}
	return list
}

/*
 * Structs are classes
 * + Private Data Fields are camel cased
 * + Public Data Fields are Upper Camel
 */
type LinkedList struct {
	Head   *node.Node
	End    *node.Node
	Length int
}

/*
 * Public functions are Upper-cased
 */
func (list *LinkedList) AddFront(item int) {
	if list.Head == nil {
		list.Head = &node.Node{Item: item, Next: nil, Prev: nil}
	} else {
		list.Head.Prev = &node.Node{Item: item, Next: list.Head, Prev: nil}
		list.Head = list.Head.Prev
	}
	list.Length++
}

func (list *LinkedList) PushBack(item int) {
	if list.Head == nil {
		list.AddFront(item)
	} else {
		var prev *node.Node = list.End
		toAdd := node.Node{Item: item, Next: nil, Prev: prev}

		if prev == nil { // update head (old list end since this is a 1 item list) to point to the new end
			toAdd.Prev = list.Head
			list.Head.Next = &toAdd
		} else { // update old list end to point to the new end
			list.End.Next = &toAdd
		}
		list.End = &toAdd
		list.Length++
	}
}

/*
 * Displays list in forward or reverse fashion
 */
func (list *LinkedList) Display(reverse bool) {
	if !reverse {
		for ptr := list.Head; ptr != nil; ptr = ptr.Next {
			fmt.Println(ptr.Item)
		}
	} else {
		for ptr := list.End; ptr != nil; ptr = ptr.Prev {
			fmt.Println(ptr.Item)
		}
	}
}

func (list *LinkedList) Clear() {
	// no need to release memory since Go does this for you
	list.Head = nil
	list.End = nil
	list.Length = 0
}

type ListNode struct {
	Value int
	Next  *ListNode
}

func removeFromList(l *ListNode, k int) *ListNode {
	// step 1: store the rference to the head of the list that isn't the value
	head := l
	for head != nil && head.Value == k {
		head = head.Next
	}

	// step 2: remove the elements from that point forward that match the value
	prev := head
	for prev != nil {
		for prev.Next != nil && prev.Next.Value == k {
			prev.Next = prev.Next.Next
		}
		prev = prev.Next
	}

	return head
}
