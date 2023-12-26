package main

import "fmt"

type UserAccount struct {
	Name   string
	Status bool
}

func changeName(user *UserAccount, name string) {
	user.Name = name
}

func testPointer() {
	user := UserAccount{"Iqbal", true}
	user2 := &user

	user2.Name = "Revantama"

	fmt.Println(user)
	fmt.Println(user2.Name)

	*user2 = UserAccount{"Iqbal 2", true}

	fmt.Println(user)
	fmt.Println(user2.Name)

	user3 := new(UserAccount)
	user4 := user3

	user4.Name = "Yuhu"
	user4.Status = false

	fmt.Println(user3)
	fmt.Println(user4)
}

func main() {
	testPointer()

	user := UserAccount{"Iqbal", true}
	changeName(&user, "Revantama")

	fmt.Println(user)
}
