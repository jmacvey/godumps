package main;

import (
	"fmt"
	"runtime"
);

func main() {
	fmt.Println("Go is running on:");
	switch os := runtime.GOOS; os {
		case "darwin": 
			fmt.Println("OSX");
		case "linux":
			fmt.Println("Linux.");
		case "windows":
			fmt.Println("Bill Gates D.");
		default:
			fmt.Println(os);
	}
}