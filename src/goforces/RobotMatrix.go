package main

import (
	"container/list"
	"fmt"
)

/* implementation of the RobotMatrix problem */

type pair struct {
	i, j, steps int
}

func getMinSteps(matrix [][]int, n, m int) int {
	iStart, jStart := -1, -1
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 2 {
				iStart = i
				jStart = j
			}
		}
	}

	steps := -1
	if iStart == -1 {
		fmt.Println("No start point found.")
		return steps
	}

	queue := list.New()
	start := pair{iStart, jStart, 0}

	queue.PushBack(start)
	p := queue.Front()
	for ; p != nil; p = p.Next() {
		i, j := p.Value.(pair).i, p.Value.(pair).j
		if i == 0 || i == n-1 || j == 0 || j == m-1 {
			steps = p.Value.(pair).steps
			break
		}

		matrix[i][j] = 1
		for k := -1; k <= 1; k += 2 {
			if matrix[i+k][j] == 0 {
				queue.PushBack(pair{i + k, j, p.Value.(pair).steps + 1})
			}
		}
		for k := -1; k <= 1; k += 2 {
			if matrix[i][j+k] == 0 {
				queue.PushBack(pair{i, j + k, p.Value.(pair).steps + 1})
			}
		}
	}

	return steps
}

func main() {
	fmt.Println("Please enter the dimensions of the matrix ")
	fmt.Println("& then please enter the matrix for robot traversal:-")

	var n, m int
	fmt.Scanln(&n, &m)
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, m)
	}

	for i := range matrix {
		for j := range matrix[i] {
			fmt.Scan(&matrix[i][j])
		}
	}

	steps := getMinSteps(matrix, n, m)
	fmt.Println("Number of steps to get to end :", steps)
}
