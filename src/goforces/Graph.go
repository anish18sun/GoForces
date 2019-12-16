package main

/* implementation of Simple Graph */

import (
	"fmt"
)

/* adjList should always be pointer type */
type vertex struct {
	value, color int
	adjList      []*vertex
}

/* normal function - for dfs()
 * takes vertex object(pointer), not vertex value
 */
func dfs(v *vertex) {
	v.color = 1
	fmt.Print(v.value, " ")
	for _, u := range v.adjList {
		if u.color == 0 {
			dfs(u)
		}
	}
}

/* receiver function - method form
 * test (without interface) or implementation
 */
func (v *vertex) dfs() {
	v.color = 1
	fmt.Print(v.value, " ")
	for _, u := range v.adjList {
		if u.color == 0 {
			u.dfs()
		}
	}
}

func main() {
	fmt.Println("Please enter the number of vertices & edges in the graph :- ")

	var n, m int
	fmt.Scanln(&n, &m)
	graph := make([]vertex, n)
	for i := range graph {
		graph[i] = vertex{i + 1, 0, make([]*vertex, 0)}
	}

	var uInd, vInd int
	fmt.Println("Please enter the edges of the graph in the adjacency list format :-")
	for i := 0; i < m; i++ {
		fmt.Scanln(&vInd, &uInd)
		/* input for graph type should always be taken this way - adding reference */
		graph[vInd-1].adjList = append(graph[vInd-1].adjList, &graph[uInd-1])
	}

	/** printing adjacency list - shows mem address */
	fmt.Println("The elements of the graph are :-")
	for _, v := range graph {
		fmt.Println("Node:", v.value, ":", v.adjList)
	}

	dfs(&graph[0])
	fmt.Println()
	// refresh graph at this point to decolor all vertices
	graph[0].dfs()
	fmt.Println()
}
