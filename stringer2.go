package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"haha", 18}
	b := Person{"hehe", 8}
	fmt.Println(a, b)
}
