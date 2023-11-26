package main

import "fmt"

func itflearn() {
	var i interface{} = "Yo"
	fmt.Println(i)

	word, ok := i.(string)
	if ok {
		fmt.Println(word)
	}
}
