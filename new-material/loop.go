package main

import "fmt"

func loop(max int) {
	for i := 0; i < max; i++ {
		fmt.Println(i)
	}
}

func main() {
	loop(10)
}
