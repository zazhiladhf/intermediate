package main

import "fmt"

type Student struct{
	Name string
	Class string
}

func (s *Student) SetMyName(name string) {
	fmt.Println("Try to change Student from ", s.Name, " name to =>", name)
	s.Name = name
}

func (s Student) CallMyName() {
	fmt.Println("Hello, My name is ", s.Name)
}

func main(){
	student := Student{Name: "Budi", Class: "7"}
	student.CallMyName()
	fmt.Println(student)

	student.SetMyName("Agus")
	fmt.Println(student)

	student.CallMyName()
	fmt.Println(student)
}