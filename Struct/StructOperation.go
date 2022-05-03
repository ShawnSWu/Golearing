package main

import "fmt"

type family struct {
	relationship string
	name         string
}

type Person struct {
	Name string
	age  int
	family
}


func main() {
	p := Person{
		Name:   "shawn",
		age:    20,
		family: family{"Mom", "Elizabeth"},
	}
	fmt.Println(p)


	//匿名Struct
	a := struct{name string}{name: "hello"}
	fmt.Println(a.name)
}
