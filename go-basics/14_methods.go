package main

import "fmt"

type CourseStruct struct {
	Name   string
	Slug   string
	Skills []string
}

func (p CourseStruct) Subscribe(name string) {
	fmt.Printf("The person %s has subscribed to %s course\n", name, p.Name)
}

func main() {
	myCourse1 := CourseStruct{Name: "Go", Slug: "go", Skills: []string{"1", "2"}}

	myCourse1.Subscribe("Berlin")

	fmt.Println(myCourse1)
}
