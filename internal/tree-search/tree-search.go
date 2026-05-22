package treesearch

import (
	"fmt"
	"tree-search-and-graph-search-comparison/internal/graph"
	"tree-search-and-graph-search-comparison/internal/utils"
)


func Run(graph graph.Graph) int {
	report_bfs := utils.Report{Time_Running: 0, Nodes_Generated: 0, Nodes_Proccessed: 0, Storage_Usage: 0}
	report_dfs :=  utils.Report{Time_Running: 0, Nodes_Generated: 0, Nodes_Proccessed: 0, Storage_Usage: 0}

	fmt.Println("----------------------------------------------------")
	fmt.Println("in tree search")
	
	bfs_res := BFS("A", "F", graph, report_bfs)
	dfs_res := DFS("A", "F", graph, report_dfs)

	report_bfs.Print(bfs_res)
	report_dfs.Print(dfs_res)
	
	return 0
}

func DFS(start string, end string, graph graph.Graph, report utils.Report) bool{
	if start == end {
		return true
    }
	
    for _, neighbor := range graph.Nodes[start] {
		if DFS(neighbor, end, graph, report) {
			return true
        }
    }
	
    return false
}


func BFS(start string, end string, graph graph.Graph, report utils.Report) bool{
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
