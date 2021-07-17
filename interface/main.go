package main

import "fmt"

// Data is interface.
type Data interface {
	Initial(name string, data []int)
	PrintData()
}

// Mydata is Struct.
type Mydata struct {
	Name string
	Data []int
}

// Initial is innit method.
func (md *Mydata) Initial(name string, data []int) {
	md.Name = name
	md.Data = data
}

// Initial is innit method.
func (md *Mydata) PrintData() {
	fmt.Println("Name; ", md.Name)
	fmt.Println("Data; ", md.Data)
}

// Check is method.
func (md *Mydata) Check() {
	fmt.Printf("Check! [%s]", md.Name)

}

func main() {
	var ob Mydata = Mydata{}
	ob.Initial("Sachiko", []int{55, 66, 77})
	ob.Check()
}
