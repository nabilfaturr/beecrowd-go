package main

import "fmt"

func main () {

	var a, b, resultAvg float64;

	fmt.Scanln(&a);
	fmt.Scanln(&b);

	resultAvg = ((a * 3.5) + (b * 7.5)) / 11;

	fmt.Printf("MEDIA = %.5f\n", resultAvg);

}