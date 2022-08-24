package main

import "fmt"

type User struct {
	Name string
}

// type NormalUser struct {
// 	name string
// }

// func (u *NormalUser) getName() string {
// 	return u.name
// }

// func (u *NormalUser) setName(name string) {
// 	if len(name) < 4 {
// 		fmt.Printf("name should contain more than 3 characters: %s\n", name)
// 	}
// 	u.name = name
// }

func main() {
	user := &User{Name: "Damir"}
	fmt.Println(user)

	user.Name = "kek"
	fmt.Println(user)

	user.Name = ""
	fmt.Println(user)

	fmt.Println("-----------")

	// user2 := &NormalUser{}
	// fmt.Println(user2)

	// user2.setName("")
	// fmt.Println(user2)

	// user2.setName("Damir")
	// fmt.Println(user2)
}
