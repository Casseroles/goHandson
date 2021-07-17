package main

import "fmt"

// Data is interface.
type Data interface {
	Initial(name string, data []int)
	PrintData()
}

// MyData is Struct.
type MyData struct {
	Name string
	Data []int
}

// Initial is innit method.
func (md *MyData) PrintData() {
	fmt.Println("Name; ", md.Name)
	fmt.Println("Data; ", md.Data)
}

func main() {
	var ob MyData = MyData{}
	ob.Initial("Sachiko", []init{55, 66, 77})
	ob.PrintData()
}
