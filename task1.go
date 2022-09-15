// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type UsingClass struct {
	age  int
	name string
}

func Calling(class UsingClass) {
	fmt.Println("My name is", class.name)
	fmt.Println("My age is", class.age)
}

func main() {
	A := UsingClass{
		age:  21,
		name: "fatima"}
	Calling(A)
}
