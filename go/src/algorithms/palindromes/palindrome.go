package main

import (
	"fmt"
)

// ListNode - List structure
type ListNode struct {
	Value int
	Next  *ListNode
}

func displayList(l *ListNode) {
	for l != nil {
		fmt.Println(l.Value)
		l = l.Next
	}
}

// creates a list from a slice of integers
func initializeList(arr []int) *ListNode {
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

// Returns
// 1. First half of list (L)
// 2. Second half of list (R) where len(R) = len(L) + 1 if total length is odd
// 3. Whether or not we've an odd length
func breakListInHalf(l *ListNode) (*ListNode, *ListNode, bool) {
	length := 0
	curr := l
	for curr != nil {
		length++
		curr = curr.Next
	}
	mid := length / 2
	curr = l

	// leave the mid point on the right half (only go to one)
	// otherwise reordering the list would be terrible
	for mid > 1 {
		curr = curr.Next
		mid--
	}

	firstHalf := l
	secondHalf := curr.Next
	curr.Next = nil

	return firstHalf, secondHalf, length%2 != 0
}

// reverses a list in place
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

func listsEqual(l *ListNode, r *ListNode) bool {
	for l != nil && r != nil {
		if l.Value != r.Value {
			return false
		}
		l = l.Next
		r = r.Next
	}
	return l == nil && r == nil
}

// Gets the last element in a list
func getLast(l *ListNode) *ListNode {
	for l != nil && l.Next != nil {
		l = l.Next
	}
	return l
}

// isPalindrome - Determines if list is palindrom in O(N) time, O(1) space
// Produces no side effects on the list
func isPalindrome(l *ListNode) bool {
	if l == nil {
		return true
	}

	firstHalf, secondHalf, isOdd := breakListInHalf(l) // O(N)

	// single length list
	if secondHalf == nil {
		return true
	}

	// if we're odd, we want to ignore the middle element
	toCheck := secondHalf
	if isOdd {
		toCheck = secondHalf.Next
	}

	// reverse first list and check it against the second
	firstHalf = reverseList(firstHalf)
	toReturn := listsEqual(firstHalf, toCheck) // O(N)

	// restore the first half and concatenate it back with the original
	firstHalf = reverseList(firstHalf)
	getLast(firstHalf).Next = secondHalf

	return toReturn
}

func main() {
	l := initializeList([]int{1, 2, 3, 4, 5, 6, 6, 5, 4, 3, 2, 1})

	fmt.Println("IS PLAINDROME", isPalindrome(l))
	displayList(l) // <- should output list unchanged
}
