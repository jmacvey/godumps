package twoHugeNumbers

import (
	"math"
)

// ListNode - linked list structure
type ListNode struct {
	Value int
	Next  *ListNode
}

func initializeList(arr []int) *ListNode {
	head := &ListNode{Value: arr[0], Next: nil}
	curr := head
	for _, v := range arr[1:] {
		curr.Next = &ListNode{Value: v, Next: nil}
		curr = curr.Next
	}
	return head
}

func addTwoHugeNumbers(a *ListNode, b *ListNode) *ListNode {
	aNumbers := listToArray(a)
	bNumbers := listToArray(b)
	maxLen := int(math.Max(float64(len(aNumbers)), float64(len(bNumbers))) + 1)

	aNumbers = append(make([]int, maxLen-len(aNumbers)), aNumbers...)
	bNumbers = append(make([]int, maxLen-len(bNumbers)), bNumbers...)

	cNumbers := make([]int, maxLen)
	result, overflow := 0, 0

	for i := maxLen - 1; i >= 0; i-- {
		result, overflow = getOverflow(aNumbers[i], bNumbers[i], overflow)
		cNumbers[i] = result
	}

	if cNumbers[0] == 0 {
		cNumbers = cNumbers[1:]
	}
	return initializeList(cNumbers)
}

func getOverflow(a int, b int, overflow int) (int, int) {
	result := a + b + overflow
	hasRemainder := 9999-result > 0
	carryOver := 0
	if hasRemainder {
		result -= 10000
		carryOver = 1
	}
	return result, carryOver
}

func listToArray(list *ListNode) []int {
	var toReturn []int
	for list != nil {
		toReturn = append(toReturn, list.Value)
		list = list.Next
	}
	return toReturn
}
