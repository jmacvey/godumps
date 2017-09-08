package main;

import (
	"LinkedList"
	"fmt"
);

func main() {
	list := LinkedList.NewLinkedList();
	for i := 0; i < 11; i++ {
		list.PushBack(i * 10);
	}
	fmt.Println("List Read Front to Back: ");
	list.Display(false);
	fmt.Println("List Read Back to Front: ");
	list.Display(true);

	fmt.Println("Clearing...");
	list.Clear();

	fmt.Println("initializing with numbers 1 - 10");
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10};
	list = LinkedList.InitializeList(arr);

	list.Display(false);
}
