package treesearch

import (
	"fmt"
	"tree-search-and-graph-search-comparison/internal/graph"
)

func Run(graph graph.Graph) int {
	fmt.Println("----------------------------------------------------")
	fmt.Println("in graph search")
	
	fmt.Println(graph)
	return 0
}

func DFS(start string, end string, graph graph.Graph) bool{
	if start == end {
		return true
    }
	
    for _, neighbor := range graph.Nodes[start] {
		if DFS(neighbor, end, graph) {
			return true
        }
    }
	
    return false
}


func BFS(start string, end string, graph graph.Graph) bool{
	queue := []string{start}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		
		if current == end{
			return true
		}

		queue = append(queue, graph.Nodes[current]...)
	}

	return false
}
