package main

import (
	"fmt"
)

func main() {
	arg.Arg()
	if !arg.Shhh {
		fmt.Printf("Hello %s!\n", arg.Name)
	}

}
