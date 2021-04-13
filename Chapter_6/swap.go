package main

import "fmt"

func swap(a *int, b *int) {
	aux := *a
	*a = *b
	*b = aux
}
func main() {
	x := 4
	y := 90
	fmt.Println(x, y)
	swap(&x, &y)
	fmt.Println(x, y)
}
