package main

/* implementation of the PrintNeatly problem(Dynamic Programming) */

import (
	"fmt"
)

/* This constant should be big enough for the input value L */
const big = 999999

func slack(i, j, L int, w []int) int {
	wLength := 0
	for k := i; k <= j; k++ {
		wLength += w[k-1]
	}
	slackness := L - (wLength + (j - i))
	if slackness < 0 {
		return big + L
	}
	return slackness
}

func lineChangePoints(w []int, n, L int) []int {
	opt := make([]int, n+1)
	sol := make([]int, n+1)
	for i := range opt {
		opt[i] = big + L
	}

	opt[0] = 0
	opt[1] = L - w[0]
	for j := 2; j < n+1; j++ {
		for i := 1; i <= j; i++ {
			cost := opt[i-1] + slack(i, j, L, w)
			if opt[j] > cost {
				opt[j] = cost
				sol[j] = i
			}
		}
	}

	return sol
}

func main() {
	fmt.Println("Please enter the number of words,")
	fmt.Println("the max-line length(L) and the length of each word :-")

	var n, L int
	fmt.Scanln(&n, &L)
	wLength := make([]int, n)
	for i := range wLength {
		fmt.Scan(&wLength[i])
	}

	fmt.Println("The points where lines can be changed:-")
	points := lineChangePoints(wLength, n, L)
	fmt.Println(points)
}
