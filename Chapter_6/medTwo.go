package main

import "fmt"

func halfWithBool(cal int) (half int, sus bool) {
	half = cal / 2
	if cal%2 == 0 {
		sus = true
	} else {
		sus = false
	}
	return
}

func main() {
	value, x := halfWithBool(1)
	fmt.Println(value, x)

	value, x = halfWithBool(2)
	fmt.Println(value, x)
}
