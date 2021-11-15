package main

import "fmt"

func hello() {
	fmt.Println("Hello!")
}

func main() {
	fmt.Println("Hello, People!")

	for i := 0; i < 15; i++ {
		fmt.Println(i)
	}

	hello()
}
