package main

/* @author anish
*  @description: program to implement Topological Sort
 */

import (
	"container/list"
	"fmt"
)

type vertex struct {
	value, color int
	adjList      []*vertex
}

func (v *vertex) dfs(topoList list) {
	v.color = 1
	fmt.Println(v.value, " ")
	for _, u := range v.adjList {
		if u.color == 0 {
			u.dfs()
		}
	}
	topoList.pushFront(v)
}

func main() {
	fmt.Println("Please enter the number of vertices & edges in the graph:-")

	var n, m int
	fmt.Scanln(&n, &m)
	graph := make([]vertex, n)
	for i := range graph {
		graph[i] = vertex{i + 1, 0, make([]*vertex, 0)}
	}

	var u, v int
	for i := 0; i < m; i++ {
		fmt.Scanln(&v, &u)
		graph[v-1].adjList = append(graph[v-1].adjList, &graph[u-1])
	}

	topoList := list.New()
	graph[0].dfs(topoList)

	fmt.Println("TopoSort on the verties: ", topoList)
}
