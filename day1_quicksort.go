package main

import "fmt"

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func SolutionBinaryGap(N int) int {
	CurLen := 0
	MaxLen := 0
	for N > 0 {
		if N%2 == 0 {
			CurLen += 1
			MaxLen = Max(MaxLen, CurLen)
		} else {
			CurLen = 0
		}
		N = N / 2
	}
	return MaxLen
}
func main() {
	number := 10000
	fmt.Println(SolutionBinaryGap(number))

}
