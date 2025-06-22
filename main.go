package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/goforj/godump"

	"github.com/bauerbrun0/writing-an-interpreter-in-go/repl"
)

type Profile struct {
	Age   int
	Email string
}

type User struct {
	Name    string
	Profile Profile
	extra   string
}

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	user2 := User{
		Name: "Alice",
		Profile: Profile{
			Age:   30,
			Email: "alice@example.com",
		},
		extra: "asd",
	}

	// Pretty-print to stdout
	godump.Dump(user2)

	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
