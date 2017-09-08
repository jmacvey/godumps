package main;

// import (
// 	"golang.org/x/tour/pic"
// );

func Pic(dx, dy int) [][]uint8 {
	toReturn := make([][]uint8, dy);
	for i := range toReturn {
		toReturn[i] = make([]uint8, dx);
		for j := range toReturn[i] {
			toReturn[i][j] = uint8(j ^ i << 8);
		}
	}
	return toReturn;
}

func main() {
	//pic.Show(Pic)
}