package main

import (
	"fmt"
	"strings"
)

type Student struct {
	rollnumber int
	name       string
	address    string
}

func NewStudent(rollno int, name string, address string) *Student {
	s := new(Student)
	s.rollnumber = rollno
	s.name = name
	s.address = address
	return s
}

type StudentList struct {
	list []*Student
}

func (ls *StudentList) CreateStudent(rollno int, name string, address string) *Student {
	st := NewStudent(rollno, name, address)
	ls.list = append(ls.list, st)
	return st
}

func (ls *StudentList) Print() {
	for i := 0; i < len(ls.list); i++ {
		fmt.Printf("%s List %d %s \n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		fmt.Println(ls.list[i].name)
		fmt.Println(ls.list[i].rollnumber)
		fmt.Println(ls.list[i].address)
	}
}

func main() {
	student := new(StudentList)
	student.CreateStudent(123, "Fatima", "qwer")
	student.CreateStudent(456, "Tahir", "qwer")
	fmt.Println(student)
	student.Print()
}
