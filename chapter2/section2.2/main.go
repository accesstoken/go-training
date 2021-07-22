package main

import (
	"flag"
	"fmt"
)

func main(){
	var (
		name string
		shhh bool
	)
	flag.StringVar(&name, "name", "<name>", "The name to say hello to")
	flag.BoolVar(&shhh, "shhh", false, "verbose")
	flag.Parse()
	if !shhh {
		fmt.Printf("Hello %s!\n", name)
	}
}
