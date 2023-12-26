package main

import "fmt"

func main() {
	count := 0
	increment := func() {
		count++
	}

	increment()
	fmt.Println(count)
}
