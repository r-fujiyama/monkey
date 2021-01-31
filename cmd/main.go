package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programing language! \n", user.Username)
	fmt.Println("Fell free to type in commands")
	repl.Start(os.Stdin, os.Stdout)
}
