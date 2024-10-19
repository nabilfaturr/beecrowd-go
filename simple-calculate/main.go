package main

import "fmt"

func main() {

	var (
		product1Code, product2Code int
		product1Price, product2Price float64
		product1Unit, product2Unit int 
	)

	fmt.Scanf("%d %d %f", &product1Code, &product1Unit, &product1Price)
	fmt.Scanf("%d %d %f", &product2Code, &product2Unit, &product2Price)

	product1TotalPrice := float64(product1Unit) * product1Price
	product2TotalPrice := float64(product2Unit) * product2Price

	totalPrice := product1TotalPrice + product2TotalPrice

	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", totalPrice)

}