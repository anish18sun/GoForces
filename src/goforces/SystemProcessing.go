package main

/* @author anish
 * @description : implmentation of the SystemProcessing problem. Given a set of
 * processing requests that arrive at each instant {x[1], x[2],..., x[n]} and
 * the capacity of the system to process those requests for each instant
 * {s[1], s[2],..., s[n]} after reboot, ie, 1 in s[1] represents 1st second
 * after reboot. Find the times when the system should be rebooted so as to
 * maximize the amount of data processed by the System.(Dynamic Programming)
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

func sum(i, j int, s, x []int) int {
	pSum := 0
	for k := i; k <= j; k++ {
		pSum += min(s[k-i], x[k-1])
	}
	return pSum
}

func getRebootTimes(x, s []int, n int) []int {
	opt := make([]int, n+1)
	sol := make([]int, n+1)
	opt[1] = min(x[0], s[1])

	for j := 2; j < n+1; j++ {
		for i := 1; i <= j; i++ {
			cost := opt[i-1] + sum(i, j, s, x)
			if opt[j] < cost {
				opt[j] = cost
				sol[j] = i
			}
		}
	}

	fmt.Println("Max Processing:", opt[n])
	return sol
}

func main() {
	fmt.Println("Please enter the number of requests & then the")
	fmt.Println("requests and processing capacities after reboot :-")

	var n int
	fmt.Scanln(&n)
	x := make([]int, n)
	s := make([]int, n+1)
	for i := range x {
		fmt.Scan(&x[i])
	}
	for i := 1; i < n+1; i++ {
		fmt.Scan(&s[i])
	}

	rebootTimes := getRebootTimes(x, s, n)
	fmt.Println("The times when reboot required", rebootTimes)
}
