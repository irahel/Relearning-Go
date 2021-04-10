package main

import "fmt"

func main() {
	var x map[string]int
	x = make(map[string]int)
	x["key"] = 10
	x["key2"] = 141
	fmt.Println(x)
}
