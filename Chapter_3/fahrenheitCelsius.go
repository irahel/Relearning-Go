package main

import "fmt"

func main() {
	fmt.Println("Digite o valor em Fahrenheit:")
	var input float64
	fmt.Scanf("%f", &input)

	input = (input - 32.0) * 5 / 9
	fmt.Println("O valor em Celsius é: ", input)
}
