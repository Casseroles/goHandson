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
func (md *MyData) Initial(name string, data []int) {
	md.Name = name
	md.Data = data
}

// Initial is innit method.
func (md *MyData) PrintData() {
	fmt.Println("Name; ", md.Name)
	fmt.Println("Data; ", md.Data)
}

// Check is method.
func (md *MyData) Check() {

}

func main() {
	var ob Data = new(MyData)
	ob.Initial("Sachiko", []int{55, 66, 77})
	ob.PrintData()
}
