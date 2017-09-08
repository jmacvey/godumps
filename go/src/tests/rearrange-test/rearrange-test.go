package main

import (
	"rearrange"
)

func main() {
	l := rearrange.InitializeList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	l = rearrange.RearrangeLastN(l, 0)
	rearrange.DisplayList(l)
}
