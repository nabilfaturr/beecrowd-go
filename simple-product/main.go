package main

import "fmt"

func main() {
	var a int
	var b int

	fmt.Scanln(&a)
	fmt.Scanln(&b)

	product := a * b

	fmt.Printf("PROD = %d\n", product)

}