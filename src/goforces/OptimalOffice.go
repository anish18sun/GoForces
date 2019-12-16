package main

/* @author anish
 * @description : Implementation of the OptimalOffice problem. Given the cost of
 * working/setup in two cities - SF and NY & also the cost of switching from one
 * city to another city. We must determine the appropriate office to work in each
 * in order to minimize the total cost of setup for all n days. (DP)
 */

import (
	"fmt"
)

const maxValue = 1 << 32

func switchSum(office []int, i, j int) int {
	officeCost := 0
	for ; i <= j; i++ {
		officeCost += office[i-1]
	}
	return officeCost
}

func getOptimalOffices(sf, ny []int, n, m int) []int {
	officeSeq := make([]string, n+1)
	opt := make([]int, n+1)
	sol := make([]int, n+1)
	const sF = "SF"
	const nY = "NY"

	/* base case */
	if sf[0] < ny[0] {
		opt[1] = sf[0]
		officeSeq[1] = sF
	} else {
		opt[1] = ny[0]
		officeSeq[1] = nY
	}

	for j := 2; j <= n+1; j++ {
		for i := 1; i <= j; i++ {
			switchInd := 0
			officeCost := 0
			var office string
			if i == 1 {
				sfCost := switchSum(sf, i, j)
				nyCost := switchSum(ny, i, j)
				if sfCost < nyCost {
					office = sF
					officeCost = sfCost
				} else {
					office = nY
					officeCost = nyCost
				}
			} else {
				switchInd = i
				if officeSeq[i-1] == sF {
					office = nY
					officeCost += opt[i-1] + m + switchSum(ny, i, j)
				} else {
					office = sF
					officeCost += opt[i-1] + m + switchSum(sf, i, j)
				}
			}

			if opt[j] > officeCost {
				sol[j] = switchInd
				opt[j] = officeCost
				officeSeq[j] = office
			}
		}
	}

	return sol
}

func main() {
	fmt.Println("Please enter the number of days, the cost of switch")
	fmt.Println("& the cost of operating in SanFransisco and NewYork for those days:-")

	var n, m int
	fmt.Scanln(&n, &m)
	sf := make([]int, n)
	ny := make([]int, n)
	for i := range sf {
		fmt.Scan(&sf[i])
	}
	for i := range ny {
		fmt.Scan(&ny[i])
	}

	switches := getOptimalOffices(sf, ny, n, m)
	fmt.Println("The optimal office switches for minimum cost:", switches)
}
