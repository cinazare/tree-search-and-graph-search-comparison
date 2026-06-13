package graphsearch

import (
	"time"

	"tree-search-and-graph-search-comparison/internal/graph"
	"tree-search-and-graph-search-comparison/internal/utils"
)

func Run(graph graph.Graph, index int) int {
	reportBFS := utils.Report{}
	reportDFS := utils.Report{}

	bfsRes := BFS("A", "F", graph, &reportBFS)
	dfsRes := DFS("A", "F", graph, &reportDFS)

	reportBFS.SaveToFile("Graph Search BFS", bfsRes, index)
	reportDFS.SaveToFile("Graph Search DFS", dfsRes, index)

	return 0
}

func DFS(start string, end string, graph graph.Graph, report *utils.Report) bool {
	startTime := time.Now()

	visited := make(map[string]bool)

	result := dfsHelper(
		start,
		end,
		graph,
		report,
		visited,
		1,
	)

	report.Time_Running = float32(time.Since(startTime).Nanoseconds())
	return result
}

func dfsHelper(
	current string,
	end string,
	graph graph.Graph,
	report *utils.Report,
	visited map[string]bool,
	depth int,
) bool {

	if visited[current] {
		return false
	}

	visited[current] = true

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
		if dfsHelper(
			neighbor,
			end,
			graph,
			report,
			visited,
			depth+1,
		) {
			return true
		}
	}

	return false
}

func BFS(start string, end string, graph graph.Graph, report *utils.Report) bool {
	startTime := time.Now()

	queue := []string{start}
	visited := map[string]bool{start: true}

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

		for _, neighbor := range graph.Nodes[current] {
			if !visited[neighbor] {
				visited[neighbor] = true
				report.Nodes_Generated++
				queue = append(queue, neighbor)
			}
		}
	}

	report.Time_Running = float32(time.Since(startTime).Nanoseconds())
	return false
}