package main

import (
	"flag"
	"fmt"
)

func main(){
	var name string
	flag.StringVar(&name, "name", "<name>", "The name to say hello to")
	flag.Parse()
	fmt.Printf("Hello %s!\n", name)
}
