package main

import (
	"fmt"
	"strconv"
	"tree-search-and-graph-search-comparison/internal/compare"
	"tree-search-and-graph-search-comparison/internal/graph"
	"tree-search-and-graph-search-comparison/internal/utils"
)

func main(){
	utils.DeleteReportFile()
	for index := range(3){

		fmt.Println("===============================================================================")
		fmt.Println("running for ID: ", index)
		sample_graph, err := graph.GetSampleHardcodedData(strconv.Itoa(index))
		if err != nil{
			// error handeling for later
		}
		compare.Run(*sample_graph, index)
	}
}