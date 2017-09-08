package main

import (
	"fmt"
	"stack"
)

func main() {
	s1 := stack.DefaultStack()

	for i := 0; i < 20; i++ {
		s1.Push(i)
	}

	for i := 0; i < 20; i++ {
		fmt.Println(s1.Pop())
	}
}
