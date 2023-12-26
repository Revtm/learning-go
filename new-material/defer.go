package main

import (
	"fmt"
	"strconv"
)

func logging() {
	err := recover()
	if err != nil {
		fmt.Println("log:", err)
	}
}

func printPrice() {
	defer logging()
	fmt.Println("Rp 10000")
}

func printPriceWithDisc(disc int) {
	defer logging()
	if disc < 1 {
		panic("Unexpected Discount: " + strconv.Itoa(disc))
	}

	fmt.Println("Rp 10000 diskon", disc, "persen")
}

func main() {
	printPrice()
	printPriceWithDisc(0)
	fmt.Println("Finished")
}
