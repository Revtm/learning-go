package main

import "fmt"

func getSomething() any {
	return 1
}
func main() {
	ans := getSomething()
	var ansStr string
	var ansInt int

	switch ans.(type) {
	case string:
		ansStr = ans.(string)
		fmt.Println(ansStr)
	case int:
		ansInt = ans.(int)
		fmt.Println(ansInt)
	}
}
