package main

/* implementation of the Rank Selection problem */

import (
	"fmt"
	"math"
	"sort"
)

const maxValue = 1 << 32

func getMedian(i, j int, arr []int) int {
	sort.Ints(arr[i:j])
	return arr[(i+j)/2]
}

func getMedianLength(n int) int {
	return int(math.Ceil(float64(n / 5)))
}

func getMedianOfMedians(j int, median []int) int {
	if j == 1 {
		return median[j-1]
	}
	return ithSmallest(j/2, 0, j-1, median)
}

func partition(arr []int, l, r, x int) int {
	var temp int
	i, j := l-1, l
	for ; j <= r; j++ {
		if arr[j] < x {
			i++
			temp = arr[j]
			arr[j] = arr[i]
			arr[i] = temp
		}
	}
	i++
	for j = i; j <= r; j++ {
		if arr[j] == x {
			break
		}
	}
	temp = arr[j]
	arr[j] = arr[i]
	arr[i] = arr[j]
	return i
}

func ithSmallest(i, l, r int, arr []int) int {
	if i < 0 || i > (r-l+1) {
		return maxValue
	}

	var j int
	n := r - l + 1
	median := make([]int, getMedianLength(n))
	for j := 0; j < n/5; j++ {
		median[j] = getMedian(l+(j*5), l+(j*5)+5, arr)
	}
	if j*5 < n {
		median[j] = getMedian(l+(j*5), l+(j*5)+n%5, arr)
		j++
	}

	x := getMedianOfMedians(j, median)
	q := partition(arr, l, r, x)
	k := q - l + 1

	if i == k {
		return x
	} else if i < k {
		return ithSmallest(i, l, q-1, arr)
	} else {
		return ithSmallest(i-k, q+1, r, arr)
	}
}

/* program logic to be verified */
func main() {
	fmt.Println("Please enter the number of values, the index")
	fmt.Println("& then the values themselves :-")

	var n, i int
	fmt.Scanln(&n, &i)
	arr := make([]int, n)
	for j := range arr {
		fmt.Scan(&arr[j])
	}

	ith := ithSmallest(i, 0, n-1, arr)
	fmt.Println("ith Smallest element:", arr[ith])
}
