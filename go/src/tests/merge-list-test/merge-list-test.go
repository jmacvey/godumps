package main

import (
	"reverseNodes"
)

func main() {
	l := reverseNodes.InitializeList([]int{1, 2, 3, 4, 5, 6})
	reverseNodes.ReverseNodesInKGroups(l, 2)

	l2 := reverseNodes.InitializeList([]int{1, 2, 3, 4, 5, 6, 7, 8})
	reverseNodes.ReverseNodesInKGroups(l2, 5)
}
