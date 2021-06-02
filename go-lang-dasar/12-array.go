package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Febrari"
	names[1] = "Supaldi"

	fmt.Println(names[0])
	fmt.Println(names[1])

	var values = [3]int{
		90,
		95,
		99,
	}

	fmt.Println((values))

	fmt.Println(len(names))
	fmt.Println(len(values))
}
