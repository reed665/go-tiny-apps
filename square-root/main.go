package main

import "fmt"

func main() {
	input := 4.0
	result := sqrt(input)
	fmt.Printf("The square root of %v equals %v\n", input, result)
}

func sqrt(x float64) float64 {
	return sqrtWithGuess(x, 1.0)
}

func sqrtWithGuess(x float64, guess float64) float64 {
	result := guess
	for i := 0; i < 10; i++ {
		result -= (result*result - x) / (2 * result)
		fmt.Printf("Interim result: %v\n", result)
	}
	return result
}
