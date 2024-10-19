package main

import "fmt"

func main() {
	var a int8
	var b int8

	fmt.Scanln(&a)
	fmt.Scanln(&b)

	SOMA := a + b

	fmt.Printf("SOMA = %d\n", SOMA)

}