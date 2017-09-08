package main;

import "fmt";

func main() {
	s, i := 1.012, 3.1234;
	fmt.Println("Before changing: ", s);
	var p *float64 = &s;
	*p = *p + 1;
	fmt.Println("After chaning: ", *p, s, &s, p, i);
	
}

