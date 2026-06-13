package graph

// for now we use hardcode graphs for our use

import (
	"fmt"
)

type Graph struct{
	Nodes map[string][]string
}


func (graph *Graph) AddToEdge(from, to string){
	graph.Nodes[from] = append(graph.Nodes[from], to)

	if _, exists := graph.Nodes[to]; !exists{
		graph.Nodes[to] = []string{}
	}

}

func GraphConstructor() *Graph {
	return &Graph{
		Nodes: make(map[string][]string),
	}
}


func GetSampleHardcodedData(id string) (*Graph, error) {
	if id == "0" {

		graph := GraphConstructor()

		// Graph 0:
		// Contains a cycle.
		// F is reachable.
		// Tree Search DFS may loop forever.
		// Graph Search DFS/BFS should succeed.

		graph.AddToEdge("A", "N0")

		for i := 0; i < 5000; i++ {
			graph.AddToEdge(
				fmt.Sprintf("N%d", i),
				fmt.Sprintf("N%d", i+1),
			)
		}

		// Cycle
		graph.AddToEdge("N5000", "N2500")

		// Goal
		graph.AddToEdge("N4999", "F")

		return graph, nil

	} else if id == "1" {

		graph := GraphConstructor()

		// Graph 1:
		// F is NOT reachable.
		// DFS/BFS must traverse thousands of nodes.

		graph.AddToEdge("A", "N0")

		for i := 0; i < 10000; i++ {
			graph.AddToEdge(
				fmt.Sprintf("N%d", i),
				fmt.Sprintf("N%d", i+1),
			)
		}

		// Disconnected component
		graph.AddToEdge("F", "G")

		return graph, nil

	} else if id == "2" {
				graph := GraphConstructor()

		// Graph 2:
		// F is reachable.
		// Large binary-tree-like graph.

		graph.AddToEdge("A", "N0")

		for i := 0; i < 5000; i++ {
			from := fmt.Sprintf("N%d", i)
			left := fmt.Sprintf("N%d", i*2+1)
			right := fmt.Sprintf("N%d", i*2+2)

			graph.AddToEdge(from, left)
			graph.AddToEdge(from, right)
		}

		graph.AddToEdge("N9998", "F")

		return graph, nil
	}

	return nil, fmt.Errorf("graph with %s id was not found", id)
}