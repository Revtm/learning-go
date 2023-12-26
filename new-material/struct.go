package main

import "fmt"

type User struct {
	Name, Email string
	Status      bool
}

func (user User) PrintHello() {
	fmt.Println("Hello", user.Name)
}

func main() {
	var user User
	user.Name = "Iqbal"
	user.Email = "miqbal*********@gmail.com"
	user.Status = true

	user2 := User{
		Name:   "Iqbal 2",
		Email:  "miqbal*********@gmail.com",
		Status: true,
	}

	user3 := User{
		"Iqbal 3", "miqbal********@gmail.com", true,
	}

	fmt.Println(user.Name)
	fmt.Println(user)
	fmt.Println(user2)
	fmt.Println(user3)
	user.PrintHello()
	user2.PrintHello()
	user3.PrintHello()
}
