package main

import "fmt"

type employee struct {
	name     string
	salary   int
	position string
}

type company struct {
	companyName string
	employees   []employee
}

func Print(class company) {
	fmt.Println("Company name: ", class.companyName)
	fmt.Println("Employee Details: ")
	for i := 0; i < len(class.employees); i++ {

		fmt.Println("Employee's name: ", class.employees[i].name)
		fmt.Println("Employee's salary: ", class.employees[i].salary)
		fmt.Println("Employee's postion: ", class.employees[i].position)
	}
}

func main() {
	emp1 := employee{name: "Fatima", salary: 9999, position: "CEO"}
	emp2 := employee{name: "Zain", salary: 564, position: "Janitor"}
	emp3 := employee{name: "Musa", salary: 7965, position: "Developer"}

	emplys := []employee{emp1, emp2, emp3}

	Comp := company{companyName: "Fatima Coop", employees: emplys}
	Print(Comp)
}
