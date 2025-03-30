package main

import "fmt"

func multiplier(x int) func(int) int {
	return func(y int) int {
		return x * y
	}
}

func main() {
	double := multiplier(2)
	triple := multiplier(3)

	fmt.Println(double(4)) // ??
	fmt.Println(triple(5)) // ??
}

// returning funcitons into functions
// 8
// 15
// the output is 8 and 15
// The multiplier function returns a function that multiplies its argument by x.
// The double variable is assigned the function returned by multiplier(2), which multiplies its argument by 2.
// The triple variable is assigned the function returned by multiplier(3), which multiplies its argument by 3.
// The double(4) call returns 8 because 4 * 2 = 8.
// The triple(5) call returns 15 because 5 * 3 = 15.
// The output is 8 and 15.
