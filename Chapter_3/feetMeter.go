package main

import "fmt"

func main() {
	fmt.Println("Digite o valor em pés:")
	var input float64
	fmt.Scanf("%f", &input)

	input = input * 0.3048
	fmt.Println("O valor em metros é: ", input)
}
