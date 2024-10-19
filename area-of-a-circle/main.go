package main

import "fmt"

func main () {
	var r float64;
	var phi float64 = 3.14159;

	fmt.Scanln(&r);

	area := phi * (r * r)


	fmt.Printf("A=%.4f\n", area)

}