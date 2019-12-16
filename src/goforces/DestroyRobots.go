package main

/*@author anish
* @description : file to implement the DestroyRobots problem(Dynamic Programming)
 */

import (
	"fmt"
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func getActivationSecs(n int, x, f []int) []int {
	opt := make([]int, n)
	sol := make([]int, n)
	opt[0] = min(x[0], f[0])

	for j := 1; j < n; j++ {
		for i := 0; i <= j; i++ {
			cost := opt[i] + min(x[j], f[j-i])
			if opt[j] < cost {
				opt[j] = cost
				sol[j] = i
			}
		}
	}

	return sol
}

/* verify solution */

func main() {
	fmt.Println("Please enter the number of seconds")
	fmt.Println("& then the robots and the charge values for each second :-")

	var n int
	fmt.Scanln(&n)
	x := make([]int, n)
	f := make([]int, n)
	for i := range x {
		fmt.Scanln(&x[i], &f[i])
	}

	activationSecs := getActivationSecs(n, x, f)
	fmt.Println("Seconds for activation :", activationSecs)
}
