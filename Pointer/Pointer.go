package main

import "fmt"

func main() {

	var x int = 10
	var xPoint *int = &x

	addBValue(x)
	fmt.Println("in main:", x)
	addByReference(xPoint)
	fmt.Println("in main:", x)
}

func addByReference(x *int) {
	*x += 10
	fmt.Println("in addByReference:", *x)
}

func addBValue(x int) {
	x += 10
	fmt.Println("in addBValue:", x)
}
