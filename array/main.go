package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("slice :", slice)
	var array [5]int
	fmt.Println("array :", array)
	copy(array[:], slice)
	fmt.Println("slice :", slice)
	fmt.Println("Array :", array)

	a:=[]int{10,20,30,40,50,60}
	b:=*(*[3]int)(a[0:3])
	fmt.Println("b :",b)
}
