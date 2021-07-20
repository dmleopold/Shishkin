package main

import "fmt"

func SolutionSquareGenerator(start int, n int) []int {
          ArrayOfSquares := []int{}
          for i := start; i < start + n; i++ {
                    ArrayOfSquares = append(ArrayOfSquares, i*i)
          }
          return ArrayOfSquares
}

func main() {
	start := 7
	amount := 3
	fmt.Println(SolutionSquareGenerator(start,amount))
	
}
