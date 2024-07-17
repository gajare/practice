package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) AddAge() Person {
	p.Age++
	return p
}

func (p Person) ReName(name string) Person {
	p.Name = name
	return p
}
func main() {
	p := Person{Name: "banti", Age: 30} //not bad
	p.AddAge()
	p.ReName("amol")
	fmt.Println("p :", p)
	p1 := Person{Name: "naina", Age: 21}.AddAge().ReName("nayu") //better
	fmt.Println("p1 :", p1)
	p2 := Person{Name: "sandip", Age: 21}.
		AddAge().
		ReName("chandip")
	fmt.Println("p2 :", p2)
}
