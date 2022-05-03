package main

import "fmt"

type Company struct {
	Name string //
	employeeCount  int
}


func main() {
	p := Company{Name: "IBM", employeeCount: 300}

	updateStruct(p)
	fmt.Println(p)  // {IBM 300}

	updateStructWithPointer(&p)
	fmt.Println(p)  // {Google 303}
}

func updateStruct(p Company) {
	p.Name = "Google"
	p.employeeCount += 3
}

func updateStructWithPointer(p *Company) {
	p.Name = "Google"
	p.employeeCount += 3
}

