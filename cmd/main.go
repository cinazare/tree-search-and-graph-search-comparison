package main



import (
	"tree-search-and-graph-search-comparison/internal/compare"
	"tree-search-and-graph-search-comparison/internal/graph"

)

func main(){
	id := "1"
	sample_graph, err := graph.GetSampleHardcodedData(id)
	if err != nil{
		// error handeling for later
	}
	compare.Run(*sample_graph)
}