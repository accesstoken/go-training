package main

import (
	"flag"
)

var (
	name string
	shhh bool
)

func arg() {
	flag.StringVar(&name, "name", "<name>", "The name to say hello to")
	flag.BoolVar(&shhh, "shhh", false, "verbose")
	flag.Parse()
}
