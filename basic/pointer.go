package main

import "fmt"

func ptlearn() {
	var p int = 1
	addressP := &p
	fmt.Printf("P: %v\n", p)
	fmt.Printf("Address: %v\n", addressP)
	fmt.Printf("*Address: %v\n", *addressP)
	fmt.Printf("&P: %v\n", &p)
}
