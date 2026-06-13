package utils

import (
	"fmt"
	"os"
)

type Report struct {
	Time_Running     float32
	Storage_Usage    float32
	Nodes_Generated  int
	Nodes_Proccessed int
}


func (r *Report) Print(result bool) {
	fmt.Println("storage usage: ", r.Storage_Usage)
	fmt.Println("generated nodes: ", r.Nodes_Generated)
	fmt.Println("proccessed nodes: ", r.Nodes_Proccessed)
	fmt.Println("time running :", r.Time_Running)
}


func (r *Report) SaveToFile(algorithm string, result bool, graphIndex int) error {
	file, err := os.OpenFile(
		"report.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = fmt.Fprintf(
		file,
		"\n==================================================\n"+
			"Algorithm: %s\n"+
			"Graph Index: %d\n"+
			"Result: %v\n"+
			"Storage Usage (nodes): %.0f\n"+
			"Generated Nodes: %d\n"+
			"Processed Nodes: %d\n"+
			"Time Running: %.0f ns\n",
		algorithm,
		graphIndex,
		result,
		r.Storage_Usage,
		r.Nodes_Generated,
		r.Nodes_Proccessed,
		r.Time_Running,
	)

	return err
}

func DeleteReportFile() error {
	if _, err := os.Stat("report.txt"); err == nil {
		return os.Remove("report.txt")
	}

	return nil
}