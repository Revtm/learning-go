package main

import "fmt"

type UserInterface interface {
	GetHelloWording() string
}

type User struct {
	Name  string
	Email string
}

func (user User) GetHelloWording() string {
	return user.Name
}

func printHelloWording(user UserInterface) {
	fmt.Println("Hello", user.GetHelloWording())
}

func main() {
	user := User{"Iqbal", "miqbal********@gmail.com"}
	printHelloWording(user)
}
