package main

import "fmt"

func main() {
	defer MultiStageDefer()() //defer didnt work on function body but works on return statementn caze return in main and main having a defer
	fmt.Println("Main function called...")
}

func MultiStageDefer() func() {
	fmt.Println("Run initialization...")
	return func() {
		fmt.Println("Run CleanUp...") // defer works only in that block
	}
}
