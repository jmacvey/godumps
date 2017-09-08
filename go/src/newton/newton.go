package main;

import (
	"fmt"
);

func sqrt(x float64, numIterations int) float64 {
	z := float64(1.0);
	for i := 0; i < numIterations; i++ {
		z -= ((z * z) - x) / (2 * z);
	}
	return z;
}

func main() {
	for i:= 1; i < 10; i++ {
		fmt.Println("Sqrt(", i, "):", sqrt(float64(i), 10));
	}
}