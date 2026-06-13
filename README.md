# Tree Search and Graph Search Comparison

## Overview

This project is a comparison between **Tree Search** and **Graph Search** approaches using two classical search algorithms:

* Breadth-First Search (BFS)
* Depth-First Search (DFS)

The goal of the project is to compare the behavior and performance of these algorithms under different graph structures and search strategies.

For each execution, the project records:

* Search result (found or not found)
* Running time
* Number of generated nodes
* Number of processed nodes
* Storage usage

The results are written to a `report.txt` file for later analysis.

---

# Project Structure

```text
.
├── cmd
│   └── main.go
│
├── internal
│   ├── compare
│   │   └── compare.go
│   │
│   ├── graph
│   │   └── graph.go
│   │
│   ├── graphsearch
│   │   └── graphsearch.go
│   │
│   ├── treesearch
│   │   └── treesearch.go
│   │
│   └── utils
│       └── utils.go
│
└── report.txt
```

---

# Folder Descriptions

## cmd

### main.go

This is the entry point of the application.

Responsibilities:

* Initializes the program.
* Loads graph samples.
* Starts the comparison process.
* Calls the comparison module.

The `main.go` file acts as the orchestrator of the entire project.

---

## internal/compare

### compare.go

Responsible for running all search strategies in sequence.

Responsibilities:

* Execute Tree Search BFS.
* Execute Tree Search DFS.
* Execute Graph Search BFS.
* Execute Graph Search DFS.
* Collect and save results.

This module provides a fair comparison between all algorithms.

---

## internal/graph

### graph.go

Contains the graph implementation and hardcoded graph samples used for testing.

Responsibilities:

* Define the graph structure.
* Add edges between nodes.
* Generate sample graphs.
* Provide different test cases.

Current graph samples include:

1. Reachable goal node.
2. Unreachable goal node.
3. Graphs containing cycles.

These samples are intentionally designed to demonstrate differences between Tree Search and Graph Search.

---

## internal/treesearch

### treesearch.go

Contains implementations of:

* Tree Search BFS
* Tree Search DFS

Characteristics:

* Does not maintain a visited set.
* May revisit nodes multiple times.
* Can enter infinite loops when cycles exist.
* Demonstrates the traditional tree search approach.

Metrics collected:

* Running time
* Storage usage
* Generated nodes
* Processed nodes

---

## internal/graphsearch

### graphsearch.go

Contains implementations of:

* Graph Search BFS
* Graph Search DFS

Characteristics:

* Uses a visited set.
* Prevents revisiting already explored nodes.
* Handles cyclic graphs safely.
* Usually more efficient than Tree Search on cyclic graphs.

Metrics collected:

* Running time
* Storage usage
* Generated nodes
* Processed nodes

---

## internal/utils

### utils.go

Contains utility structures and helper functions.

Responsibilities:

* Report generation.
* Saving results to files.
* Deleting previous reports.
* Printing search statistics.

The `Report` structure stores:

```go
type Report struct {
    Time_Running     float32
    Storage_Usage    float32
    Nodes_Generated  int
    Nodes_Proccessed int
}
```

---

# Algorithms

## Breadth-First Search (BFS)

Breadth-First Search explores nodes level by level.

Example:

```text
A
├── B
├── C
└── D
```

Traversal order:

```text
A → B → C → D
```

Properties:

* Complete for finite graphs.
* Finds the shortest path in unweighted graphs.
* Requires more memory than DFS.

---

## Depth-First Search (DFS)

Depth-First Search explores one branch completely before backtracking.

Example:

```text
A
├── B
│   └── D
└── C
```

Traversal order:

```text
A → B → D → C
```

Properties:

* Requires less memory than BFS.
* Does not guarantee the shortest path.
* Can get trapped in deep branches.

---

# Tree Search

Tree Search treats every generated node as a new node.

Characteristics:

* No visited set.
* Can revisit the same state many times.
* Can loop forever in cyclic graphs.
* Simpler implementation.

Advantages:

* Easy to implement.

Disadvantages:

* Higher memory consumption in repeated states.
* Poor performance on cyclic graphs.

---

# Graph Search

Graph Search maintains a visited set.

Characteristics:

* Previously visited states are not explored again.
* Prevents infinite loops.
* More efficient on cyclic graphs.

Advantages:

* Avoids repeated work.
* Handles cycles safely.
* Generally more efficient.

Disadvantages:

* Requires additional memory for the visited set.

---

# Performance Metrics

Each algorithm records the following metrics:

## Running Time

Execution time measured in nanoseconds.

## Generated Nodes

Number of nodes created or discovered during search.

## Processed Nodes

Number of nodes actually explored.

## Storage Usage

Maximum number of nodes stored simultaneously.

For:

* BFS → Maximum queue size.
* DFS → Maximum recursion depth.

---

# Running the Project

## Prerequisites

Install Go:

https://go.dev/dl/

Verify installation:

```bash
go version
```

---

## Clone the Repository

```bash
git clone <repository-url>
```

Enter the project directory:

```bash
cd tree-search-and-graph-search-comparison
```

---

## Install Dependencies

```bash
go mod tidy
```

---

## Run the Project

```bash
go run ./cmd
```

or

```bash
go run cmd/main.go
```

depending on the project structure.

---

# Output

After execution, a file named:

```text
report.txt
```

will be generated in the project directory.

Example:

```text
==================================================
Algorithm: Graph Search BFS
Graph Index: 0
==================================================
Result: true
Storage Usage (nodes): 12
Generated Nodes: 1000
Processed Nodes: 523
Time Running: 156723 ns
--------------------------------------------------
```

The report contains results for all algorithms and all graph samples.

---

# Purpose of the Project

The primary objective of this project is to study and compare:

* BFS vs DFS
* Tree Search vs Graph Search

through practical implementations and empirical measurements.

The project demonstrates how search strategy influences:

* Execution time
* Memory consumption
* Node generation
* Node processing
* Behavior in cyclic graphs

and provides a clear comparison between classical search approaches used in Artificial Intelligence and Graph Theory.
