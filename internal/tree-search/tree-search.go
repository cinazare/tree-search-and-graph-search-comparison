package treesearch

import (
	"time"
	"tree-search-and-graph-search-comparison/internal/graph"
	"tree-search-and-graph-search-comparison/internal/utils"
	"fmt"
)

const MaxDepth = 3000

func Run(graph graph.Graph, index int) int {
	reportBFS := utils.Report{}
	reportDFS := utils.Report{}

	bfsRes := BFS("A", "F", graph, &reportBFS)


	dfsRes := DFS("A", "F", graph, &reportDFS)
	reportBFS.SaveToFile("Tree Search BFS", bfsRes, index)
	reportDFS.SaveToFile("Tree Search DFS", dfsRes, index)
	return 0
}

func DFS(start string, end string, graph graph.Graph, report *utils.Report) bool {
	startTime := time.Now()

	result := dfsHelper(start, end, graph, report, 1)

	report.Time_Running = float32(time.Since(startTime).Nanoseconds())
	return result
}

func dfsHelper(
	current string,
	end string,
	graph graph.Graph,
	report *utils.Report,
	depth int,
) bool {

	if depth > MaxDepth {
		fmt.Println("Search stopped: maximum depth reached")
		return false
	}

	report.Nodes_Proccessed++

	if float32(depth) > report.Storage_Usage {
		report.Storage_Usage = float32(depth)
	}

	if current == end {
		return true
	}

	neighbors := graph.Nodes[current]
	report.Nodes_Generated += len(neighbors)

	for _, neighbor := range neighbors {
		if dfsHelper(neighbor, end, graph, report, depth+1) {
			return true
		}
	}

	return false
}


func BFS(start string, end string, graph graph.Graph, report *utils.Report) bool {
	startTime := time.Now()

	queue := []string{start}
	report.Nodes_Generated = 1

	for len(queue) > 0 {
		if float32(len(queue)) > report.Storage_Usage {
			report.Storage_Usage = float32(len(queue))
		}

		current := queue[0]
		queue = queue[1:]

		report.Nodes_Proccessed++

		if current == end {
			report.Time_Running = float32(time.Since(startTime).Nanoseconds())
			return true
		}

		neighbors := graph.Nodes[current]

		report.Nodes_Generated += len(neighbors)

		queue = append(queue, neighbors...)
	}

	report.Time_Running = float32(time.Since(startTime).Nanoseconds())
	return false
}
