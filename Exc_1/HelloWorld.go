package main

import "fmt"

// Define a struct
type Person struct {
	Name string
	Age  int
}

// Method with receiver type to print struct contents
func (p Person) Print() {
	fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
}

func main() {
	p := Person{Name: "Karim", Age: 20}
	p.Print()
}
