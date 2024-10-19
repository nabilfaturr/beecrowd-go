package main

import "fmt"

func main() {
	var n float64
	fmt.Scanln(&n)

	switch {
	case n >= 0 && n <= 25:
		fmt.Println("Intervalo [0,25]")
	case n > 25 && n <= 50:
		fmt.Println("Intervalo (25,50]")
	case n > 50 && n <= 75:
		fmt.Println("Intervalo (50,75]")
	case n > 75 && n <= 100:
		fmt.Println("Intervalo (75,100]")
	default:
		fmt.Println("Fora de intervalo")
	}
}
