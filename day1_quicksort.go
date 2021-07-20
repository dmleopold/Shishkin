/******************************************************************************

                            Online Go Lang Compiler.
                Code, Compile, Run and Debug Go Lang program online.
Write your code in this editor and press "Run" button to execute it.

*******************************************************************************/
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
