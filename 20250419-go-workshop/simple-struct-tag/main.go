package main

import "fmt"

type User struct {
	Name string `example:"name"`
}

func (u *User) String() string {
	return fmt.Sprintf("Hi! My name is " + u.Name)
}

func main() {
	// Example usage of the User struct
	user := &User{
		Name: "Sammy",
	}

	fmt.Println(user)
}
