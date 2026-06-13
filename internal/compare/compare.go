package compare

import (
	"time"
	"tree-search-and-graph-search-comparison/internal/graph-search"
	"tree-search-and-graph-search-comparison/internal/tree-search"
	"tree-search-and-graph-search-comparison/internal/graph"
)

func Run(graph graph.Graph, index int) string {
    
    start1 := time.Now()
    graphsearch.Run(graph, index)
    t1 := time.Since(start1)	

	start2 := time.Now()
    treesearch.Run(graph, index)
    t2 := time.Since(start2)

    return "Algo1: " + t1.String() + ", Algo2: " + t2.String()

}