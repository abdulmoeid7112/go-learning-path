package main

import "fmt"

type Courses struct {
	Name   string
	Slug   string
	Skills []string
}

func (p Courses) Subscribe(name string) {
	fmt.Printf("The person %s has subscribed to %s course\n", name, p.Name)
}

type Career struct {
	Courses
}

func (p Career) Subscribe(name string) {
	fmt.Printf("The person %s has subscribed to %s carrer\n", name, p.Name)
}

type Organization interface {
	Subscribe(name string)
}

func callSubscribe(p Organization) {
	p.Subscribe("Professor")
}

func InterfaceTest() {
	myCourse := Courses{Name: "Go", Slug: "go", Skills: []string{"1", "2"}}
	myCourse.Subscribe("Berlin")

	myCareer := new(Career)

	myCareer.Name = "Go 2"
	myCareer.Slug = "go2"
	myCareer.Skills = []string{"3", "4"}

	callSubscribe(myCourse)
	callSubscribe(myCareer)
}

func main() {
	InterfaceTest()
}
