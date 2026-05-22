package utils

import "fmt"

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