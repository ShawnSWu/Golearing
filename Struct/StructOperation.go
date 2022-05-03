package main

import "fmt"

type Person struct {
	Name string //
	age  int
}

type Point struct {
	X int
	Y int
}

func main() {
	p := Person{Name: "shawn", age: 20}
	p2 := p

	p.age = 25
	fmt.Println(p)  //{shawn 25}
	fmt.Println(p2) //{shawn 20}
}
