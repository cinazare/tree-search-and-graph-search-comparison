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
	if id == "0"{
		graph := GraphConstructor()
		// Defining path: A -> C -> F
		graph.AddToEdge("A", "B")
		graph.AddToEdge("A", "C")
		graph.AddToEdge("B", "D")
		graph.AddToEdge("B", "E")
		graph.AddToEdge("C", "F") // F is now reachable!
		graph.AddToEdge("E", "F") // Another path to F
		
		return graph, nil
		
	} else if id == "1"{
		graph := GraphConstructor()
		
		graph.AddToEdge("A", "B")
		graph.AddToEdge("A", "C")
		graph.AddToEdge("B", "D")
		graph.AddToEdge("C", "E")
		
		// Component 2 (disconnected from A)
		graph.AddToEdge("F", "G")
	
		return graph, nil
		
	} else if id == "2"{
	
		graph := GraphConstructor()
		// A -> B -> C -> B creates a loop
		// B -> F makes F reachable despite the loop
	
		graph.AddToEdge("A", "B")
		graph.AddToEdge("B", "C")
		graph.AddToEdge("C", "B")
		graph.AddToEdge("B", "F")
	
		return graph, nil
	
	}
	return nil, fmt.Errorf("graph with %s id was not found", id)
}