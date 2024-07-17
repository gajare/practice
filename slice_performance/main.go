package main

import "fmt"

func main() {
	//Wrong
	s1 := make([]int, 5)
	s1 = append(s1, 1)
	fmt.Println("s1 :", s1)	//s1 : [0 0 0 0 0 1]
	fmt.Println("length")
	//correct
	s2 := make([]int, 0, 5)
	s2 = append(s2, 1)
	fmt.Println("s2 :", s2)	//s2 : [1]
}
