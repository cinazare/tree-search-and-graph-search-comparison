package graph

// for now we use hardcode graphs for our use 

import "fmt"

type Graph struct{
	Nodes map[string][]string
}


func GraphConstructor(data map[string][]string) *Graph {
	return &Graph{Nodes: data}
}


func GetSampleHardcodedData(id string) (*Graph, error) {
	if id == "1"{
		return GraphConstructor(map[string][]string{
			"A": {"B", "C"},
			"B": {"D", "E"},
			"C": {"A"},
			"D": {},
			"E": {"B"},
		}), nil              
	}
	return nil, fmt.Errorf("graph with %s id was not found", id)
}