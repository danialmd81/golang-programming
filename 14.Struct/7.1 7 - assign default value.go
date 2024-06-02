package main

import "fmt"

//Student strct
type Student struct {
	Name string
	Age  int
}

//Instant Method
func (s *Student) Instant() {
	if s.Name == "" {
		s.Name = "Robert"
	}
	if s.Age == 0 {
		s.Age = 25
	}
}

func main() {
	student1 := Student{}
	student1.Instant()
	fmt.Println(student1)

	student2 := Student{Name: "Julia"}
	student2.Instant()
	fmt.Println(student2)

	student3 := Student{Age: 30}
	student3.Instant()
	fmt.Println(student3)
}
